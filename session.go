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
	"github.com/bytedance/sonic"
)

type Session struct {
	Id            string                 `json:"id,omitempty"`
	CreateTime    int64                  `json:"createTime,omitempty"`
	Data          map[string]interface{} `json:"data,omitempty"`
	TokenSignList []TokenSign            `json:"tokenSignList,omitempty"`
	store         Store
}

func (s *Session) Get(key string) (interface{}, bool) {
	if s.Data == nil {
		return nil, false
	}
	v, ok := s.Data[key]

	return v, ok
}

func (s *Session) Set(key string, val interface{}) {
	if s.Data == nil {
		s.Data = make(map[string]interface{})
	}
	s.Data[key] = val
	s.store.Update(s.Id, s)
}

func (s *Session) UnmarshalBinary(data []byte) error {
	return sonic.Unmarshal(data, s)
}

func (s *Session) MarshalBinary() (data []byte, err error) {
	return sonic.Marshal(s)
}

func (s *Session) AddTokenSign(ts TokenSign) {
	notExist := true
	for i := range s.TokenSignList {
		if ts.Val == s.TokenSignList[i].Val {
			notExist = false
			break
		}
	}
	if notExist {
		s.TokenSignList = append(s.TokenSignList, ts)
		s.store.Update(s.Id, s)
	}
}

func (s *Session) GetTokenSignByDevice(device string) (TokenSign, bool) {
	for i := range s.TokenSignList {
		if s.TokenSignList[i].Device == device {
			return s.TokenSignList[i], true
		}
	}

	return TokenSign{}, false
}

func (s *Session) DelTokenSignByToken(token string) {
	tokenSignList := s.TokenSignList
	for i := range tokenSignList {
		if tokenSignList[i].Val == token {
			s.TokenSignList = removeTokenSignByIndex(i, tokenSignList)
			s.store.Update(s.Id, s)
		}
	}
}

func removeTokenSignByIndex(index int, tokenSignList []TokenSign) []TokenSign {
	result := make([]TokenSign, 0, len(tokenSignList)-1)
	result = append(result, tokenSignList[:index]...)
	if index != len(tokenSignList)-1 {
		result = append(result, tokenSignList[index+1:]...)
	}

	return result
}

type TokenSign struct {
	Val    string `json:"val,omitempty"`
	Device string `json:"device,omitempty"`
}

func (t *TokenSign) MarshalBinary() (data []byte, err error) {
	return sonic.Marshal(t)
}

func (t *TokenSign) UnmarshalBinary(data []byte) error {
	return sonic.Unmarshal(data, t)
}
