package model

import (
	"go-mysql-transfer/service/oracle/utils"
	"strings"
)

type Record struct {
	schemaName  string
	tableName   string
	primaryKeys []ColumnValue
	columns     []ColumnValue
}

func NewRecord(schemaName string, tableName string, primaryKeys []ColumnValue, columns []ColumnValue) *Record {
	return &Record{schemaName: schemaName, tableName: tableName, primaryKeys: primaryKeys, columns: columns}
}

func (r *Record) SchemaName() string {
	return r.schemaName
}

func (r *Record) TableName() string {
	return r.tableName
}

func (r *Record) PrimaryKeys() []ColumnValue {
	return r.primaryKeys
}

func (r *Record) Columns() []ColumnValue {
	return r.columns
}

func (r *Record) SetSchemaName(schemaName string) {
	r.schemaName = schemaName
}

func (r *Record) SetTableName(tableName string) {
	r.tableName = tableName
}

func (r *Record) SetPrimaryKeys(primaryKeys []ColumnValue) {
	r.primaryKeys = primaryKeys
}

func (r *Record) SetColumns(columns []ColumnValue) {
	r.columns = columns
}

func (r *Record) getColumnByNameWithFalse(columnName string) ColumnValue {
	return r.getColumnByName(columnName, false)
}

func (r *Record) getColumnByName(columnName string, returnNilNotExists bool) ColumnValue {
	for _, column := range r.columns {
		re := column.Column()
		if strings.EqualFold(re.Name(), columnName) {
			return column
		}
	}

	for _, pk := range r.primaryKeys {
		re := pk.Column()
		if strings.EqualFold(re.Name(), columnName) {
			return pk
		}
	}

	if returnNilNotExists {
		return ColumnValue{}
	} else {
		panic("not found column[" + columnName + "]")
	}
}

func (r *Record) removeColumnByNameWithFalse(columnName string) ColumnValue {
	return r.removeColumnByName(columnName, false)
}

func (r *Record) removeColumnByName(columnName string, returnNilNotExists bool) ColumnValue {
	var remove ColumnValue
	var idx int
	for _idx, pk := range r.primaryKeys {
		re := pk.Column()
		if strings.EqualFold(re.Name(), columnName) {
			remove = pk
			idx = _idx
			break
		}
	}
	if remove != (ColumnValue{}) && idx > -1 {
		utils.RemoveItemByIndex(&r.primaryKeys, idx)
		return remove
	} else {
		for _idx, column := range r.columns {
			remove = column
			idx = _idx
			break
		}

		if remove != (ColumnValue{}) && idx > -1 {
			utils.RemoveItemByIndex(&r.columns, idx)
			return remove
		}
	}

	if returnNilNotExists {
		return ColumnValue{}
	} else {
		panic("not found column[" + columnName + "]")
	}

}

func (r *Record) addPrimaryKey(primaryKey ColumnValue) {

}
