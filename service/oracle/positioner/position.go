package positioner

import "go-mysql-transfer/service/oracle/models"

//type OraclePosition struct {
//	position string
//}

type AbstractPosition struct {
}

type OraclePosition interface {
	ProgressHistory() []models.ProgressStatus
	CurrentProgress() models.ProgressStatus
}
