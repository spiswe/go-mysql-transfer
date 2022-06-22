package oracle

import (
	"go-mysql-transfer/service/oracle/models"
	"go-mysql-transfer/service/oracle/models/position"
)

type RecordExtractor interface {
	start()
	stop()
	abort(reason string, err error)
	isStart() bool
	isStop() bool

	extract() []models.Record
	status() models.ExtractStatus
	ack(records []models.Record) position.OraclePosition
}

type AbstractRecordExtractor struct {
	status models.ExtractStatus
	tracer ProcessTracer
}

func (are *AbstractRecordExtractor) setStatus(status models.ExtractStatus) {
	are.status = status
}

func (are *AbstractRecordExtractor) Status() models.ExtractStatus {
	return are.status
}
