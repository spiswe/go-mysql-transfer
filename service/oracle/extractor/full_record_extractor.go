package extractor

import (
	"fmt"
	"go-mysql-transfer/service/oracle"
	"go-mysql-transfer/service/oracle/models"
	meta "go-mysql-transfer/service/oracle/oracle_meta"
	"strings"
)

const (
	FORMAT        string = "select * from (select %s from %s.%s t where %s > ? order by {3} asc) where rownum <= ?"
	MIN_PK_FORMAT string = "select min(%s) from {%s}.{%s}"
)

type OracleFullRecordExtractor struct {
	absOraRecExt  AbstractOracleRecordExtractor
	oracleContext models.OracleContext
	queue         chan models.Record
	extractSQL    string
	extractThread oracle.Pool
	getMinPkSQL   string
}

func NewOracleFullRecordExtractor(oracleContext models.OracleContext) *OracleFullRecordExtractor {
	return &OracleFullRecordExtractor{oracleContext: oracleContext}
}

func (o *OracleFullRecordExtractor) start() {
	extractor := o.absOraRecExt.Extractor()
	extractor.Start()
	tableMeta := o.oracleContext.TableMeta()
	primaryKey := tableMeta.PrimaryKeys()[0].Name()
	schemaName := tableMeta.Schema()
	tableName := tableMeta.Name()

	if o.extractSQL == "" {
		columnWithPk := tableMeta.GetColumnsWithPrimary()
		colString := meta.SQLTemplateControl.MakeColumn(columnWithPk)
		o.extractSQL = fmt.Sprintf(FORMAT, colString, schemaName, tableName, primaryKey)
	}

	if o.getMinPkSQL == "" && strings.Trim(primaryKey, " ") != "" {
		o.getMinPkSQL = fmt.Sprintf(MIN_PK_FORMAT, primaryKey, schemaName, tableName)
	}

}

func (o *OracleFullRecordExtractor) ContinueExtractor(context models.OracleContext) {
	position := context.LastPosition()

}
