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

var (
	DefaultConfig = Config{
		TokenName:      "GSecurity",
		Timeout:        2592000,
		IsConcurrent:   true,
		IsShare:        true,
		TokenStyle:     UUID,
		ReadFromQuery:  false,
		ReadFromCookie: true,
		ReadFromHeader: false,
		WriteToHeader:  false,
		WriteToCookie:  true,
		Cookie: Cookie{
			Path:     "/",
			Secure:   false,
			HttpOnly: false,
			SameSite: "Lax",
		},
	}
)

type Config struct {
	// token name (also Cookie name and data persistence prefix)
	TokenName string
	// token validity period
	Timeout int64
	// whether to allow concurrent logins with the same account (if true, allow concurrent logins, if false,
	// new logins will crowd out old logins)
	IsConcurrent bool
	// whether to share a token when multiple people log in to the same account (if true, all logins share a token;
	// if false, create a token for each login)
	IsShare bool
	// generated token's style, can be customized
	TokenStyle GenerateToken
	// read token from query, default true
	ReadFromQuery bool
	// read token from header, default true
	ReadFromHeader bool
	// read token from cookie, default true
	ReadFromCookie bool
	// set token to header, default false
	WriteToHeader bool
	// set token to cookie, default true
	WriteToCookie bool
	Cookie        Cookie
}

type Cookie struct {
	Domain   string
	Path     string
	Secure   bool
	HttpOnly bool
	SameSite string
}

func checkConfig(conf Config) {
	if conf.TokenName == "" {
		conf.TokenName = DefaultConfig.TokenName
	}
	if conf.TokenStyle == nil {
		conf.TokenStyle = DefaultConfig.TokenStyle
	}
	if conf.Timeout == 0 {
		conf.Timeout = DefaultConfig.Timeout
	}
}
