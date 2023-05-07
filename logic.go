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
	"time"
)

const (
	defaultLoginType = "login"
	DefaultDevice    = "default-device"
	sessionKey       = "%s:%s:session:%d"
	tokenKey         = "%s:%s:token:%s"
)

type Logic struct {
	Store
	Config
	LoginType string
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
	s, ok := l.GetObj(sessionValueKey)
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
			session.DelTokenSignByToken(tokenSign.Val)
			if len(session.TokenSignList) == 0 {
				l.Store.Delete(session.Id)
			}
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

	return token
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