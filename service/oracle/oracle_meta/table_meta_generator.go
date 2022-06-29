package meta

import (
	"fmt"
	"go-mysql-transfer/service/oracle/database"
	"go-mysql-transfer/service/oracle/oracle_meta/schema"
	"go-mysql-transfer/service/oracle/utils"
	"go-mysql-transfer/util/logs"
	"strings"
)

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

var TableMetaGeneratorController *TableMetaGenerator

func (t *TableMetaGenerator) getTableMeta(
	dataSource database.DataSource,
	schemaName string,
	tableName string) Table {
	//dataSource.Conn()
	return Table{}
}

func (t *TableMetaGenerator) GetTableMetasWithoutColumn(
	dataSource database.DataSource,
	schemaName string,
	tableName string) []Table {

	var result []Table

	dt := fmt.Sprintf("%T", dataSource.Conn().Driver())
	if dt == "*go_ora.OracleDriver" && schemaName == "" && tableName == "" {
		metadatas, err := schema.TableNames(dataSource.Conn(), schemaName)
		if err != nil {
			logs.Error("Get Table metaData error " + err.Error())
		}
		for _, metadata := range metadatas {
			if !(strings.HasPrefix(strings.ToUpper(metadata[1]), "MLOG$_") ||
				strings.HasPrefix(strings.ToUpper(metadata[1]), "RUPD$_")) {
				table := NewTable("TABLE", metadata[0], metadata[1])
				result = append(result, *table)
			}
		}
	} else {
		// use column type func to check if the table exists
		ct, _ := schema.ColumnTypes(dataSource.Conn(), schemaName, tableName)
		if len(ct) != 0 &&
			!(strings.HasPrefix(strings.ToUpper(tableName), "MLOG$_") ||
				strings.HasPrefix(strings.ToUpper(tableName), "RUPD$_")) {
			table := NewTable("TABLE", schemaName, tableName)
			result = append(result, *table)
		}
	}
	return result
}

func (t *TableMetaGenerator) BuildColumns(dataSource database.DataSource, table *Table) {
	columns, _ := schema.ColumnTypes(dataSource.Conn(), table.Schema(), table.Name())
	primaryKeys, _ := schema.PrimaryKey(dataSource.Conn(), table.Schema(), table.Name())
	var columnList, primaryKeyList []ColumnMeta
	for _, col := range columns {
		if utils.SliceContains(primaryKeys, col.Name()) {
			primaryKeyMeta := NewColumnMeta(col.Name(), utils.SwitchToGoOraType(col.DatabaseTypeName()))
			primaryKeyList = append(primaryKeyList, *primaryKeyMeta)
		} else {
			columnMeta := NewColumnMeta(col.Name(), utils.SwitchToGoOraType(col.DatabaseTypeName()))
			columnList = append(columnList, *columnMeta)
		}

	}
	table.SetColumns(columnList)
	table.SetPrimaryKeys(primaryKeyList)
}
