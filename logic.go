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
	"fmt"
	"go.uber.org/zap"
	"time"
)

const (
	defaultLoginType = "login"
	defaultDisableService
	DefaultDevice = "default-device"
	sessionKey    = "%s:%s:session:%d"
	tokenKey      = "%s:%s:token:%s"
	disableKey    = "%s:%s:disable:%s:%d"
)

type Logic struct {
	Store
	Config
	LoginType     string
	authorization Authorization
}

func NewLogic(conf Config, store Store) *Logic {
	return NewLogicWithLoginType(defaultLoginType, conf, store)
}

func NewLogicWithLoginType(loginType string, config Config, store Store) *Logic {
	checkConfig(config)

	return &Logic{
		LoginType: loginType,
		Config:    config,
		Store:     store,
	}
}

func (l *Logic) GetConfig() Config {
	return l.Config
}

func (l *Logic) GetStore() Store {
	return l.Store
}

// GetIdByToken if not login or token is invalid, return 0
func (l *Logic) GetIdByToken(token string) int64 {
	cacheId, ok := l.Store.Get(l.buildTokenKey(token))
	if ok {
		id := cacheId.(int64)
		if isValidId(id) {
			return id
		}
	}

	return 0
}

// GetSessionByToken get session by token, if token not exist, return nil
func (l *Logic) GetSessionByToken(token string) *Session {
	tokenValue := l.buildTokenKey(token)
	loginId, ok := l.Get(tokenValue)
	if !ok {
		return nil
	}

	id := loginId.(int64)
	if isValidId(id) {
		return l.GetSessionById(id)
	}

	return nil
}

// GetSessionById get session by id, if id not exist, return nil
func (l *Logic) GetSessionById(id int64) *Session {
	sessionValueKey := l.buildSessionKey(id)
	s, ok := l.GetSession(sessionValueKey)
	if ok {
		session := s.(*Session)
		if session.store == nil {
			session.store = l.Store
		}
		return session
	}

	return nil
}

// GetSessionByIdOrCreate get session by id, if not exist, create and return
func (l *Logic) GetSessionByIdOrCreate(id int64) *Session {
	s := l.GetSessionById(id)
	if s == nil {
		s = l.createSession(id)
	}

	return s
}

// IsLoginByToken validate token is login
func (l *Logic) IsLoginByToken(token string) bool {
	v, ok := l.Get(l.buildTokenKey(token))

	return ok && isValidId(v.(int64))
}

// CheckLoginByToken check token is login, if it's login, return nil, else return error
func (l *Logic) CheckLoginByToken(token string) error {
	id, ok := l.Get(l.buildTokenKey(token))
	if ok {
		return checkValidId(id.(int64))
	}

	return ErrNotLogin
}

// createSession create session by id
func (l *Logic) createSession(id int64) *Session {
	s := &Session{
		CreateTime: time.Now().Unix(),
		Id:         l.buildSessionKey(id),
		store:      l.Store,
	}
	l.Set(s.Id, s, l.Timeout)

	return s
}

// LogoutByToken logout with token
func (l *Logic) LogoutByToken(token string) {
	tokenKey := l.buildTokenKey(token)
	loginId, ok := l.Store.Get(tokenKey)
	if ok {
		l.Store.Delete(tokenKey)
		session := l.GetSessionById(loginId.(int64))
		if session != nil {
			session.DelTokenSignByToken(token)
			if len(session.TokenSignList) == 0 {
				l.Store.Delete(session.Id)
			}
			log.Info("user logout", zap.String("loginType", l.LoginType), zap.Any("loginId", loginId))
		}
	}
}

// Logout with id
func (l *Logic) Logout(id int64) {
	l.LogoutByIdAndDevice(id, DefaultDevice)
}

// LogoutByIdAndDevice logout with id and device
func (l *Logic) LogoutByIdAndDevice(id int64, device string) {
	session := l.GetSessionById(id)
	if session != nil {
		tokenSign, ok := session.GetTokenSignByDevice(device)
		if ok {
			l.removeTokenSign(session, tokenSign.Val)
			l.Store.Delete(l.buildTokenKey(tokenSign.Val))
			log.Info("user logout", zap.String("loginType", l.LoginType), zap.Int64("loginId", id))
		}
	}
}

// GetTokenByIdAndDevice get token by id and device, if not exist, return ""
func (l *Logic) GetTokenByIdAndDevice(id int64, device string) string {
	session := l.GetSessionById(id)
	if session != nil {
		if tokenSign, ok := session.GetTokenSignByDevice(device); ok {
			return l.buildTokenKey(tokenSign.Val)
		}
	}

	return ""
}

// Login use id login
func (l *Logic) Login(id int64) string {
	return l.LoginWithDevice(id, DefaultDevice)
}

