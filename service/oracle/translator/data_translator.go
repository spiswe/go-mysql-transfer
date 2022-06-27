package translator

import "go-mysql-transfer/service/oracle/models"

type DataTranslator interface {
	TranslatorSchema() string
	TranslatorTable() string
	TranslatorRecord(record models.Record) bool
	TranslatorRecordList(records []models.Record) []models.Record
}
