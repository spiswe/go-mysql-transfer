package oracle

import (
	"fmt"
	"go-mysql-transfer/service/oracle/models"
	"go-mysql-transfer/util/logs"
	"strings"
)

var TracerController = &ProcessTracer{}

const (
	FullFormat  = "{Preparing: %d, fulling: %d, success: %d, failed: %d}"
	IncFormat   = "{Preparing:%d, incing: %d, success: %d, failed: %d}"
	CheckFormat = "{Preparing:%d, checking: %d, success:%d, failed:%d}"
	AllFormat   = "{Preparing:%d, fulling: %d, checking: %d, success: %d, failed: %d}"
)

type ProcessTracer struct {
	total  int
	mode   models.RunMode
	status map[string]models.ProgressStatus
}

func NewProcessTracer(mode models.RunMode, total int) *ProcessTracer {
	return &ProcessTracer{total: total, mode: mode}
}

func (pt *ProcessTracer) Update(tableName string, progress models.ProgressStatus) {
	if pt.status[tableName] != models.FAILED {
		pt.status[tableName] = progress
	}
}

func (pt *ProcessTracer) Print(detail bool) {
	var fulling, incing, failed, success = 0, 0, 0, 0
	var fullingTables, incingTables, failedTables, successTables []string
	var msg = ""

	for tableName, progress := range pt.status {
		if progress == models.FULLING {
			fulling++
			fullingTables = append(fullingTables, tableName)
		} else if progress == models.INCING {
			incing++
			incingTables = append(incingTables, tableName)
		} else if progress == models.FAILED {
			failed++
			failedTables = append(failedTables, tableName)
		} else if progress == models.SUCCESS {
			success++
			successTables = append(successTables, tableName)
		}

	}

	var unknown = pt.total - fulling - incing - failed - success

	if pt.mode == models.ALL {
		msg = fmt.Sprintf(AllFormat, unknown, fulling, incing, success, failed)
	} else if pt.mode == models.FULL {
		msg = fmt.Sprintf(FullFormat, unknown, fulling, success, failed)
	} else if pt.mode == models.INC {
		msg = fmt.Sprintf(IncFormat, unknown, incing, success, failed)
	} else if pt.mode == models.CHECK {
		msg = fmt.Sprintf(CheckFormat, unknown, fulling, success, failed)
	}

	logs.Info(msg)
	tableListString := ""
	if detail {
		if fulling > 0 {
			tableListString = strings.Join(fullingTables, ",")
			if pt.mode == models.CHECK {
				logs.Info("checking: " + tableListString)
			} else {
				logs.Info("fulling: " + tableListString)
			}
		}
		if incing > 0 {
			tableListString = strings.Join(incingTables, ",")
			logs.Info("incing: " + tableListString)
		}
		if failed > 0 {
			tableListString = strings.Join(failedTables, ",")
			logs.Info("failed: " + tableListString)
		}
		tableListString = strings.Join(successTables, ",")
		logs.Info("success: " + tableListString)
	}

}
