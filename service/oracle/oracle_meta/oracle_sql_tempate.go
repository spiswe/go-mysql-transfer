package meta

var OracleSQLTemplateControl = &OracleSQLTemplate{}

type OracleSQLTemplate struct {
	sqlTemplate SQLTemplate
}

func (o *OracleSQLTemplate) SqlTemplate() SQLTemplate {
	return o.sqlTemplate
}

func (o *OracleSQLTemplate) SetSqlTemplate(sqlTemplate SQLTemplate) {
	o.sqlTemplate = sqlTemplate
}

func (o *OracleSQLTemplate) getMergeSQL(schemaName string, tableName string, pkNames []string, colNames []string) string {
	return ""
}
