package oracle

import (
	"go-mysql-transfer/service/oracle/database"
	extractor "go-mysql-transfer/service/oracle/extractor"
	"go-mysql-transfer/service/oracle/models"
	"go-mysql-transfer/service/oracle/positioner"
)

type Instance struct {
	abs             AbstractLifeCycle
	instanceContext models.OracleContext
	extractor       extractor.RecordExtractor
	applier         RecordApplier
	translator      DataTranslator
	positioner      positioner.RecordPositioner
	//todo
	//alarmService, alarmReceiver
	tableController TableController
	progressTracer  ProcessTracer
	//todo aggregation
	dbType       database.DBType
	tableShitKey string
	executor     Pool
}

func NewInstance(context models.OracleContext) *Instance {
	tableMeta := context.TableMeta()
	return &Instance{
		instanceContext: context,
		extractor:       nil,
		applier:         nil,
		translator:      nil,
		positioner:      nil,
		tableController: TableController{},
		progressTracer:  ProcessTracer{},
		dbType:          "",
		tableShitKey:    tableMeta.GetFullName(),
	}
}

func (i *Instance) start() {
	i.abs.Start()
	i.tableController.Acquire()
	tableMeta := i.instanceContext.TableMeta()
	executorName := "Instance-" + tableMeta.GetFullName()
	//if i.executor. == (Pool{}){

	}
}
