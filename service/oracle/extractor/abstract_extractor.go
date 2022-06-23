package extractor

import (
	"go-mysql-transfer/service/oracle"
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
	absLifeCycle  oracle.AbstractLifeCycle
	extractStatus models.ExtractStatus
	tracer        oracle.ProcessTracer
}

func (a *AbstractRecordExtractor) AbsLifeCycle() oracle.AbstractLifeCycle {
	return a.absLifeCycle
}

func (a *AbstractRecordExtractor) SetAbsLifeCycle(absLifeCycle oracle.AbstractLifeCycle) {
	a.absLifeCycle = absLifeCycle
}

func (are *AbstractRecordExtractor) extract() []models.Record {
	//TODO implement me
	panic("implement me")
}

func (are *AbstractRecordExtractor) status() models.ExtractStatus {
	//TODO implement me
	panic("implement me")
}

func (are *AbstractRecordExtractor) ack(records []models.Record) position.OraclePosition {
	//TODO implement me
	panic("implement me")
}

func (are *AbstractRecordExtractor) Tracer() oracle.ProcessTracer {
	return are.tracer
}

func (are *AbstractRecordExtractor) SetTracer(tracer oracle.ProcessTracer) {
	are.tracer = tracer
}

func (are *AbstractRecordExtractor) setStatus(status models.ExtractStatus) {
	are.extractStatus = status
}

func (are *AbstractRecordExtractor) Status() models.ExtractStatus {
	return are.extractStatus
}
