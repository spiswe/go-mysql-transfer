package oracle

import "go-mysql-transfer/util/logs"

type LifeCycle interface {
	Start()
	Stop()
	Abort(reason string, err error)
	IsStart() bool
	IsStop() bool
}

type AbstractLifeCycle struct {
	Running bool `default:"false" json:"running,omitempty"`
}

func (alc *AbstractLifeCycle) Start() {
	if alc.Running {
		return
	} else {
		alc.Running = true
	}
}

func (alc *AbstractLifeCycle) Stop() {
	if !alc.Running {
		return
	} else {
		alc.Running = false
	}
}

func (alc *AbstractLifeCycle) Abort(reason string, err error) {
	logs.Error("abort caused by" + reason + err.Error())
	alc.Stop()
}

func (alc *AbstractLifeCycle) IsStart() bool {
	return alc.Running
}

func (alc *AbstractLifeCycle) IsStop() bool {
	return !alc.IsStart()
}
