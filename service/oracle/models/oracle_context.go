package models

import (
	"go-mysql-transfer/service/oracle/models/position"
	"go-mysql-transfer/service/oracle/oracle_meta"
	"go-mysql-transfer/util/database"
)

type OracleContext struct {
	lastPosition         position.AbstractPosition // 最后一次同步的position 记录
	tableMeta            meta.Table                // 对应table oracle_meta
	ignoreSchema         bool                      // 设置是否忽略schema
	runMode              RunMode                   // 运行模式
	onceCrawNum          int                       // 单次提取记录数量
	TPSLimit             int                       `default:"0"` // 0 及以下代表不限制
	sourceDS             database.DataSource       // datasource
	targetDS             database.DataSource       // target datasource
	batchApply           bool                      // 是否批量
	skipApplierException bool                      // 是否跳过异常
	SourceCodeEncoding   string                    `default:"UTF-8"` //source encoding
	TargetCodeEncoding   string                    `default:"UTF-8"` // target encoding
	mViewLogType         string                    // 物化视图日志类型
	tablePKs             map[string][]string       // pk
}

func NewOracleContextFull(
	lastPosition position.AbstractPosition,
	tableMeta meta.Table,
	ignoreSchema bool,
	runMode RunMode,
	onceCrawNum int,
	TPSLimit int,
	sourceDS database.DataSource,
	targetDS database.DataSource,
	batchApply bool,
	skipApplierException bool,
	sourceCodeEncoding string,
	targetCodeEncoding string,
	mViewLogType string,
	tablePKs map[string][]string,
) (o *OracleContext) {
	return &OracleContext{
		lastPosition:         lastPosition,
		tableMeta:            tableMeta,
		ignoreSchema:         ignoreSchema,
		runMode:              runMode,
		onceCrawNum:          onceCrawNum,
		TPSLimit:             TPSLimit,
		sourceDS:             sourceDS,
		targetDS:             targetDS,
		batchApply:           batchApply,
		skipApplierException: skipApplierException,
		SourceCodeEncoding:   sourceCodeEncoding,
		TargetCodeEncoding:   targetCodeEncoding,
		mViewLogType:         mViewLogType,
		tablePKs:             tablePKs,
	}
}

func NewOracleContext(
	ignoreSchema bool,
	runMode RunMode,
	onceCrawNum int,
	TPSLimit int,
	sourceDS database.DataSource,
	targetDS database.DataSource,
	batchApply bool,
	skipApplierException bool,
	sourceCodeEncoding string,
	targetCodeEncoding string,
	mViewLogType string,
	tablePKs map[string][]string,
) (o *OracleContext) {
	return &OracleContext{
		ignoreSchema:         ignoreSchema,
		runMode:              runMode,
		onceCrawNum:          onceCrawNum,
		TPSLimit:             TPSLimit,
		sourceDS:             sourceDS,
		targetDS:             targetDS,
		batchApply:           batchApply,
		skipApplierException: skipApplierException,
		SourceCodeEncoding:   sourceCodeEncoding,
		TargetCodeEncoding:   targetCodeEncoding,
		mViewLogType:         mViewLogType,
		tablePKs:             tablePKs,
	}
}

func (o *OracleContext) setTablePKs(tablePKs map[string][]string) {
	o.tablePKs = tablePKs
}

func (o *OracleContext) getTablePKs() map[string][]string {
	return o.tablePKs
}

func (o *OracleContext) getMViewLogType() string {
	return o.mViewLogType
}

func (o *OracleContext) setLastPosition(position position.AbstractPosition) {
	o.lastPosition = position
}

func (o *OracleContext) getOnceCrawNum() int {
	return o.onceCrawNum
}

func (o *OracleContext) setOnceCrawNum(onceCrawNum int) {
	o.onceCrawNum = onceCrawNum
}

func (o *OracleContext) getTargetDS() database.DataSource {
	return o.targetDS
}

func (o *OracleContext) getSourceDS() database.DataSource {
	return o.sourceDS
}

func (o *OracleContext) setTargetDS(targetDS database.DataSource) {
	o.targetDS = targetDS
}

func (o *OracleContext) setSourceDS(SourceDS database.DataSource) {
	o.sourceDS = SourceDS
}

func (o *OracleContext) isBatchApply() bool {
	return o.batchApply
}

func (o *OracleContext) setBatchApply(batchApply bool) {
	o.batchApply = batchApply
}

func (o *OracleContext) getSourceCodeEncoding() string {
	return o.SourceCodeEncoding
}

func (o *OracleContext) setSourceCodeEncoding(sourceCodeEncoding string) {
	o.SourceCodeEncoding = sourceCodeEncoding
}

func (o *OracleContext) setTargetCodeEncoding(targetCodeEncoding string) {
	o.TargetCodeEncoding = targetCodeEncoding
}

func (o *OracleContext) getTargetCodeEncoding() string {
	return o.TargetCodeEncoding
}

func (o *OracleContext) getTableMeta() meta.Table {
	return o.tableMeta
}

func (o *OracleContext) setTableMeta(tableMeta meta.Table) {
	o.tableMeta = tableMeta
}

func (o *OracleContext) getTPSLimit() int {
	return o.TPSLimit
}

func (o *OracleContext) setTPSLimit(TPSLimit int) {
	o.TPSLimit = TPSLimit
}

func (o *OracleContext) getRunMode() RunMode {
	return o.runMode
}

func (o *OracleContext) setRunMode(mode RunMode) {
	o.runMode = mode
}

func (o *OracleContext) isIgnoreSchema() bool {
	return o.ignoreSchema
}

func (o *OracleContext) setIgnoreSchema(ignoreSchema bool) {
	o.ignoreSchema = ignoreSchema
}

func (o *OracleContext) setSkipApplierException(skipApplierException bool) {
	o.skipApplierException = skipApplierException
}

func (o *OracleContext) isSkipApplierException() bool {
	return o.skipApplierException
}

func (o *OracleContext) cloneOracleContext() *OracleContext {
	return NewOracleContext(
		o.isIgnoreSchema(),
		o.getRunMode(),
		o.getOnceCrawNum(),
		o.getTPSLimit(),
		o.getSourceDS(),
		o.getTargetDS(),
		o.isBatchApply(),
		o.isSkipApplierException(),
		o.getSourceCodeEncoding(),
		o.getTargetCodeEncoding(),
		o.getMViewLogType(),
		o.getTablePKs(),
	)
}
