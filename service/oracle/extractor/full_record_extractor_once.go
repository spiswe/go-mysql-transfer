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
}

type ContinueExtractor struct {
	// goroutine

}

func NewOracleFullRecordExtractorOnce(oracleContext models.OracleContext) *OracleFullRecordExtractorOnce {
	return &OracleFullRecordExtractorOnce{oracleContext: oracleContext}
}

func (o *OracleFullRecordExtractorOnce) SetExtractSQL(extractSQL string) {
	o.extractSQL = extractSQL
}

func (o *OracleFullRecordExtractorOnce) start() {
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

func (o *OracleFullRecordExtractorOnce) stop() {
	// todo thread.stop
}

func (o *OracleFullRecordExtractorOnce) Extract() []models.Record {
	record := make([]models.Record, o.oracleContext.OnceCrawNum())
	for i := 0; i < o.oracleContext.OnceCrawNum(); i++ {
		//todo put record to queue, need add queue
	}
	return record
}

func (o *OracleFullRecordExtractorOnce) ContinueExtractor(rows *sql.Rows) {
	for rows.Next() {
		var columns []models.ColumnValue
		var pks []models.ColumnValue
		tableMeta := o.oracleContext.TableMeta()
		for _, pk := range tableMeta.PrimaryKeys() {
			extractor := o.absOraRecExt.Extractor()
			c := extractor.getColumnValue(rows, o.oracleContext.GetSourceCodeEncoding(), pk)
			pks = append(pks, c)
		}

		for _, col := range tableMeta.Columns() {
			extractor := o.absOraRecExt.Extractor()
			c := extractor.getColumnValue(rows, o.oracleContext.GetSourceCodeEncoding(), col)
			columns = append(columns, c)
		}

		re := models.NewRecord(tableMeta.Schema(), tableMeta.Name(), pks, columns)
		logs.Info(re.SchemaName())
		// todo add re to queue
	}
	extractor := o.absOraRecExt.Extractor()
	extractor.SetStatus(models.TABLEEND)
}
