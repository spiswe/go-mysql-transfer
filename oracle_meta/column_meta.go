/*
table column oracle_meta
*/

package meta

type ColumnMeta struct {
	name       string
	columnType int
}

func NewColumnMeta(name string, columnType int) (cm *ColumnMeta) {
	return &ColumnMeta{name: name, columnType: columnType}
}

func (cm *ColumnMeta) getName() string {
	return cm.name
}

func (cm *ColumnMeta) getType() int {
	return cm.columnType
}

func (cm *ColumnMeta) setName(name string) {
	cm.name = name
}

func (cm *ColumnMeta) setType(columnType int) {
	cm.columnType = columnType
}

func (cm *ColumnMeta) clone() ColumnMeta {
	return *NewColumnMeta(cm.name, cm.columnType)
}

// todo 序列化实现
func (cm *ColumnMeta) toString() string {
	return ""
}
