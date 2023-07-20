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

import "go.uber.org/zap"

type logger struct {
	*zap.Logger
	enable bool
}

func newLogger() *zap.Logger {
	l, _ := zap.NewProduction(zap.AddCallerSkip(1), zap.WithCaller(true))

	return l
}

var log = &logger{
	Logger: newLogger(),
	enable: true,
}

func SetLogger(l *zap.Logger) {
	log.Logger = l
}

func SetLoggerEnable(enable bool) {
	log.enable = enable
}

func (l *logger) Info(m string, fields ...zap.Field) {
	if l.enable {
		l.Logger.Info(m, fields...)
	}
}

func (l *logger) Warn(m string, fields ...zap.Field) {
	if l.enable {
		l.Logger.Warn(m, fields...)
	}
}

func (l *logger) Debug(m string, fields ...zap.Field) {
	if l.enable {
		l.Logger.Debug(m, fields...)
	}
}

func (l *logger) Error(m string, fields ...zap.Field) {
	if l.enable {
		l.Logger.Error(m, fields...)
	}
}
