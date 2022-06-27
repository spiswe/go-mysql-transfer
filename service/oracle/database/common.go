package database

import (
	"database/sql"
	go_ora "github.com/sijms/go-ora/v2"
	"go-mysql-transfer/global"
	"go-mysql-transfer/util/logs"
	"os"
)

type DBType string

const (
	MySQL DBType = "mysql"

	Oracle DBType = "oracle"

	PostgreSQL DBType = "postgres"
)

type DataSourceFactory struct {
	dataSources map[DataSourceConfig]DataSource
}

func (dsf *DataSourceFactory) DataSources() map[DataSourceConfig]DataSource {
	return dsf.dataSources
}

func (dsf *DataSourceFactory) GetDataSource(config DataSourceConfig) DataSource {
	return dsf.dataSources[config]
}

func (dsf *DataSourceFactory) SetDataSources(dataSources map[DataSourceConfig]DataSource) {
	dsf.dataSources = dataSources
}

func (dsf *DataSourceFactory) Start() {
	//dsf.SetDataSource()
	//dsf.SetDataSource(dsf.DataSource())
	for k, _ := range dsf.DataSources() {
		datasource := dsf.createDataSourceByCfg(*k.Config())
		dsf.dataSources[k] = *datasource
	}
}

func (dsf *DataSourceFactory) createDataSourceByCfg(cfg global.Config) *DataSource {
	// for test
	server := cfg.OracleAddr
	port := cfg.OraclePort
	service := cfg.OracleService
	user := cfg.OracleUser
	password := cfg.OraclePassword

	if cfg.Flavor != string(Oracle) {
		panic("not support yet")
	}

	url := go_ora.BuildUrl(server, port, service, user, password, nil)

	conn, err := sql.Open(cfg.Flavor, url)
	if err != nil {
		logs.Error("Can't open the driver: " + err.Error())
		//logs.Error("open database connection error: " + err.Error())
		os.Exit(-1)
	}

	err = conn.Ping()
	// todo set conn config
	if err != nil {
		logs.Error("Can't ping connection: " + err.Error())
		os.Exit(-1)
	}

	return NewDataSource(conn)
}

func (dsf *DataSourceFactory) createDataSourceByDSN() {}
