// Package storage /*
package storage

import (
	"github.com/siddontang/go-mysql/mysql"

	"go-mysql-transfer/global"
)

type PositionStorage interface {
	Initialize() error
	Save(pos mysql.Position) error
	Get() (mysql.Position, error)
}

func NewPositionStorage() PositionStorage {
	if global.Cfg().IsCluster() {
		if global.Cfg().IsZk() {
			return &zkPositionStorage{}
		}
		if global.Cfg().IsEtcd() {
			return &etcdPositionStorage{}
		}
	}

	return &boltPositionStorage{}
}
