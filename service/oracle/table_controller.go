package oracle

import (
	"context"
	"golang.org/x/sync/semaphore"
	"sync"
)

type TableController struct {
	tableWg *sync.WaitGroup
	queue   chan Instance // 任务
	sem     *semaphore.Weighted
}

func NewTableController(total int, concurrent int) *TableController {
	tableWg := sync.WaitGroup{}
	tableWg.Add(total)
	sem := semaphore.NewWeighted(int64(concurrent))
	return &TableController{
		tableWg: &tableWg,
		sem:     sem,
	}
}

func (t *TableController) Acquire() {
	_ = t.sem.Acquire(context.Background(), 1)
}

func (t *TableController) Release(instance Instance) {
	t.sem.Release(1)
	t.queue <- instance
	t.tableWg.Done()
}

func (t *TableController) TakeDone() Instance {
	return <-t.queue
}

func (t *TableController) WaitForDone() {
	t.tableWg.Wait()
}
