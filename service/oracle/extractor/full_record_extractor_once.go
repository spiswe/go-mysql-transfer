package extractor

import (
	"database/sql"
	"fmt"
	"go-mysql-transfer/service/oracle"
	"go-mysql-transfer/service/oracle/models"
	meta "go-mysql-transfer/service/oracle/oracle_meta"
	"go-mysql-transfer/util/logs"
)

const (
	FORMATALL string = "select /*+parallel(t)*/ {0} from {1}.{2} t"
)

type OracleFullRecordExtractorOnce struct {
	absOraRecExt  AbstractOracleRecordExtractor
	oracleContext models.OracleContext
	extractSQL    string
	queue         chan models.Record
}

//type ContinueExtractor struct {
//	// goroutine
//}

func NewOracleFullRecordExtractorOnce(oracleContext models.OracleContext) *OracleFullRecordExtractorOnce {
	return &OracleFullRecordExtractorOnce{oracleContext: oracleContext}
}

func (o *OracleFullRecordExtractorOnce) SetExtractSQL(extractSQL string) {
	o.extractSQL = extractSQL
}

func (o *OracleFullRecordExtractorOnce) Start() {
	// todo 需要增加对当前运行状态的判断
	//o.absOraRecExt.extractor.AbsLifeCycle()
	//extractor.AbsLifeCycle()
	tableMeta := o.oracleContext.TableMeta()
	if o.extractSQL == "" {
		columns := meta.SQLTemplateControl.MakeColumn(tableMeta.GetColumnsWithPrimary())
		o.extractSQL = fmt.Sprintf(FORMATALL, columns, tableMeta.Schema(), tableMeta.Name())
	}
	// todo 需要增加多线程支持

	// update status
	oracle.TracerController.Update(tableMeta.GetFullName(), models.FULLING)
}

func (o *OracleFullRecordExtractorOnce) Stop() {
	// todo thread.stop
}

func (o *OracleFullRecordExtractorOnce) Extract() []models.Record {
	records := make([]models.Record, o.oracleContext.OnceCrawNum())
	for i := 0; i < o.oracleContext.OnceCrawNum(); i++ {
		//todo put record to queue, need add queue
		record := <-o.queue
		if record.TableName() != "" {
			records = append(records, record)
		} else if o.absOraRecExt.Extractor().status == models.TABLEEND {
			// todo 确认是否真的end了
			record := <-o.queue
			if record.TableName() != "" {
				records = append(records, record)
			} else {
				// real end
				break
			}
		} else {
			// cannot get data
			i--
			continue
		}
	}
	return records
}

func (o *OracleFullRecordExtractorOnce) ContinueExtractor(rows *sql.Rows) {
	for rows.Next() {
		tableMeta := o.oracleContext.TableMeta()

		pks := o.absOraRecExt.GetColumnValues(rows, o.oracleContext.GetSourceCodeEncoding(), tableMeta.PrimaryKeys())
		cms := o.absOraRecExt.GetColumnValues(rows, o.oracleContext.GetSourceCodeEncoding(), tableMeta.Columns())

		re := models.NewRecord(tableMeta.Schema(), tableMeta.Name(), pks, cms)
		logs.Info(re.SchemaName())
		o.queue <- *re
	}
	extractor := o.absOraRecExt.Extractor()
	extractor.SetStatus(models.TABLEEND)
}
