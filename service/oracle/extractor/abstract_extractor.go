package extractor

import (
	"go-mysql-transfer/service/oracle"
	"go-mysql-transfer/service/oracle/models"
	"go-mysql-transfer/service/oracle/positioner"
)

type RecordExtractor interface {
	oracle.LifeCycle
	Extract() []models.Record
	Status() models.ExtractStatus
	Ack(records []models.Record) positioner.OraclePosition
}

type AbstractRecordExtractor struct {
	absLifeCycle oracle.AbstractLifeCycle
	status       models.ExtractStatus
	tracer       oracle.ProcessTracer
}

func (a *AbstractRecordExtractor) AbsLifeCycle() oracle.AbstractLifeCycle {
	return a.absLifeCycle
}

func (a *AbstractRecordExtractor) SetAbsLifeCycle(absLifeCycle oracle.AbstractLifeCycle) {
	a.absLifeCycle = absLifeCycle
}

func (a *AbstractRecordExtractor) SetStatus(status models.ExtractStatus) {
	a.status = status
}

func (a *AbstractRecordExtractor) Tracer() oracle.ProcessTracer {
	return a.tracer
}

func (a *AbstractRecordExtractor) SetTracer(tracer oracle.ProcessTracer) {
	a.tracer = tracer
}

func (a *AbstractRecordExtractor) Start() {
	a.absLifeCycle.Start()
}

func (a *AbstractRecordExtractor) Stop() {
	a.absLifeCycle.Stop()
}

func (a *AbstractRecordExtractor) Abort(reason string, err error) {
	a.absLifeCycle.Abort(reason, err)
}

func (a *AbstractRecordExtractor) IsStart() bool {
	return a.absLifeCycle.IsStart()
}

func (a *AbstractRecordExtractor) IsStop() bool {
	return a.absLifeCycle.IsStop()
}

// Extract different extractor use different extract method
func (a *AbstractRecordExtractor) Extract() []models.Record {
	//TODO implement me
	panic("implement me")
}

func (a *AbstractRecordExtractor) Status() models.ExtractStatus {
	return a.status
}

// Ack different extractor use different ack method
func (a *AbstractRecordExtractor) Ack(records []models.Record) positioner.OraclePosition {
	//TODO implement me
	panic("implement me")
}
