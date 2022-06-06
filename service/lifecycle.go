package service

import "go-mysql-transfer/util/logs"

type LifeCycle interface {
	start()
	stop()
	abort(reason string, err error)
	isStart() bool
	isStop() bool
}

type AbstractLifeCycle struct {
	Running bool `default:"false" json:"running,omitempty"`
}

func (alc *AbstractLifeCycle) isStart() bool {
	return alc.Running
}

func (alc *AbstractLifeCycle) isStop() bool {
	return !alc.isStart()
}

func (alc *AbstractLifeCycle) start() {
	if alc.Running {
		return
	} else {
		alc.Running = true
	}
}

func (alc *AbstractLifeCycle) stop() {
	if !alc.Running {
		return
	} else {
		alc.Running = false
	}
}

func (alc *AbstractLifeCycle) abort(reason string, err error) {
	logs.Error("abort caused by" + reason + err.Error())
	alc.stop()
}
