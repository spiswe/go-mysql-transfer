package models

type ProgressStatus int

const (

	// UNKNOW know nothing
	UNKNOW ProgressStatus = iota

	// MARKING 标记ing
	MARKING

	// FULLING 全量ing
	FULLING

	// INCING 增量ing
	INCING

	// CLEARING 清理ing
	CLEARING

	// FAILED 清理物化视图
	FAILED

	// SUCCESS 成功
	SUCCESS
)
