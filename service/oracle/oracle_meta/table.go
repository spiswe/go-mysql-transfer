/*
table oracle_meta management for oracle
*/

package meta

type Table struct {
	tableType   string
	schema      string
	name        string
	primaryKeys []ColumnMeta
	columns     []ColumnMeta

	//ext key 其他
	extKey string
}

func (t *Table) TableType() string {
	return t.tableType
}

func (t *Table) SetTableType(tableType string) {
	t.tableType = tableType
}

func (t *Table) Schema() string {
	return t.schema
}

func (t *Table) SetSchema(schema string) {
	t.schema = schema
}

func (t *Table) Name() string {
	return t.name
}

func (t *Table) SetName(name string) {
	t.name = name
}

func (t *Table) PrimaryKeys() []ColumnMeta {
	return t.primaryKeys
}

func (t *Table) SetPrimaryKeys(primaryKeys []ColumnMeta) {
	t.primaryKeys = primaryKeys
}

func (t *Table) Columns() []ColumnMeta {
	return t.columns
}

func (t *Table) SetColumns(columns []ColumnMeta) {
	t.columns = columns
}

func (t *Table) ExtKey() string {
	return t.extKey
}

func (t *Table) GetColumnsWithPrimary() []ColumnMeta {
	result := append(t.columns, t.primaryKeys...)
	return result
}

func NewTable(tableType string, schema string, name string) *Table {
	return &Table{tableType: tableType, schema: schema, name: name}
}

func (t *Table) GetFullName() string {
	return t.schema + t.name
}

func NewTableWithColumn(
	tableType string,
	schema string,
	name string,
	primaryKeys []ColumnMeta,
	columns []ColumnMeta) *Table {
	return &Table{
		tableType:   tableType,
		name:        name,
		schema:      schema,
		primaryKeys: primaryKeys,
		columns:     columns}
}

func (t *Table) ToString() string {
	return ""
}

//func (t *Table) getType() string {
//	return t.tableType
//}
//
//func (t *Table) setType(tableType string) {
//	t.tableType = tableType
//}
//
//func (t *Table) getName() string {
//	return t.name
//}
//
//func (t *Table) setName(name string) {
//	t.name = name
//}
//
//func (t *Table) getSchema() string {
//	return t.schema
//}
//
//func (t *Table) setSchema(schema string) {
//	t.schema = schema
//}
//
//func (t *Table) getPrimaryKeys() []ColumnMeta {
//	return t.primaryKeys
//}
//
//func (t *Table) setPrimaryKeys(primaryKeys []ColumnMeta) {
//	t.primaryKeys = primaryKeys
//}
//
//func (t *Table) addPrimaryKeys(primaryKey ColumnMeta) {
//	t.primaryKeys = append(t.primaryKeys, primaryKey)
//}
//
//func (t *Table) getColumns() []ColumnMeta {
//	return t.columns
//}
//
//func (t *Table) setColumns(columns []ColumnMeta) {
//	t.columns = columns
//}
//
//func (t *Table) AddColumn(column ColumnMeta) {
//	t.columns = append(t.columns, column)
//}
//
//func (t *Table) GetColumnWithPrimaryKey() []ColumnMeta {
//	var result []ColumnMeta
//	result = append(t.primaryKeys, t.columns...)
//	return result
//}
//
//func (t *Table) GetExtKey() string {
//	return t.extKey
//}
//
//func (t *Table) SetExtKey(extKey string) {
//	t.extKey = extKey
//}
//