// LoginWithDevice login with id and device
func (l *Logic) LoginWithDevice(id int64, device string) string {
	token := l.assignToken(id, device)
	l.createLoginSession(id, device, token)

	log.Info("user login", zap.String("loginType", l.LoginType), zap.Int64("loginId", id),
		zap.String("device", device))

	return token
}

func (l *Logic) Kick(id int64) {
	l.kickWithDevice(id, DefaultDevice)
}

func (l *Logic) kickWithDevice(id int64, device string) {
	session := l.GetSessionById(id)
	if session != nil {
		tokenSign, ok := session.GetTokenSignByDevice(device)
		if ok {
			l.removeTokenSign(session, tokenSign.Val)
			l.Store.Update(l.buildTokenKey(tokenSign.Val), BeKick)
			log.Info("kick user", zap.String("loginType", l.LoginType), zap.Int64("loginId", id),
				zap.String("device", device))
		}
	}
}

func (l *Logic) KickWithToken(token string) {
	session := l.GetSessionByToken(token)
	if session != nil {
		tokenSign, ok := session.GetTokenSignByToken(token)
		if ok {
			l.removeTokenSign(session, tokenSign.Val)
			l.Store.Update(l.buildTokenKey(tokenSign.Val), BeKick)
			log.Info("kick user with token", zap.String("loginType", l.LoginType),
				zap.String("device", tokenSign.Device), zap.String("token", token))
		}
	}
}

// assignToken assign token with id and device
func (l *Logic) assignToken(id int64, device string) string {
	var session *Session
	session = l.GetSessionById(id)
	if !l.IsConcurrent {
		if session != nil {
			if tokenSign, ok := session.GetTokenSignByDevice(device); ok {
				// remove tokenSign
				session.DelTokenSignByToken(tokenSign.Val)
				// mark token as BeReplace
				l.Store.Update(l.buildTokenKey(tokenSign.Val), BeReplace)
			}
		}
	}

	if l.IsConcurrent && l.IsShare {
		if session != nil {
			if tokenSign, ok := session.GetTokenSignByDevice(device); ok {
				return tokenSign.Val
			}
		}
	}

	return l.TokenStyle()
}

func (l *Logic) buildSessionKey(id int64) string {
	return fmt.Sprintf(sessionKey, l.Config.TokenName, l.LoginType, id)
}

func (l *Logic) buildTokenKey(token string) string {
	return fmt.Sprintf(tokenKey, l.TokenName, l.LoginType, token)
}

func (l *Logic) buildDisableKey(id int64, service string) string {
	return fmt.Sprintf(disableKey, l.TokenName, l.LoginType, service, id)
}

func (l *Logic) createLoginSession(id int64, device string, token string) {
	session := l.GetSessionByIdOrCreate(id)

	ts := TokenSign{
		Val:    token,
		Device: device,
	}
	session.AddTokenSign(ts)

	l.Set(l.buildTokenKey(token), id, l.Timeout)
}

// GetTokenTimeout get expire time by token, return value is the number of seconds
func (l *Logic) GetTokenTimeout(token string) int64 {
	return l.GetExTime(l.buildTokenKey(token))
}

func (l *Logic) DisableWithLevelAndService(id, level, exTime int64, service string) {
	l.Store.Set(l.buildDisableKey(id, service), level, exTime)
	log.Info("disable user", zap.String("loginType", l.LoginType), zap.String("service", service),
		zap.Int64("loginId", id), zap.Int64("disable_level", level), zap.Int64("disable_time", exTime))
}

func (l *Logic) IsDisableWithLevelAndService(id, level int64, service string) bool {
	err := l.CheckDisableWithLevelAndService(id, level, service)

	return err != nil
}

func (l *Logic) CheckDisableWithLevelAndService(id, level int64, service string) error {
	key := l.buildDisableKey(id, service)
	cLevel, ok := l.Store.Get(key)
	if ok && cLevel.(int64) >= level {
		return NewErrDisable(level, service)
	}

	return nil
}

func (l *Logic) RmDisableWithServices(id int64, services ...string) {
	for _, service := range services {
		l.Store.Delete(l.buildDisableKey(id, service))
		log.Info("remove disable", zap.String("loginType", l.LoginType), zap.Int64("loginId", id),
			zap.Strings("services", services))
	}
}

func (l *Logic) DisableExTime(id int64, service string) int64 {
	return l.Store.GetExTime(l.buildDisableKey(id, service))
}

func (l *Logic) removeTokenSign(session *Session, token string) {
	session.DelTokenSignByToken(token)
	if len(session.TokenSignList) == 0 {
		l.Store.Delete(session.Id)
	}
}

func (l *Logic) SetAuthorization(a Authorization) {
	l.authorization = a
}

func (l *Logic) GetPermissionList(id int64) []string {
	if l.authorization != nil {
		return l.authorization.GetPermissionList(id)
	}

	return []string{}
}

func (l *Logic) GetRoleList(id int64) []string {
	if l.authorization != nil {
		return l.authorization.GetRoleList(id)
	}

	return []string{}
}
