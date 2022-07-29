package positioner

import (
	"go-mysql-transfer/service/oracle/models"
	"go-mysql-transfer/service/oracle/utils"
)

type IDPosition struct {
	position        string
	id              int
	progressHistory []models.ProgressStatus
	currentProgress models.ProgressStatus
}

//func (i *IDPosition) ProgressHistory() {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (i *IDPosition) CurrentProgress() {
//	//TODO implement me
//	panic("implement me")
//}

func (i *IDPosition) ID() int { return i.id }

func (i *IDPosition) SetID(id int) { i.id = id }

func (i *IDPosition) ProgressHistory() []models.ProgressStatus { return i.progressHistory }

func (i *IDPosition) CurrentProgress() models.ProgressStatus { return i.currentProgress }

func (i *IDPosition) SetProgressHistory(progressHistory []models.ProgressStatus) {
	i.progressHistory = progressHistory
}

func (i *IDPosition) SetCurrentProgress(currentProgress models.ProgressStatus) {
	if i.currentProgress != currentProgress && i.currentProgress != models.UNKNOW {
		i.progressHistory = append(i.progressHistory, currentProgress)
	}

	i.currentProgress = currentProgress
}

func (i *IDPosition) IsInHistory(progress models.ProgressStatus) bool {
	return utils.SliceContains(i.progressHistory, progress)
}
