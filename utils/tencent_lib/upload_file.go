package tencent_lib

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type (
	FileInfo struct {
		FilePath   string                // file path and objKey is the same. / symbol will be translated to folder automatically in tencent cos.
		File       multipart.File        //
		FileHeader *multipart.FileHeader //
		AclOpts    *cos.ACLHeaderOptions //
	}

	UploadResult struct {
		Path       string
		StatusCode int
	}
)

func UploadFile(ctx context.Context, fileInfo FileInfo) (UploadResult, error) {
	path := fmt.Sprintf("%s/%s", tencentLib.Env, fileInfo.FilePath)
	resp, err := tencentLib.TencentCloudClient.Object.Put(
		ctx,
		path,
		fileInfo.File,
		&cos.ObjectPutOptions{
			ACLHeaderOptions: fileInfo.AclOpts,
		},
	)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return UploadResult{}, err
	}

	return UploadResult{
		StatusCode: resp.StatusCode,
		Path:       path,
	}, nil
}

func GetObjectUrl(ctx context.Context, objKey string) (string, error) {
	url := tencentLib.TencentCloudClient.Object.GetObjectURL(
		fmt.Sprintf("%s/%s", tencentLib.Env, objKey),
	)

	return url.String(), nil
}

func GetPresignedObjectUrl(ctx context.Context, objKey string, exp time.Duration) (string, error) {
	presignedURL, err := tencentLib.TencentCloudClient.Object.GetPresignedURL(
		ctx,
		http.MethodGet,
		objKey,
		tencentLib.TencentCloudSecretID,
		tencentLib.TencentCloudSecretKey,
		exp,
		nil,
	)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return "", err
	}

	return presignedURL.String(), nil
}
