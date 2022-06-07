package model

import (
	"go-mysql-transfer/model/position"
	meta "go-mysql-transfer/oracle_meta"
)

type OracleContext struct {
	// 最后一次同步的position 记录
	lastPosition position.AbstractPosition
	// 对应table oracle_meta
	tableMeta meta.Table
	// 设置是否忽略schema
	ignoreSchema bool
	// 运行模式
	runMode RunMode
}
