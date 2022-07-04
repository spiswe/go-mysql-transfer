package oracle

import (
	"golang.org/x/sync/semaphore"
	"sync"
)

type TableController struct {
	tableWg sync.WaitGroup
	work    chan func() // 任务
	sem     semaphore.Weighted
}
