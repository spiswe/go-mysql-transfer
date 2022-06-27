package database

import (
	"database/sql"
	"go-mysql-transfer/global"
)

type DataSource struct {
	conn *sql.DB
}

func (d *DataSource) Conn() *sql.DB {
	return d.conn
}

func (d *DataSource) SetConn(conn *sql.DB) {
	d.conn = conn
}

func NewDataSource(conn *sql.DB) *DataSource {
	return &DataSource{conn: conn}
}

type DataSourceConfig struct {
	server         string
	userName       string
	password       string
	service        string // oracle service name
	port           int
	dbType         DBType
	sourceEncoding string
	config         *global.Config
}

func NewDataSourceConfig(cfg *global.Config) *DataSourceConfig {
	return &DataSourceConfig{
		server:         cfg.OracleAddr,
		userName:       cfg.OracleUser,
		password:       cfg.OraclePassword,
		service:        cfg.OracleService,
		port:           cfg.OraclePort,
		dbType:         DBType(cfg.Flavor),
		sourceEncoding: cfg.OracleEncoding,
		config:         cfg}
}

func (d *DataSourceConfig) Server() string {
	return d.server
}

func (d *DataSourceConfig) SetServer(server string) {
	d.server = server
}

func (d *DataSourceConfig) UserName() string {
	return d.userName
}

func (d *DataSourceConfig) SetUserName(userName string) {
	d.userName = userName
}

func (d *DataSourceConfig) Password() string {
	return d.password
}

func (d *DataSourceConfig) SetPassword(password string) {
	d.password = password
}

func (d *DataSourceConfig) Service() string {
	return d.service
}

func (d *DataSourceConfig) SetService(service string) {
	d.service = service
}

func (d *DataSourceConfig) Port() int {
	return d.port
}

func (d *DataSourceConfig) SetPort(port int) {
	d.port = port
}

func (d *DataSourceConfig) DbType() DBType {
	return d.dbType
}

func (d *DataSourceConfig) SetDbType(dbType DBType) {
	d.dbType = dbType
}

func (d *DataSourceConfig) SourceEncoding() string {
	return d.sourceEncoding
}

func (d *DataSourceConfig) SetSourceEncoding(sourceEncoding string) {
	d.sourceEncoding = sourceEncoding
}

func (d *DataSourceConfig) Config() *global.Config {
	return d.config
}

func (d *DataSourceConfig) SetConfig(config *global.Config) {
	d.config = config
}

func NewDataSourceConfigFull(
	server string,
	userName string,
	password string,
	service string,
	port int,
	dbType DBType,
	sourceEncoding string,
	config *global.Config) *DataSourceConfig {
	return &DataSourceConfig{
		server:         server,
		userName:       userName,
		password:       password,
		service:        service,
		port:           port,
		dbType:         dbType,
		sourceEncoding: sourceEncoding,
		config:         config}
}
