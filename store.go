/*
 * Copyright 2023 veerdone
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gsecurity

import (
	"sync"
	"time"
)

type Store interface {
	Set(key string, val interface{}, exTime int64)
	Get(key string) (interface{}, bool)
	GetObj(key string) (interface{}, bool)
	Update(key string, val interface{})
	Delete(key string)
	GetExTime(key string) int64
	UpdateObjTimeout(key string, exTime int64)
}

// NewDefaultStore clearSleepTime is the number of seconds between each check when using default storage
func NewDefaultStore(clearSleepTime int) Store {
	d := &defaultStoreImpl{
		exMap:          sync.Map{},
		dMap:           sync.Map{},
		clearSleepTime: clearSleepTime,
	}
	d.clearGoroutine()

	return d
}

type defaultStoreImpl struct {
	dMap           sync.Map
	exMap          sync.Map
	clearSleepTime int
}

func (s *defaultStoreImpl) GetObj(key string) (interface{}, bool) {
	return s.Get(key)
}

func (s *defaultStoreImpl) Set(key string, val interface{}, exTime int64) {
	s.dMap.Store(key, val)
	if exTime != NeverExpire {
		exTime = time.Now().Unix() + exTime
	}
	s.exMap.Store(key, exTime)
}

func (s *defaultStoreImpl) Get(key string) (interface{}, bool) {
	s.clearKeyByTimeout(key)

	return s.dMap.Load(key)
}

func (s *defaultStoreImpl) Update(key string, val interface{}) {
	s.dMap.Store(key, val)
}

func (s *defaultStoreImpl) Delete(key string) {
	s.dMap.Delete(key)
	s.exMap.Delete(key)
}

func (s *defaultStoreImpl) GetExTime(key string) int64 {
	value, ok := s.exMap.Load(key)
	if ok {
		return value.(int64)
	}

	return NotValueExist
}

func (s *defaultStoreImpl) UpdateObjTimeout(key string, exTime int64) {
	s.exMap.Store(key, exTime)
}

func (s *defaultStoreImpl) clearKeyByTimeout(key string) {
	val, ok := s.exMap.Load(key)
	if ok {
		s.clearKeyByValueTimeout(key, val.(int64))
	}
}

func (s *defaultStoreImpl) clearKeyByValueTimeout(key string, exTime int64) {
	if exTime != NeverExpire && exTime < time.Now().Unix() {
		s.exMap.Delete(key)
		s.dMap.Delete(key)
	}
}

func (s *defaultStoreImpl) clearGoroutine() {
	go func() {
		for {
			s.exMap.Range(func(key, value any) bool {
				exTime := value.(int64)
				s.clearKeyByValueTimeout(key.(string), exTime)
				return true
			})
			time.Sleep(time.Second * time.Duration(s.clearSleepTime))
		}
	}()
}
