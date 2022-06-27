package oracle

import (
	"go-mysql-transfer/global"
	"go-mysql-transfer/service/oracle/database"
	"go-mysql-transfer/service/oracle/models"
	meta "go-mysql-transfer/service/oracle/oracle_meta"
	"go-mysql-transfer/service/oracle/translator"
	"strings"
)

type Controller struct {
	configuration     *global.Config
	dataSourceFactory database.DataSourceFactory
	//config
	// todo

	runMode       models.RunMode
	globalContext models.OracleContext
	sourceDBType  database.DBType
	targetDBType  database.DBType
	// todo
	// alarmService translator DIR
	// table controller to set table extractor
	processTracer ProcessTracer
	instance      Instance
}

type TableHolder struct {
	table        meta.Table
	ignoreSchema bool
	translator   translator.DataTranslator
}

func NewTableHolder(table meta.Table) *TableHolder {
	return &TableHolder{table: table}
}

func NewController(configuration *global.Config) *Controller {
	return &Controller{configuration: configuration}
}

func (c *Controller) start() {
	var dbconfig map[database.DataSourceConfig]database.DataSource
	sourceConfig := database.NewDataSourceConfig(c.configuration)
	dbconfig[*sourceConfig] = database.DataSource{}

	// connect to all databases
	c.dataSourceFactory.SetDataSources(dbconfig)
	c.dataSourceFactory.Start() // set all database connection

	c.SetRunMode(models.RunMode(c.configuration.OracleRunMode))
	c.SetSourceDBType(database.DBType(c.configuration.Flavor))
	// todo, maybe it is useless
	c.SetTargetDBType(database.DBType(c.configuration.Flavor))
	c.globalContext = c.InitGlobalContext()
	// todo add alarm
	// todo add extractorDump, statBufferSize, statePrintInterval

}

func (c *Controller) TableHolder(table meta.Table) {}

//func (c *Controller) InitTables() []TableHolder {
//	logs.Info("check table privileges ...")
//	tableWhiteList := c.configuration.OracleTableWhiteList
//	tableBlackList := c.configuration.OracleTableBlackList
//	isEmpty := true
//	for _, table := range tableWhiteList {
//		if len(table) == 0 {
//			isEmpty = false
//			break
//		}
//	}
//	var tables []TableHolder
//	// target DBType is not important
//	if !isEmpty {
//		for _, obj := range tableWhiteList {
//			whiteTable := c.getTable(obj)
//			if !c.isBlackTable(whiteTable, tableBlackList) {
//				tableInfo := strings.Split(whiteTable, ".")
//				ignoreSchema := false
//				if len(tableInfo) == 1 {
//					whiteTables :=
//				}
//			}
//
//		}
//	}
//}

func (c *Controller) isBlackTable(table string, tableBlackList []string) bool {
	for _, obj := range tableBlackList {
		if table == obj {
			return true
		} else {
			return false
		}
	}
	return false
}

func (c *Controller) getTable(tableName string) string {
	paramArray := strings.Split(tableName, "#")
	if len(paramArray) >= 1 && paramArray[0] != "" {
		return paramArray[0]
	} else {
		return ""
	}
}

func (c *Controller) Configuration() *global.Config {
	return c.configuration
}

func (c *Controller) SetConfiguration(configuration *global.Config) {
	c.configuration = configuration
}

func (c *Controller) DataSourceFactory() database.DataSourceFactory {
	return c.dataSourceFactory
}

func (c *Controller) SetDataSourceFactory(dataSourceFactory database.DataSourceFactory) {
	c.dataSourceFactory = dataSourceFactory
}

func (c *Controller) RunMode() models.RunMode {
	return c.runMode
}

func (c *Controller) SetRunMode(runMode models.RunMode) {
	c.runMode = runMode
}

func (c *Controller) GlobalContext() models.OracleContext {
	return c.globalContext
}

func (c *Controller) SetGlobalContext(globalContext models.OracleContext) {
	c.globalContext = globalContext
}

func (c *Controller) SourceDBType() database.DBType {
	return c.sourceDBType
}

func (c *Controller) SetSourceDBType(sourceDBType database.DBType) {
	c.sourceDBType = sourceDBType
}

func (c *Controller) TargetDBType() database.DBType {
	return c.targetDBType
}

func (c *Controller) SetTargetDBType(targetDBType database.DBType) {
	c.targetDBType = targetDBType
}

func (c *Controller) ProcessTracer() ProcessTracer {
	return c.processTracer
}

func (c *Controller) SetProcessTracer(processTracer ProcessTracer) {
	c.processTracer = processTracer
}

func (c *Controller) Instance() Instance {
	return c.instance
}

func (c *Controller) SetInstance(instance Instance) {
	c.instance = instance
}

// InitDataSource todo need to split different database type && do not use to modify code
func (c *Controller) InitDataSource(initType string) database.DataSource {
	dataSourceConfig := database.NewDataSourceConfig(global.Cfg())
	return c.dataSourceFactory.GetDataSource(*dataSourceConfig)
}

func (c *Controller) InitGlobalContext() models.OracleContext {
	context := models.OracleContext{}
	//context.SetSourceDS()
	context.SetSourceDS(c.InitDataSource("source"))
	context.SetSourceCodeEncoding(c.configuration.OracleEncoding)
	context.SetTargetCodeEncoding(c.configuration.OracleEncoding)
	context.SetBatchApply(c.configuration.OracleBatchApply)
	context.SetOnceCrawNum(c.configuration.OracleCrawNum)
	context.SetTPSLimit(c.configuration.OracleTPSLimit)
	context.SetIgnoreSchema(c.configuration.OracleIgnoreSchema)
	context.SetSkipApplierException(c.configuration.OracleSkipApplierException)
	context.SetRunMode(c.runMode)
	return context
}
