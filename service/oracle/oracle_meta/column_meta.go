/*
table column oracle_meta
*/

package meta

import go_ora "github.com/sijms/go-ora/v2"

type ColumnMeta struct {
	name       string
	columnType go_ora.OracleType
}

func (cm *ColumnMeta) Name() string {
	return cm.name
}

func (cm *ColumnMeta) SetName(name string) {
	cm.name = name
}

func (cm *ColumnMeta) ColumnType() go_ora.OracleType {
	return cm.columnType
}

func (cm *ColumnMeta) SetColumnType(columnType go_ora.OracleType) {
	cm.columnType = columnType
}

func NewColumnMeta(name string, columnType go_ora.OracleType) (cm *ColumnMeta) {
	return &ColumnMeta{name: name, columnType: columnType}
}

func (cm *ColumnMeta) clone() ColumnMeta {
	return *NewColumnMeta(cm.name, cm.columnType)
}

// todo 序列化实现
func (cm *ColumnMeta) toString() string {
	return ""
}
