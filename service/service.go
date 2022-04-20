// Package service /*
package service

import (
	"go-mysql-transfer/global"
	"go-mysql-transfer/service/election"
)

var (
	_transferService *TransferService
	_electionService election.Service
	_clusterService  *ClusterService
)

func Initialize() error {
	transferService := &TransferService{
		loopStopSignal: make(chan struct{}, 1),
	}
	err := transferService.initialize()
	if err != nil {
		return err
	}
	_transferService = transferService

	if global.Cfg().IsCluster() {
		_clusterService = &ClusterService{
			electionSignal: make(chan bool, 1),
		}
		_electionService = election.NewElection(_clusterService.electionSignal)
	}

	return nil
}

func StartUp() {
	if global.Cfg().IsCluster() {
		_clusterService.boot()
	} else {
		_transferService.StartUp()
	}
}

func Close() {
	_transferService.Close()
}

func TransferServiceIns() *TransferService {
	return _transferService
}

func ClusterServiceIns() *ClusterService {
	return _clusterService
}
