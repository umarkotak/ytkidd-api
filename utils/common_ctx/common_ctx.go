package common_ctx

import (
	"context"
	"net/http"
)

type (
	CommonCtxKeyType string

	UserAuth struct {
		GUID     string `json:"guid"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	CommonCtx struct {
		DeviceID        string // X-Device-Id. device fingerprint must unique as possible per device
		DeviceOs        string // X-Device-Os. device os, Eg: android/ios/windows/ubuntu/other
		AppVersion      string // X-App-Version. Eg: 1.20.30
		ActivitySession string // X-Activity-Session. generated every opening app

		UserAuth UserAuth
	}
)

var (
	CommonCtxKey = CommonCtxKeyType("common_ctx")
)

func Get(r *http.Request) CommonCtx {
	commonCtx := CommonCtx{
		UserAuth: UserAuth{},
	}

	v := r.Context().Value(CommonCtxKey)

	if v == nil {
		return commonCtx
	}

	commonCtx, ok := v.(CommonCtx)

	if !ok {
		return commonCtx
	}

	return commonCtx
}

func GetFromCtx(ctx context.Context) CommonCtx {
	commonCtx := CommonCtx{
		UserAuth: UserAuth{},
	}

	v := ctx.Value(CommonCtxKey)

	if v == nil {
		return commonCtx
	}

	commonCtx, ok := v.(CommonCtx)

	if !ok {
		return commonCtx
	}

	return commonCtx
}

func FromRequestHeader(r *http.Request) CommonCtx {
	deviceOs := r.Header.Get("X-Device-Os")
	if deviceOs == "" {
		deviceOs = "unknown"
	}

	return CommonCtx{
		DeviceID:        r.Header.Get("X-Device-Id"),
		DeviceOs:        deviceOs,
		ActivitySession: r.Header.Get("X-Activity-Session"),
		AppVersion:      r.Header.Get("X-App-Version"),
	}
}
