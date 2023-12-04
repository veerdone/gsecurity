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

package echoadaptor

import (
	"github.com/labstack/echo/v4"
	"github.com/veerdone/gsecurity"
	"net/http"
	"time"
)

type echoAdaptor struct {
	echo.Context
}

func New(c echo.Context) *echoAdaptor {
	return &echoAdaptor{
		Context: c,
	}
}

func (e *echoAdaptor) GetFromHeader(tokenName string) string {
	return e.Context.Request().Header.Get(tokenName)
}

func (e *echoAdaptor) GetFromQuery(tokenName string) string {
	return e.Context.QueryParam(tokenName)
}

func (e *echoAdaptor) GetFromCookie(tokenName string) string {
	cookie, err := e.Context.Cookie(tokenName)
	if err != nil {
		return ""
	}

	return cookie.Value
}

func (e *echoAdaptor) SetHeader(headerName, headerValue string) {
	e.Context.Response().Header().Set(headerName, headerValue)
}

func (e *echoAdaptor) Get(key string) interface{} {
	return e.Context.Get(key)
}

func (e *echoAdaptor) Set(key string, v interface{}) {
	e.Context.Set(key, v)
}

func (e *echoAdaptor) SetCookie(conf gsecurity.Config, token string) {
	sameSite := http.SameSiteDefaultMode
	switch conf.Cookie.SameSite {
	case "None":
		sameSite = http.SameSiteNoneMode
	case "Lax":
		sameSite = http.SameSiteLaxMode
	case "strict":
		sameSite = http.SameSiteStrictMode
	}

	c := http.Cookie{
		Name:     conf.TokenName,
		Value:    token,
		Path:     conf.Cookie.Path,
		Domain:   conf.Cookie.Domain,
		Expires:  time.Now().Add(time.Second * time.Duration(conf.Timeout)),
		Secure:   conf.Cookie.Secure,
		HttpOnly: conf.Cookie.HttpOnly,
		SameSite: sameSite,
	}
	e.Context.SetCookie(&c)
}
