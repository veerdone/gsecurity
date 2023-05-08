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

var defaultSecurity = &Security{
	Logic: NewLogic(DefaultConfig, NewDefaultStore(30)),
}

func SetDefaultSecurity(l *Logic) {
	defaultSecurity.Logic = l
}

func Login(id int64) string {
	return defaultSecurity.Login(id)
}

func LoginAndSet(id int64, a adaptor.Adaptor) string {
	return defaultSecurity.LoginAndSet(id, a)
}

func IsLogin(a adaptor.Adaptor) bool {
	return defaultSecurity.IsLogin(a)
}

func CheckLogin(a adaptor.Adaptor) error {
	return defaultSecurity.CheckLogin(a)
}

func Sessions(a adaptor.Adaptor) *Session {
	return defaultSecurity.Session(a)
}

func Logout(a adaptor.Adaptor) {
	defaultSecurity.Logout(a)
}

func LogoutById(id int64) {
	defaultSecurity.LogoutById(id)
}

func LogoutByIdAndDevice(id int64, device string) {
	defaultSecurity.LogoutByIdAndDevice(id, device)
}

func Disable(id, exTime int64) {
	defaultSecurity.DisableWithLevelAndService(id, 1, exTime, defaultDisableService)
}

func DisableWithLevel(id, level, exTime int64) {
	defaultSecurity.DisableWithLevelAndService(id, level, exTime, defaultDisableService)
}

func DisableWithService(id, exTime int64, service string) {
	defaultSecurity.DisableWithLevelAndService(id, 1, exTime, service)
}

// DisableWithLevelAndService disable with the id, level, expire time of seconds and service
func DisableWithLevelAndService(id, level, exTime int64, service string) {
	defaultSecurity.DisableWithLevelAndService(id, level, exTime, service)
}

func IsDisable(id int64) bool {
	return defaultSecurity.IsDisableWithLevelAndService(id, 1, defaultDisableService)
}

func IsDisableWithLevel(id, level int64) bool {
	return defaultSecurity.IsDisableWithLevelAndService(id, level, defaultDisableService)
}

func IsDisableWithService(id int64, service string) bool {
	return defaultSecurity.IsDisableWithLevelAndService(id, 1, service)
}

// IsDisableWithLevelAndService check is disable with id, level and service
func IsDisableWithLevelAndService(id, level int64, service string) bool {
	return defaultSecurity.IsDisableWithLevelAndService(id, level, service)
}

func CheckDisable(id int64) error {
	return defaultSecurity.CheckDisableWithLevelAndService(id, 1, defaultDisableService)
}

func CheckDisableWithLevel(id, level int64) error {
	return defaultSecurity.CheckDisableWithLevelAndService(id, level, defaultDisableService)
}

func CheckDisableWithService(id int64, service string) error {
	return defaultSecurity.CheckDisableWithLevelAndService(id, 1, service)
}

// CheckDisableWithLevelAndService check disable with id, level and service,
// if it's disabled, return ErrDisable
func CheckDisableWithLevelAndService(id, level int64, service string) error {
	return defaultSecurity.CheckDisableWithLevelAndService(id, level, service)
}

func RmDisable(id int64) {
	defaultSecurity.RmDisableWithServices(id, defaultDisableService)
}

// RmDisableWithServices remove disable with id and services
func RmDisableWithServices(id int64, services ...string) {
	defaultSecurity.RmDisableWithServices(id, services...)
}

func DisableExTime(id int64) int64 {
	return defaultSecurity.DisableExTime(id)
}

// DisableExTimeWithService get disabled expire time, if never expire return NeverExpire,
// if not disable return NotValueExist
func DisableExTimeWithService(id int64, services string) int64 {
	return defaultSecurity.DisableExTimeWithService(id, services)
}
