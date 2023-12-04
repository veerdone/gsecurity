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

const ContextKey = "GSecurityContextKey"

type Adaptor interface {
	GetFromHeader(tokenName string) string
	GetFromQuery(tokenName string) string
	GetFromCookie(tokenName string) string
	SetCookie(conf Config, token string)
	SetHeader(headerName, headerVal string)
	Get(key string) interface{}
	Set(key string, val interface{})
}
