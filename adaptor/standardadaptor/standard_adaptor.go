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

package standardadaptor

import (
	"github.com/veerdone/gsecurity"
	"github.com/veerdone/gsecurity/adaptor"
	"net/http"
	"sync"
)

type standardAdaptor struct {
	*http.Request
	m sync.Map
}

func New(req *http.Request) adaptor.Adaptor {
	return &standardAdaptor{Request: req, m: sync.Map{}}
}

func (a *standardAdaptor) GetFromHeader(tokenName string) string {
	return a.Header.Get(tokenName)
}

func (a *standardAdaptor) GetFromQuery(tokenName string) string {
	return a.URL.Query().Get(tokenName)
}

func (a *standardAdaptor) GetFromCookie(tokenName string) string {
	cookie, err := a.Cookie(tokenName)
	if cookie != nil && err == nil {
		return cookie.Value
	}

	return ""
}

func (a *standardAdaptor) SetHeader(headerName, headerVal string) {
	a.Response.Header.Add(headerVal, headerVal)
}

func (a *standardAdaptor) Get(key string) interface{} {
	value, _ := a.m.Load(key)

	return value
}

func (a *standardAdaptor) Set(key string, val interface{}) {
	a.m.Store(key, val)
}

func (a *standardAdaptor) SetCookie(conf gsecurity.Config, token string) {

}
