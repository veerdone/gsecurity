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

package hertzadaptor

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/veerdone/gsecurity"
	"time"
)

type hertzAdaptor struct {
	*app.RequestContext
}

func New(c *app.RequestContext) gsecurity.Adaptor {
	return hertzAdaptor{
		RequestContext: c,
	}
}

func (h hertzAdaptor) GetFromHeader(tokenName string) string {
	return string(h.GetHeader(tokenName))
}

func (h hertzAdaptor) GetFromQuery(tokenName string) string {
	return h.Query(tokenName)
}

func (h hertzAdaptor) GetFromCookie(tokenName string) string {
	return string(h.Cookie(tokenName))
}

func (h hertzAdaptor) SetCookie(conf gsecurity.Config, token string) {
	c := protocol.Cookie{}
	c.SetKey(conf.TokenName)
	c.SetValue(token)
	c.SetDomain(conf.Cookie.Domain)
	c.SetHTTPOnly(conf.Cookie.HttpOnly)
	c.SetPath(conf.Cookie.Path)
	expire := time.Now().Add(time.Duration(conf.Timeout) * time.Second)
	c.SetExpire(expire)
	c.SetSecure(conf.Cookie.Secure)

	h.Response.Header.SetCookie(&c)
}

func (h hertzAdaptor) SetHeader(headerName, headerVal string) {
	h.RequestContext.Header(headerName, headerVal)
}

func (h hertzAdaptor) Get(key string) interface{} {
	value, _ := h.RequestContext.Get(key)

	return value
}
