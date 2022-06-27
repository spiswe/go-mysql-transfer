package models

import (
	"go-mysql-transfer/service/oracle/database"
	"go-mysql-transfer/service/oracle/models/position"
	"go-mysql-transfer/service/oracle/oracle_meta"
)

type OracleContext struct {
	lastPosition         position.AbstractPosition // 最后一次同步的position 记录
	tableMeta            meta.Table                // 对应table oracle_meta
	ignoreSchema         bool                      // 设置是否忽略schema
	runMode              RunMode                   // 运行模式
	onceCrawNum          int                       // 单次提取记录数量
	tpsLimit             int                       `default:"0"` // 0 及以下代表不限制
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
		tpsLimit:             TPSLimit,
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
		tpsLimit:             TPSLimit,
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

func (o *OracleContext) SetTablePKs(tablePKs map[string][]string) {
	o.tablePKs = tablePKs
}

func (o *OracleContext) TablePKs() map[string][]string {
	return o.tablePKs
}

func (o *OracleContext) MViewLogType() string {
	return o.mViewLogType
}

func (o *OracleContext) SetLastPosition(position position.AbstractPosition) {
	o.lastPosition = position
}

func (o *OracleContext) OnceCrawNum() int {
	return o.onceCrawNum
}

func (o *OracleContext) SetOnceCrawNum(onceCrawNum int) {
	o.onceCrawNum = onceCrawNum
}

func (o *OracleContext) TargetDS() database.DataSource {
	return o.targetDS
}

func (o *OracleContext) SourceDS() database.DataSource {
	return o.sourceDS
}

func (o *OracleContext) SetTargetDS(targetDS database.DataSource) {
	o.targetDS = targetDS
}

func (o *OracleContext) SetSourceDS(SourceDS database.DataSource) {
	o.sourceDS = SourceDS
}

func (o *OracleContext) IsBatchApply() bool {
	return o.batchApply
}

func (o *OracleContext) SetBatchApply(batchApply bool) {
	o.batchApply = batchApply
}

func (o *OracleContext) GetSourceCodeEncoding() string {
	return o.SourceCodeEncoding
}

func (o *OracleContext) SetSourceCodeEncoding(sourceCodeEncoding string) {
	o.SourceCodeEncoding = sourceCodeEncoding
}

func (o *OracleContext) SetTargetCodeEncoding(targetCodeEncoding string) {
	o.TargetCodeEncoding = targetCodeEncoding
}

func (o *OracleContext) GetTargetCodeEncoding() string {
	return o.TargetCodeEncoding
}

func (o *OracleContext) TableMeta() meta.Table {
	return o.tableMeta
}

func (o *OracleContext) SetTableMeta(tableMeta meta.Table) {
	o.tableMeta = tableMeta
}

func (o *OracleContext) GetTPSLimit() int {
	return o.tpsLimit
}

func (o *OracleContext) SetTPSLimit(TPSLimit int) {
	o.tpsLimit = TPSLimit
}

func (o *OracleContext) GetRunMode() RunMode {
	return o.runMode
}

func (o *OracleContext) SetRunMode(mode RunMode) {
	o.runMode = mode
}

func (o *OracleContext) IsIgnoreSchema() bool {
	return o.ignoreSchema
}

func (o *OracleContext) SetIgnoreSchema(ignoreSchema bool) {
	o.ignoreSchema = ignoreSchema
}

func (o *OracleContext) SetSkipApplierException(skipApplierException bool) {
	o.skipApplierException = skipApplierException
}

func (o *OracleContext) IsSkipApplierException() bool {
	return o.skipApplierException
}

func (o *OracleContext) CloneOracleContext() *OracleContext {
	return NewOracleContext(
		o.IsIgnoreSchema(),
		o.GetRunMode(),
		o.OnceCrawNum(),
		o.GetTPSLimit(),
		o.SourceDS(),
		o.TargetDS(),
		o.IsBatchApply(),
		o.IsSkipApplierException(),
		o.GetSourceCodeEncoding(),
		o.GetTargetCodeEncoding(),
		o.MViewLogType(),
		o.TablePKs(),
	)
}
