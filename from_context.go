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
	"context"
	"errors"
)

var ErrNoAdaptorInContext = errors.New("no adaptor in context")

// LoginAndSetFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func LoginAndSetFromCtx(id int64, ctx context.Context) string {
	if a, ok := getFromCtx(ctx); ok {
		return defaultSecurity.LoginAndSet(id, a)
	}

	return ""
}

// IsLoginFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func IsLoginFromCtx(ctx context.Context) bool {
	if a, ok := getFromCtx(ctx); ok {
		return defaultSecurity.IsLogin(a)
	}

	return false
}

// CheckLoginFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func CheckLoginFromCtx(ctx context.Context) error {
	if a, ok := getFromCtx(ctx); ok {
		return defaultSecurity.CheckLogin(a)
	}

	return ErrNoAdaptorInContext
}

// GetLoginIdFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func GetLoginIdFromCtx(ctx context.Context) int64 {
	if a, ok := getFromCtx(ctx); ok {
		return defaultSecurity.GetLoginId(a)
	}

	return 0
}

// GetLoginTokenFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func GetLoginTokenFromCtx(ctx context.Context) string {
	if a, ok := getFromCtx(ctx); ok {
		return defaultSecurity.GetToken(a)
	}

	return ""
}

// SessionsFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func SessionsFromCtx(ctx context.Context) *Session {
	if a, ok := getFromCtx(ctx); ok {
		return defaultSecurity.Session(a)
	}

	return nil
}

// LogoutFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func LogoutFromCtx(ctx context.Context) {
	if a, ok := getFromCtx(ctx); ok {
		defaultSecurity.Logout(a)
	}
}

// HasRoleFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func HasRoleFromCtx(ctx context.Context, role string) bool {
	if a, ok := getFromCtx(ctx); ok {
		return defaultSecurity.HasRole(a, role)
	}

	return false
}

// HasRoleOrFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func HasRoleOrFromCtx(ctx context.Context, roles ...string) bool {
	if a, ok := getFromCtx(ctx); ok {
		return defaultSecurity.HasRoleOr(a, roles...)
	}

	return false
}

// HasRoleAndFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func HasRoleAndFromCtx(ctx context.Context, roles ...string) bool {
	if a, ok := getFromCtx(ctx); ok {
		return defaultSecurity.HasRoleAnd(a, roles...)
	}

	return false
}

// HasPermissionFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func HasPermissionFromCtx(ctx context.Context, p string) bool {
	if a, ok := getFromCtx(ctx); ok {
		return defaultSecurity.HasPermission(a, p)
	}

	return false
}

// HasPermissionOrFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func HasPermissionOrFromCtx(ctx context.Context, ps ...string) bool {
	if a, ok := getFromCtx(ctx); ok {
		return defaultSecurity.HasPermissionOr(a, ps...)
	}

	return false
}

// HasPermissionAndFromCtx The method at the end of FromCtx can be used only when Adaptor is the value and ContextKey is the Key in the Context
func HasPermissionAndFromCtx(ctx context.Context, ps ...string) bool {
	if a, ok := getFromCtx(ctx); ok {
		return defaultSecurity.HasPermissionAnd(a, ps...)
	}

	return false
}

func getFromCtx(ctx context.Context) (Adaptor, bool) {
	v := ctx.Value(ContextKey)
	if v == nil {
		return nil, false
	}
	a, ok := v.(Adaptor)

	return a, ok
}
