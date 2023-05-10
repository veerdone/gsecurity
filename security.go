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

import "github.com/veerdone/gsecurity/adaptor"

type Security struct {
	*Logic
}

func NewSecurity(l *Logic) *Security {
	return &Security{Logic: l}
}

// Login use id login, return token
func (s *Security) Login(id int64) string {
	return s.Logic.Login(id)
}

// LoginAndSet use id to login, and set token to cookie, return token
func (s *Security) LoginAndSet(id int64, a adaptor.Adaptor) string {
	token := s.Login(id)
	if s.WriteToCookie {
		a.SetCookie(s.Config, token)
	}
	if s.WriteToHeader {
		a.SetHeader(s.TokenName, token)
	}

	return token
}

// LoginWithDevice use id and device to login, return token
func (s *Security) LoginWithDevice(id int64, device string) string {
	return s.Logic.LoginWithDevice(id, device)
}

func (s *Security) getToken(a adaptor.Adaptor) string {
	if s.ReadFromHeader {
		if token := a.GetFromHeader(s.TokenName); token != "" {
			return token
		}
	}
	if s.ReadFromCookie {
		if token := a.GetFromCookie(s.TokenName); token != "" {
			return token
		}
	}
	if s.ReadFromQuery {
		if token := a.GetFromQuery(s.TokenName); token != "" {
			return token
		}
	}

	return ""
}

// IsLogin get token from adaptor.Adaptor and check token login or not
func (s *Security) IsLogin(a adaptor.Adaptor) bool {
	token := s.getToken(a)

	return s.IsLoginByToken(token)
}

// CheckLogin get token from adaptor.Adaptor and check token login or not,
// if not login, return ErrNotLogin
func (s *Security) CheckLogin(a adaptor.Adaptor) error {
	token := s.getToken(a)

	return s.CheckLoginByToken(token)
}

// GetToken get token from adaptor.Adaptor
func (s *Security) GetToken(a adaptor.Adaptor) string {
	return s.getToken(a)
}

// Session get token from adaptor.Adaptor then get Session by token
func (s *Security) Session(a adaptor.Adaptor) *Session {
	token := s.getToken(a)

	return s.GetSessionByToken(token)
}

// Logout get token from adaptor.Adaptor then use token to logout
func (s *Security) Logout(a adaptor.Adaptor) {
	token := s.getToken(a)

	s.Logic.LogoutByToken(token)
}

// LogoutById logout of the id
func (s *Security) LogoutById(id int64) {
	s.Logic.Logout(id)
}

// LogoutByIdAndDevice logout of the id and device
func (s *Security) LogoutByIdAndDevice(id int64, device string) {
	s.Logic.LogoutByIdAndDevice(id, device)
}

func (s *Security) Disable(id, exTime int64) {
	s.Logic.DisableWithLevelAndService(id, 1, exTime, defaultDisableService)
}

func (s *Security) DisableWithLevel(id, level, exTime int64) {
	s.Logic.DisableWithLevelAndService(id, level, exTime, defaultDisableService)
}

func (s *Security) DisableWithService(id, exTime int64, service string) {
	s.Logic.DisableWithLevelAndService(id, 1, exTime, service)
}

// DisableWithLevelAndService disable with the id, level, expire time of seconds and service
func (s *Security) DisableWithLevelAndService(id, level, exTime int64, service string) {
	s.Logic.DisableWithLevelAndService(id, level, exTime, service)
}

func (s *Security) IsDisable(id int64) bool {
	return s.Logic.IsDisableWithLevelAndService(id, 1, defaultDisableService)
}

func (s *Security) IsDisableWithLevel(id, level int64) bool {
	return s.Logic.IsDisableWithLevelAndService(id, level, defaultDisableService)
}

func (s *Security) IsDisableWithService(id int64, service string) bool {
	return s.Logic.IsDisableWithLevelAndService(id, 1, service)
}

// IsDisableWithLevelAndService check is disable with id, level and service
func (s *Security) IsDisableWithLevelAndService(id, level int64, service string) bool {
	return s.Logic.IsDisableWithLevelAndService(id, level, service)
}

func (s *Security) CheckDisable(id int64) error {
	return s.Logic.CheckDisableWithLevelAndService(id, 1, defaultDisableService)
}

func (s *Security) CheckDisableWithLevel(id, level int64) error {
	return s.Logic.CheckDisableWithLevelAndService(id, level, defaultDisableService)
}

func (s *Security) CheckDisableWithService(id int64, service string) error {
	return s.Logic.CheckDisableWithLevelAndService(id, 1, service)
}

// CheckDisableWithLevelAndService check disable with id, level and service,
// if it's disabled, return ErrDisable
func (s *Security) CheckDisableWithLevelAndService(id, level int64, service string) error {
	return s.Logic.CheckDisableWithLevelAndService(id, level, service)
}

func (s *Security) RmDisable(id int64) {
	s.Logic.RmDisableWithServices(id, defaultDisableService)
}

// RmDisableWithServices remove disable with id and services
func (s *Security) RmDisableWithServices(id int64, services ...string) {
	s.Logic.RmDisableWithServices(id, services...)
}

func (s *Security) DisableExTime(id int64) int64 {
	return s.Logic.DisableExTime(id, defaultDisableService)
}

// DisableExTimeWithService get disabled expire time, if never expire return NeverExpire,
// if not disable return NotValueExist
func (s *Security) DisableExTimeWithService(id int64, service string) int64 {
	return s.Logic.DisableExTime(id, service)
}
