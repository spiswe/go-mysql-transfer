package extractor

import (
	"fmt"
	"go-mysql-transfer/service/oracle/models"
	meta "go-mysql-transfer/service/oracle/oracle_meta"
)

const (
	FORMATALL string = "select /*+parallel(t)*/ {0} from {1}.{2} t"
)

type OracleFullRecordExtractorOnce struct {
	absOraRecExt  AbstractOracleRecordExtractor
	oracleContext models.OracleContext
	extractSQL    string
}

func NewOracleFullRecordExtractorOnce(oracleContext models.OracleContext) *OracleFullRecordExtractorOnce {
	return &OracleFullRecordExtractorOnce{oracleContext: oracleContext}
}

func (o *OracleFullRecordExtractorOnce) SetExtractSQL(extractSQL string) {
	o.extractSQL = extractSQL
}

func (o *OracleFullRecordExtractorOnce) start() {
	//o.absOraRecExt.extractor.AbsLifeCycle()
	//extractor.AbsLifeCycle()
	if o.extractSQL == "" {
		tableMeta := o.oracleContext.TableMeta()
		columns := meta.SQLTemplateControl.MakeColumn(tableMeta.GetColumnsWithPrimary())
		o.extractSQL = fmt.Sprintf(FORMATALL, columns, tableMeta.Schema(), tableMeta.Name())
	}

	// todo 增加线程控制调用
}
