package database

type DBType int

const (
	MySQL DBType = iota

	Oracle

	PostgreSQL
)

type DataSourceFactory struct {
}
