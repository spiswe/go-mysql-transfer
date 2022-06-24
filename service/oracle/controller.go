package oracle

import (
	"database/sql"
	go_ora "github.com/sijms/go-ora/v2"
	"go-mysql-transfer/service/oracle/database"
	"go-mysql-transfer/service/oracle/models"
	"go-mysql-transfer/util/logs"
)

type Controller struct {
	configuration     Configuration
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

func NewController(configuration Configuration) *Controller {
	return &Controller{configuration: configuration}
}

func (c *Controller) Configuration() Configuration {
	return c.configuration
}

func (c *Controller) SetConfiguration(configuration Configuration) {
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
//func (c *Controller) InitDataSource(dbType database.DBType) database.DataSource {
func (c *Controller) InitDataSource(dbType database.DBType) *sql.DB {
	// todo use config to add
	server := "192.168.2.120"
	port := 15122 // need int
	service := "LHR11G"
	user := "makang"
	password := "makang"
	oracleUrl := go_ora.BuildUrl(server, port, service, user, password, nil)
	db, err := sql.Open("oracle", oracleUrl)
	if err != nil {
		logs.Error("build database connection error " + err.Error())
	}
	return db
}
