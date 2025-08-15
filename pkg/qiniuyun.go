package pkg

import (
	"context"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"mime/multipart"
)

func QiNiuYun(file *multipart.FileHeader) (string, error) {
	accessKey := "b9kZ-QYKlCoM-0xR-FA8ZWUeIBvr6aRqcM9WDeN4"
	secretKey := "40CFvBU5zuK7cqoqEuZFSQ4ErW_ZnA76W-rRcgWY"
	mac := credentials.NewCredentials(accessKey, secretKey)
	bucket := "2301ayuanzhichao2"
	key := file.Filename
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	open, _ := file.Open()

	err := uploadManager.UploadReader(context.Background(), open, &uploader.ObjectOptions{
		BucketName: bucket,
		ObjectName: &key,
		CustomVars: map[string]string{
			"name": "github logo",
		},
		FileName: key,
	}, nil)
	return "http://t0o58f3jw.hn-bkt.clouddn.com/" + key, err
}
