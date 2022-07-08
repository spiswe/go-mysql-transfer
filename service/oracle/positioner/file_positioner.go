package positioner

import "go-mysql-transfer/service/oracle"

type FileRecordPositioner struct {
	oa       oracle.AbstractLifeCycle
	position OraclePosition
}

func (f FileRecordPositioner) GetLatest() OraclePosition {
	return f.position
}

func (f FileRecordPositioner) Persist(position OraclePosition) {
	f.position = position
}
