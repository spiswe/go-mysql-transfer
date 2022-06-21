/*
table column oracle_meta
*/

package meta

type ColumnMeta struct {
	name       string
	columnType int
}

func (cm *ColumnMeta) Name() string {
	return cm.name
}

func (cm *ColumnMeta) SetName(name string) {
	cm.name = name
}

func (cm *ColumnMeta) ColumnType() int {
	return cm.columnType
}

func (cm *ColumnMeta) SetColumnType(columnType int) {
	cm.columnType = columnType
}

func NewColumnMeta(name string, columnType int) (cm *ColumnMeta) {
	return &ColumnMeta{name: name, columnType: columnType}
}

func (cm *ColumnMeta) clone() ColumnMeta {
	return *NewColumnMeta(cm.name, cm.columnType)
}

// todo 序列化实现
func (cm *ColumnMeta) toString() string {
	return ""
}
