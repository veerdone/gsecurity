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

package ginadaptor

import (
	"github.com/gin-gonic/gin"
	"github.com/veerdone/gsecurity"
	"github.com/veerdone/gsecurity/adaptor"
)

type ginAdaptor struct {
	*gin.Context
}

func New(c *gin.Context) adaptor.Adaptor {
	return &ginAdaptor{Context: c}
}

func (a *ginAdaptor) GetFromHeader(tokenName string) string {
	return a.GetHeader(tokenName)
}

func (a *ginAdaptor) GetFromQuery(tokenName string) string {
	return a.Query(tokenName)
}

func (a *ginAdaptor) GetFromCookie(tokenName string) string {
	cookie, err := a.Cookie(tokenName)
	if cookie != "" && err == nil {
		return cookie
	}

	return ""
}

func (a *ginAdaptor) SetHeader(headerName, headerVal string) {
	a.Context.Header(headerName, headerVal)
}

func (a *ginAdaptor) Get(key string) interface{} {
	value, exists := a.Context.Get(key)
	if exists {
		return value
	}

	return nil
}

func (a *ginAdaptor) Set(key string, val interface{}) {
	a.Context.Set(key, val)
}

func (a *ginAdaptor) SetCookie(conf gsecurity.Config, token string) {
	a.Context.SetCookie(conf.TokenName, token, int(conf.Timeout), conf.Cookie.Path,
		conf.Cookie.Domain, conf.Cookie.Secure, conf.Cookie.HttpOnly)
}
