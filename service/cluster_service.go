// Package service /*
package service

import (
	"log"

	"go-mysql-transfer/global"
	"go-mysql-transfer/metrics"
)

type ClusterService struct {
	electionSignal chan bool //选举信号
}

func (s *ClusterService) boot() error {
	log.Println("start master election")
	err := _electionService.Elect()
	if err != nil {
		return err
	}

	s.startElectListener()

	return nil
}

func (s *ClusterService) startElectListener() {
	go func() {
		for {
			select {
			case selected := <-s.electionSignal:
				global.SetLeaderNode(_electionService.Leader())
				global.SetLeaderFlag(selected)
				if selected {
					metrics.SetLeaderState(metrics.LeaderState)
					_transferService.StartUp()
				} else {
					metrics.SetLeaderState(metrics.FollowerState)
					_transferService.stopDump()
				}
			}
		}

	}()
}

func (s *ClusterService) Nodes() []string {
	return _electionService.Nodes()
}
