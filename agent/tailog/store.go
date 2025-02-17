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

// Package tailog help to store tail logs
package tailog

import (
	"container/ring"
	"strings"
	"sync"
)

// Store implement ring save logs.
type Store struct {
	log      *ring.Ring
	capacity int
	m        sync.Mutex
}

// NewStore creates Store.
func NewStore(capacity int) *Store {
	return &Store{
		log:      ring.New(capacity),
		capacity: capacity,
	}
}

// Write writes log for store.
func (l *Store) Write(b []byte) (int, error) {
	l.m.Lock()
	defer l.m.Unlock()

	if l.capacity == 0 {
		return len(b), nil
	}

	l.log.Value = string(b)
	l.log = l.log.Next()
	return len(b), nil
}

// Resize to update capacity.
func (l *Store) Resize(capacity int) {
	l.m.Lock()
	defer l.m.Unlock()

	if l.capacity == capacity {
		return
	}

	old := l.log

	l.log = ring.New(capacity)
	l.capacity = capacity
	if l.capacity == 0 {
		return
	}

	old.Do(func(p interface{}) {
		if p != nil {
			l.log.Value = p
			l.log = l.log.Next()
		}
	})
}

// GetLogs return all logs.
func (l *Store) GetLogs() []string {
	l.m.Lock()
	defer l.m.Unlock()

	if l.capacity == 0 {
		return nil
	}

	logs := make([]string, 0, l.capacity)

	replacer := getColorReplacer()
	l.log.Do(func(p interface{}) {
		if p != nil {
			logs = append(logs, replacer.Replace(p.(string)))
		}
	})

	return logs
}

func getColorReplacer() *strings.Replacer {
	const (
		red    = "\x1b[31m"
		yellow = "\x1b[33m"
		blue   = "\x1b[36m"
		gray   = "\x1b[37m"
		end    = "\x1b[0m"
	)

	return strings.NewReplacer(red, "", yellow, "", blue, "", gray, "", end, "")
}
