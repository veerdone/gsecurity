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

var defaultSecurity = &Security{
	Logic: NewLogic(DefaultConfig, NewDefaultStore(30)),
}

func SetDefaultSecurity(l *Logic) {
	defaultSecurity.Logic = l
}

func Login(id int64) string {
	return defaultSecurity.Login(id)
}

func LoginAndSet(id int64, a Adaptor) string {
	return defaultSecurity.LoginAndSet(id, a)
}

func IsLogin(a Adaptor) bool {
	return defaultSecurity.IsLogin(a)
}

func Sessions(a Adaptor) *Session {
	return defaultSecurity.Session(a)
}

func Logout(a Adaptor) {
	defaultSecurity.Logout(a)
}
