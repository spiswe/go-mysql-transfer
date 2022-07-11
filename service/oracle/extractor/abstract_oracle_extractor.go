package extractor

import (
	"database/sql"
	"go-mysql-transfer/service/oracle/models"
	meta "go-mysql-transfer/service/oracle/oracle_meta"
	"reflect"
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

func (a *AbstractOracleRecordExtractor) GetColumnValues(rs *sql.Rows, encoding string, col []meta.ColumnMeta) []models.ColumnValue {
	columnTypes, _ := rs.ColumnTypes()
	rowValues := make([]reflect.Value, len(columnTypes))
	rowResult := make([]interface{}, len(columnTypes))
	var columnValues []models.ColumnValue
	for i := 0; i < len(columnTypes); i++ {
		rowValues[i] = reflect.New(reflect.PtrTo(columnTypes[i].ScanType()))
		rowResult[i] = rowValues[i].Interface()
	}

	_ = rs.Scan(rowResult...)

	for i := 0; i < len(rowValues); i++ {
		if rv := rowValues[i].Elem(); rv.IsNil() {
			rowResult[i] = nil
		} else {
			rowResult[i] = rv.Elem().Interface()
		}

		//columnValue := ColumnValue{column: ColumnMeta{name: col[i].name, columnType: col[i].columnType}, value: rowResult[i]}
		columnValue := models.NewColumnValueWithoutCheck(col[i], rowResult[i])
		columnValues = append(columnValues, *columnValue)
	}
	return columnValues
}

//func (a *AbstractRecordExtractor) getColumnValue(
//	rs *sql.Rows,
//	encoding string,
//	col meta.ColumnMeta) models.ColumnValue {
//	if col.ColumnType() == go_ora.DATE {
//		value := rs.Scan()
//	}
//	//var value interface{}
//	//if col.ColumnType() == go_ora.DATE {
//	//	var ret sql.NullTime
//	//	value = rs.Scan(&ret)
//	//}
//
//	// todo need clone
//	cv := models.ColumnValue{}
//	return cv
//}
