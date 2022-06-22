package models

import (
	"go-mysql-transfer/service/oracle/oracle_meta"
)

type ColumnValue struct {
	column meta.ColumnMeta
	value  interface{}
	check  bool `default:"true"` // 是否需要做对比
}

func NewColumnValueWithCheck(column meta.ColumnMeta, value interface{}, check bool) *ColumnValue {
	return &ColumnValue{column: column, value: value, check: check}
}

func NewColumnValueWithoutCheck(column meta.ColumnMeta, value interface{}) *ColumnValue {
	return &ColumnValue{column: column, value: value}
}

func (cv *ColumnValue) Value() interface{} {
	return cv.value
}

func (cv *ColumnValue) SetValue(value interface{}) {
	cv.value = value
}

func (cv *ColumnValue) Column() meta.ColumnMeta {
	return cv.column
}

func (cv *ColumnValue) SetColumn(column meta.ColumnMeta) {
	cv.column = column
}

func (cv *ColumnValue) isCheck() bool {
	return cv.check
}

func (cv *ColumnValue) SetCheck(check bool) {
	cv.check = check
}

func (cv *ColumnValue) cloneColumValue() *ColumnValue {
	return NewColumnValueWithoutCheck(cv.Column(), cv.Value())
}
