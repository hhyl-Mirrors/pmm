// pmm-agent
// Copyright 2019 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connectionuptime

import (
	"sync"
	"time"
)

// Service calculates connection uptime between agent and server
type Service struct {
	mx sync.Mutex

	connectedStatuses []bool

	windowPeriodSeconds int
	indexLastStatus     int
	startTime           time.Time
	lastStatusTimestamp time.Time
}

// NewService creates new instance of Service
func NewService(windowPeriod time.Duration) *Service {
	return &Service{
		windowPeriodSeconds: int(windowPeriod.Seconds()),
		connectedStatuses:   make([]bool, int(windowPeriod.Seconds())),
	}
}

// RegisterConnectionStatus adds new connection status
func (s *Service) RegisterConnectionStatus(timestamp time.Time, connected bool) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.registerConnectionStatus(timestamp, connected)
}

func (s *Service) registerConnectionStatus(timestamp time.Time, connected bool) {
	if s.startTime.IsZero() {
		s.startTime = timestamp
		s.lastStatusTimestamp = timestamp
		s.connectedStatuses[0] = connected
		s.indexLastStatus = 0

		return
	}

	secondsFromLastEvent := int(timestamp.Unix() - s.lastStatusTimestamp.Unix())
	for i := s.indexLastStatus + 1; i < (s.indexLastStatus + secondsFromLastEvent); i++ {
		// set the same status to elements of previous connection status
		s.connectedStatuses[i%s.windowPeriodSeconds] = s.connectedStatuses[s.indexLastStatus]
	}

	s.indexLastStatus = (s.indexLastStatus + secondsFromLastEvent) % s.windowPeriodSeconds
	s.connectedStatuses[s.indexLastStatus] = connected
	s.lastStatusTimestamp = timestamp
}

// GetConnectedUpTimeSince calculates connected uptime between agent and server
// based on the connection statuses
func (s *Service) GetConnectedUpTimeSince(toTime time.Time) float32 {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.fillStatusesUntil(toTime)
	return s.calculateConnectionUpTime(toTime)
}

func (s *Service) calculateConnectionUpTime(toTime time.Time) float32 {
	totalNumOfSeconds := s.getTotalNumberOfSeconds(toTime)
	startIndex := s.getStartIndex(totalNumOfSeconds)
	connectedSeconds := s.getNumOfConnectedSeconds(startIndex, totalNumOfSeconds)

	return float32(connectedSeconds) / float32(totalNumOfSeconds) * 100
}

func (s *Service) getTotalNumberOfSeconds(toTime time.Time) int {
	totalNumOfSeconds := s.windowPeriodSeconds
	diffInSecondsBetweenStartTimeAndToTime := int(toTime.Unix() - s.startTime.Unix())
	if diffInSecondsBetweenStartTimeAndToTime < s.windowPeriodSeconds {
		totalNumOfSeconds = diffInSecondsBetweenStartTimeAndToTime
	}
	return totalNumOfSeconds
}

func (s *Service) getStartIndex(size int) int {
	startElement := s.indexLastStatus - size
	if startElement < 0 {
		startElement = s.windowPeriodSeconds + startElement
	}
	return startElement
}

func (s *Service) getNumOfConnectedSeconds(startIndex int, totalNumOfSeconds int) int {
	endIndex := startIndex + totalNumOfSeconds
	connectedSeconds := 0
	for i := startIndex; i < endIndex; i++ {
		if s.connectedStatuses[i%s.windowPeriodSeconds] {
			connectedSeconds++
		}
	}
	return connectedSeconds
}

// fill values in the slice until toTime
func (s *Service) fillStatusesUntil(toTime time.Time) {
	s.registerConnectionStatus(toTime, s.connectedStatuses[s.indexLastStatus])
}
