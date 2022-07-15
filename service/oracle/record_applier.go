package oracle

type RecordApplier interface {
	start()
	stop()
	abort(reason string, err error)
	isStart() bool
	isStop() bool
}
