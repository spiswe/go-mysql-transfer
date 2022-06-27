package meta

import "go-mysql-transfer/service/oracle/database"

const (
	MLOGQUERYSQL       = "select master,log_table from all_mview_logs where master = %s"
	MLOGSCHEMAQUERYSQL = "select master,log_table from all_mview_logs where master = %s and log_owner = %s "
	QUERYSHAREDKEY     = "show partitions from %s"
)

type TableMetaGenerator struct {
	mLogQuerySql       string
	mLogSchemaQuerySQl string
	querySharedKey     string
}

func (t *TableMetaGenerator) getTableMeta(dataSource database.DataSource, schemaName string, tableName string) Table {
	//dataSource.Conn()
	return Table{}
}

//func (t *TableMetaGenerator)
