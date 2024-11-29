package tencent_lib

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
)

// Tencent advance permission control: https://www.tencentcloud.com/document/product/647/35157
// Tencent go lib ref: https://github.com/tencentyun/tls-sig-api-v2-golang/blob/master/tencentyun/TLSSigAPI.go#L156
const (
	HostPrivilege     = uint32(255)
	AudiencePrivilege = uint32(42)
)

// expirySecond -> expiry after second, eg: 86400 mean expire after 1 day
func GenGenericSig(ctx context.Context, userID string, expirySecond int) (string, error) {
	sig, err := tencentyun.GenUserSig(
		tencentLib.TencentSdkAppID, tencentLib.TencentSdkAppSecret, userID, expirySecond,
	)
	return sig, err
}

func GenHostSig(ctx context.Context, roomID, userID string, expirySecond int) (string, error) {
	if !tencentLib.EnableAdvancePermissionControl {
		return GenGenericSig(ctx, userID, expirySecond)
	}

	sig, err := GenAdvancedSig(ctx, roomID, userID, expirySecond, HostPrivilege)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
	}
	return sig, err
}

func GenAudienceSig(ctx context.Context, roomID, userID string, expirySecond int) (string, error) {
	if !tencentLib.EnableAdvancePermissionControl {
		return GenGenericSig(ctx, userID, expirySecond)
	}

	sig, err := GenAdvancedSig(ctx, roomID, userID, expirySecond, AudiencePrivilege)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
	}
	return sig, err
}

func GenAdvancedSig(ctx context.Context, roomID, userID string, expirySecond int, privilegeMap uint32) (string, error) {
	sig, err := tencentyun.GenPrivateMapKeyWithStringRoomID(
		tencentLib.TencentSdkAppID, tencentLib.TencentSdkAppSecret, userID, expirySecond, roomID, privilegeMap,
	)
	return sig, err
}
