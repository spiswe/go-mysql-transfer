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

func NewTable(tableType string, schema string, name string) *Table {
	return &Table{tableType: tableType, schema: schema, name: name}
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

func (t *Table) getType() string {
	return t.tableType
}

func (t *Table) setType(tableType string) {
	t.tableType = tableType
}

func (t *Table) getName() string {
	return t.name
}

func (t *Table) setName(name string) {
	t.name = name
}

func (t *Table) getSchema() string {
	return t.schema
}

func (t *Table) setSchema(schema string) {
	t.schema = schema
}

func (t *Table) getPrimaryKeys() []ColumnMeta {
	return t.primaryKeys
}

func (t *Table) setPrimaryKeys(primaryKeys []ColumnMeta) {
	t.primaryKeys = primaryKeys
}

func (t *Table) addPrimaryKeys(primaryKey ColumnMeta) {
	t.primaryKeys = append(t.primaryKeys, primaryKey)
}

func (t *Table) getColumns() []ColumnMeta {
	return t.columns
}

func (t *Table) setColumns(columns []ColumnMeta) {
	t.columns = columns
}

func (t *Table) addColumn(column ColumnMeta) {
	t.columns = append(t.columns, column)
}

func (t *Table) getColumnWithPrimaryKey() []ColumnMeta {
	var result []ColumnMeta
	result = append(t.primaryKeys, t.columns...)
	return result
}

func (t *Table) getExtKey() string {
	return t.extKey
}

func (t *Table) setExtKey(extKey string) {
	t.extKey = extKey
}

func (t *Table) getFullName() string {
	return t.schema + t.name
}

func (t *Table) toString() string {
	return ""
}
