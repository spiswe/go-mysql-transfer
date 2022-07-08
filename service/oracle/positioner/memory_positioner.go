package positioner

import "go-mysql-transfer/service/oracle"

type MemoryRecordPositioner struct {
	oa       oracle.AbstractLifeCycle
	position OraclePosition
}

func (m MemoryRecordPositioner) GetLatest() OraclePosition {
	return m.position
}

func (m MemoryRecordPositioner) Persist(position OraclePosition) {
	m.position = position
}
