package oracle

import (
	extractor "go-mysql-transfer/service/oracle/extractor"
	"go-mysql-transfer/service/oracle/models"
)

type Instance struct {
	instanceContext models.OracleContext
	extractor       extractor.RecordExtractor
	applier         RecordApplier
	translator      DataTranslator
}
