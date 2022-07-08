package positioner

import "go-mysql-transfer/service/oracle"

type ZKRecordPositioner struct {
	oa       oracle.AbstractLifeCycle
	position OraclePosition
}

func (Z ZKRecordPositioner) GetLatest() OraclePosition {
	return Z.position
}

func (Z ZKRecordPositioner) Persist(position OraclePosition) {
	Z.position = position
}
