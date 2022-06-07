package model

type RunMode int

const (

	// MARK 增量记录
	MARK RunMode = iota

	// INC 增量
	INC

	// FULL 全量
	FULL

	// ALL FULL + INC
	ALL

	// CHECK 数据检查
	CHECK

	// CLEAR 清理物化视图
	CLEAR
)
