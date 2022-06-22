package oracle

import (
	"go-mysql-transfer/service/oracle/models"
)

type Instance struct {
	instanceContext models.OracleContext
	extractor       RecordExtractor
	applier         RecordApplier
	translator      DataTranslator
}
