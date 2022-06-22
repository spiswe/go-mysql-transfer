package models

type ExtractStatus int

const (

	// NORMAL 正常提取
	NORMAL ExtractStatus = iota

	// CATCHUP 追赶
	CATCHUP

	// NOUPDATE 无更新
	NOUPDATE

	// TABLEEND 结束提取
	TABLEEND
)
