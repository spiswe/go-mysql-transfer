package oracle

import (
	"go-mysql-transfer/global"
	"go-mysql-transfer/service/oracle/database"
	"go-mysql-transfer/service/oracle/extractor"
	"go-mysql-transfer/service/oracle/models"
	meta "go-mysql-transfer/service/oracle/oracle_meta"
	"go-mysql-transfer/service/oracle/positioner"
	"go-mysql-transfer/service/oracle/translator"
	"go-mysql-transfer/service/oracle/utils"
	"go-mysql-transfer/util/logs"
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
	tableController TableController
	processTracer   ProcessTracer
	instance        Instance
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
	var dbConfig map[database.DataSourceConfig]database.DataSource
	sourceConfig := database.NewDataSourceConfig(c.configuration)
	dbConfig[*sourceConfig] = database.DataSource{}

	// connect to all databases
	c.dataSourceFactory.SetDataSources(dbConfig)
	c.dataSourceFactory.Start() // set all database connection

	c.SetRunMode(models.RunMode(c.configuration.OracleRunMode))
	c.SetSourceDBType(database.DBType(c.configuration.Flavor))
	// todo, maybe it is useless
	c.SetTargetDBType(database.DBType(c.configuration.Flavor))
	c.globalContext = c.InitGlobalContext()
	concurrent := c.configuration.OracleTableConcurrentEnable
	tableMetas := c.InitTables()
	// todo add alarm
	// todo add extractorDump, statBufferSize, statePrintInterval
	// set default temporary
	// for test
	statBufferSize := 16384
	statPrintInterval := 5
	alarmReceiver := "spiswe@outlook.com"
	retryTimes := 3
	// milliseconds
	retryInterval := 1000
	noupdateThresoldDefault := -1
	noUpdateThresold := 0
	useExtractorExecutor := true
	useApplierExecutor := true
	threadSize := 1
	if concurrent {
		threadSize = c.configuration.OracleTableConcurrentSize
	}
	c.tableController = *NewTableController(len(tableMetas), threadSize)
	c.processTracer = *NewProcessTracer(c.runMode, len(tableMetas))
	if threadSize < len(tableMetas) {
		noupdateThresoldDefault = 3
	}
	if useExtractorExecutor {
		extractorSize := 5
		extractorExecutor, _ := NewExecutorPool(uint64(extractorSize))
	}

	if useApplierExecutor {
		applierSize := 5
		applierExecutor, _ := NewExecutorPool(uint64(applierSize))
	}

	for _, tableHolder := range tableMetas {
		oracleContext := c.BuildContext(c.globalContext, tableHolder.table, tableHolder.ignoreSchema)
		positioner := c.ChoosePositioner(tableHolder)
		//extractor :=
	}

}

//func (c *Controller) TableHolder(table meta.Table) {}

// InitTables todo 可能存在大小写问题，待测试确认
func (c *Controller) InitTables() []TableHolder {
	logs.Info("check table privileges ...")
	tableWhiteList := c.configuration.OracleTableWhiteList
	tableBlackList := c.configuration.OracleTableBlackList
	isEmpty := true
	for _, table := range tableWhiteList {
		if len(table) == 0 {
			isEmpty = false
			break
		}
	}
	var tables []TableHolder
	// target DBType is not important
	if !isEmpty {
		for _, obj := range tableWhiteList {
			whiteTable := c.getTable(obj)
			if !c.isBlackTable(whiteTable, tableBlackList) {
				tableInfo := strings.Split(whiteTable, ".")
				ignoreSchema := false
				var whiteTables []meta.Table
				if len(tableInfo) == 1 {
					whiteTables = meta.TableMetaGeneratorController.
						GetTableMetasWithoutColumn(c.globalContext.SourceDS(), "", tableInfo[1])
					ignoreSchema = true
				} else if len(tableInfo) == 2 {
					whiteTables = meta.TableMetaGeneratorController.
						GetTableMetasWithoutColumn(c.globalContext.SourceDS(), tableInfo[0], tableInfo[1])
				} else {
					logs.Error("table " + whiteTable + " is not valid")
				}

				if len(whiteTables) == 0 {
					logs.Error("table " + whiteTable + " is not found")
				}

				for _, table := range whiteTables {
					if (!c.isBlackTable(table.Name(), tableBlackList)) &&
						(!c.isBlackTable(table.GetFullName(), tableBlackList)) {
						meta.TableMetaGeneratorController.BuildColumns(c.globalContext.SourceDS(), &table)
						// ext keys used to build drds sql
						// may be use translator to build ext key, later decide
						holder := NewTableHolder(table)
						holder.ignoreSchema = ignoreSchema
						if !utils.SliceContains(tables, holder) {
							tables = append(tables, *holder)
						}
					}
				}
			}
		}
	} else {
		tablemetas := meta.TableMetaGeneratorController.GetTableMetasWithoutColumn(c.globalContext.SourceDS(), "", "")
		for _, tablemeta := range tablemetas {
			if (!c.isBlackTable(tablemeta.Name(), tableBlackList)) && (!c.isBlackTable(tablemeta.GetFullName(), tableBlackList)) {
				meta.TableMetaGeneratorController.BuildColumns(c.globalContext.SourceDS(), &tablemeta)
				holder := NewTableHolder(tablemeta)
				if !utils.SliceContains(tables, holder) {
					tables = append(tables, *holder)
				}
			}
		}
	}
	logs.Info("checks done, all tables ok")
	return tables
}

//
//func (c *Controller) BuildExtKeys(table meta.Table, tableStr string, targetDB database.DBType) DataTranslator {
//	return DataTranslator().translator
//}

func (c *Controller) ChooseExtractor(
	tableHolder TableHolder,
	context models.OracleContext,
	runMode models.RunMode,
	positioner positioner.RecordPositioner) extractor.RecordExtractor {

	once := c.configuration.OracleExtractorOnce

	if c.sourceDBType == database.Oracle {
		if runMode == models.FULL || runMode == models.CHECK {
			tableName := tableHolder.table.Name()
			fullName := tableHolder.table.GetFullName()
			extractSQL := c.configuration.OracleExtractorSQL

		}
	}

}

func (c *Controller) ChoosePositioner(tableHolder TableHolder) positioner.RecordPositioner {
	positionerMode := c.configuration.OraclePositionerMode
	switch positionerMode {
	// todo need to finish file/zk
	case "memory":
		p := positioner.MemoryRecordPositioner{}
		return p
	case "file":
		p := positioner.FileRecordPositioner{}
		return p
	case "zk":
		p := positioner.ZKRecordPositioner{}
		return p
	default:
		p := positioner.MemoryRecordPositioner{}
		return p
	}
}

func (c *Controller) BuildContext(
	context models.OracleContext,
	table meta.Table,
	ignoreSchema bool) models.OracleContext {
	result := context.CloneOracleContext()
	result.SetTableMeta(table)
	// 识别是否存在schema 定义
	if ignoreSchema {
		result.SetIgnoreSchema(ignoreSchema)
	}

	return *result
}

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
