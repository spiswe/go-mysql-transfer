package extractor

import (
	"database/sql"
	go_ora "github.com/sijms/go-ora/v2"
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
	rs *sql.Rows,
	encoding string,
	col meta.ColumnMeta) models.ColumnValue {
	if col.ColumnType() == go_ora.DATE {
		value := rs.Scan()
	}
	//var value interface{}
	//if col.ColumnType() == go_ora.DATE {
	//	var ret sql.NullTime
	//	value = rs.Scan(&ret)
	//}

	// todo need clone
	cv := models.ColumnValue{}
	return cv
}
