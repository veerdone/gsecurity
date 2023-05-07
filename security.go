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

type Security struct {
	*Logic
}

func NewSecurity(l *Logic) *Security {
	return &Security{Logic: l}
}

func (s *Security) Login(id int64) string {
	return s.Logic.Login(id)
}

func (s *Security) LoginAndSet(id int64, a Adaptor) string {
	token := s.Login(id)
	a.SetCookie(s.Config, token)

	return token
}

func (s *Security) LoginWithDevice(id int64, device string) string {
	return s.Logic.LoginWithDevice(id, device)
}

func (s *Security) IsLogin(a Adaptor) bool {
	token := a.GetToken(s.TokenName)

	return s.IsLoginByToken(token)
}

func (s *Security) GetToken(a Adaptor) string {
	return a.GetToken(s.TokenName)
}

func (s *Security) Session(a Adaptor) *Session {
	token := a.GetToken(s.TokenName)

	return s.GetSessionByToken(token)
}

func (s *Security) Logout(a Adaptor) {
	token := a.GetToken(s.TokenName)

	s.Logic.LogoutByToken(token)
}

func (s *Security) LogoutById(id int64) {
	s.Logic.Logout(id)
}

func (s *Security) LogoutByIdAndDevice(id int64, device string) {
	s.Logic.LogoutByIdAndDevice(id, device)
}
