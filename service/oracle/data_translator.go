package oracle

import (
	"go-mysql-transfer/service/oracle/models"
)

type DataTranslator interface {
	translatorSchema() string
	translatorTable() string
	translator(record models.Record)        // single record
	translatorBatch(record []models.Record) // batch record
}
