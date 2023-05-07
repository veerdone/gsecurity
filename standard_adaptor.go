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

import "net/http"

type standardAdaptor struct {
	*http.Request
}

func Standard(req *http.Request) Adaptor {
	return &standardAdaptor{Request: req}
}

func (a *standardAdaptor) GetToken(tokenName string) string {
	cookie, err := a.Cookie(tokenName)
	if cookie != nil && err != nil {
		return cookie.Value
	}

	if header := a.Header.Get(tokenName); header != "" {
		return header
	}

	return a.URL.Query().Get(tokenName)
}

func (a *standardAdaptor) SetCookie(conf Config, token string) {

}
