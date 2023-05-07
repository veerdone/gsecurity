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

import (
	"github.com/google/uuid"
	"math/rand"
	"unsafe"
)

type GenerateToken func() string

var (
	UUID    GenerateToken = uuid.NewString
	Rand32  GenerateToken = rand32
	Rand64  GenerateToken = rand64
	Rand128 GenerateToken = rand128
)

const (
	baseNumber     = "1234567890"
	baseChar       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	baseCharNumber = baseNumber + baseChar
)

func rand32() string {
	return randString(32)
}

func rand64() string {
	return randString(64)
}

func rand128() string {
	return randString(128)
}

func randString(l int) string {
	b := make([]byte, l)
	for i := 0; i < l; i++ {
		b[i] = baseCharNumber[rand.Intn(len(baseCharNumber))]
	}

	return sliceByteToString(b)
}

func sliceByteToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
