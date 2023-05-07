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

import "errors"

const (
	// NeverExpire token never expire
	NeverExpire = -1
	// NotValueExist Store getExTime if value not exist return this value
	NotValueExist = -2
	// BeReplace being knocked off the line
	BeReplace = -4
)

var (
	ErrBeReplace = errors.New("has been replaced")
	ErrNotLogin  = errors.New("not login")
	abnormalMap  = map[int64]error{
		BeReplace: ErrBeReplace,
	}
)

// isValidId check id is valid
func isValidId(id int64) bool {
	return checkValidId(id) == nil
}

// checkValidId check id is valid, if it is not valid, return error
func checkValidId(id int64) error {
	return abnormalMap[id]
}
