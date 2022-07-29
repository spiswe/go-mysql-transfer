package extractor

import (
	"database/sql"
	"fmt"
	"go-mysql-transfer/service/oracle"
	"go-mysql-transfer/service/oracle/models"
	meta "go-mysql-transfer/service/oracle/oracle_meta"
	"go-mysql-transfer/service/oracle/positioner"
	"reflect"
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

type ContinueExtractorFull struct {
	ds      *sql.DB
	id      interface{}
	running bool
}

func (c *ContinueExtractorFull) NewContinueExtractorFull(extractor *OracleFullRecordExtractor) {
	position := extractor.oracleContext.LastPosition()
	datasource := extractor.oracleContext.SourceDS()
	c.ds = datasource.Conn()

	if position != nil {

		// tell you that this is all for m_logs
		idPosition := position.(*positioner.IDPosition)

		if idPosition.CurrentProgress() == models.FULLING {
			id := idPosition.ID()
		}
	}
}

func (c *ContinueExtractorFull) GetMinID(sql string) interface{} {
	var minId interface{}
	if c.ds != nil && strings.Trim(sql, " ") != "" {
		rows, _ := c.ds.Query(sql)
		defer func() {
			_ = rows.Close()
		}()

		for rows.Next() {
			_ = rows.Scan(&minId)
			break
		}

		if minId != nil {
			if reflect.TypeOf(minId) == reflect.TypeOf(1) {
				minId = minId.(int)
			} else {
				minId = ""
			}
		} else {
			if reflect.TypeOf(minId) == reflect.TypeOf(1) {
				minId = 0
			} else {
				minId = ""
			}
		}
		return minId
	} else {
		panic("datasource or getMinPkSQL is nil")
	}
}

//func (o *OracleFullRecordExtractor) ContinueExtractor(context models.OracleContext) {
//	position := context.LastPosition()
//	o.oracleContext
//
//}
