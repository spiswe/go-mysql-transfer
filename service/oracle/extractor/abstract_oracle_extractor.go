package extractor

import (
	"database/sql"
	"go-mysql-transfer/service/oracle/models"
	meta "go-mysql-transfer/service/oracle/oracle_meta"
)

// https://github.com/golang/go/wiki/SQLDrivers

type AbstractOracleRecordExtractor struct {
	extractor AbstractRecordExtractor
}

func (a *AbstractOracleRecordExtractor) Extractor() AbstractRecordExtractor {
	return a.extractor
}

func (a *AbstractOracleRecordExtractor) SetExtractor(extractor AbstractRecordExtractor) {
	a.extractor = extractor
}

func (a *AbstractRecordExtractor) getColumnValue(
	rs *sql.Row,
	encoding string,
	col meta.ColumnMeta) models.ColumnValue {

	//var value interface{}
	//if col.ColumnType() == go_ora.DATE {
	//	var ret sql.NullTime
	//	value = rs.Scan(&ret)
	//}

	// todo need clone
	cv := models.ColumnValue{}
	return cv
}
