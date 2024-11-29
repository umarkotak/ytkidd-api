package tencent_lib

import (
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentLib struct {
	Env                            string
	EnableAdvancePermissionControl bool

	// used for tencent cloud storage
	TencentCloudEndpoint  string
	TencentCloudSecretID  string
	TencentCloudSecretKey string

	// used for tencent rtc
	TencentSdkAppID     int
	TencentSdkAppSecret string

	TencentCloudClient *cos.Client
}

var (
	tencentLib TencentLib
)

func Initialize(t TencentLib) {
	tencentLib := TencentLib{
		Env:                            t.Env,
		EnableAdvancePermissionControl: t.EnableAdvancePermissionControl,
		TencentCloudEndpoint:           t.TencentCloudEndpoint,
		TencentCloudSecretID:           t.TencentCloudSecretID,
		TencentCloudSecretKey:          t.TencentCloudSecretKey,
		TencentSdkAppID:                t.TencentSdkAppID,
		TencentSdkAppSecret:            t.TencentSdkAppSecret,
	}

	tencentBucketUrl, err := url.Parse(tencentLib.TencentCloudEndpoint)
	if err != nil {
		logrus.Error(err)
	}

	tencentCloudBaseUrl := &cos.BaseURL{BucketURL: tencentBucketUrl}

	tencentCloudClient := cos.NewClient(tencentCloudBaseUrl, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  tencentLib.TencentCloudSecretID,
			SecretKey: tencentLib.TencentCloudSecretKey,
		},
	})

	tencentLib.TencentCloudClient = tencentCloudClient
}
