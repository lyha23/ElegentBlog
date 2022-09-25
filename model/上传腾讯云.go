package model

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/wejectchen/ginblog/utils"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

func Upload(_ multipart.File, file *multipart.FileHeader) (string, error) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse(utils.TencentSever)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  utils.TencentAccessKey, // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: utils.TencentSecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	timestamp := time.Now().Unix()
	filename := file.Filename
	if len(file.Filename) > 10 { // 文件名过长，截取
		filename = file.Filename[len(file.Filename)-7:]
	}
	name := fmt.Sprintf("myblog/%d%d/%d_%s", time.Now().Year(), time.Now().Month(), timestamp, filename)

	fd, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fd.Close()
	_, err = c.Object.Put(context.Background(), name, fd, nil)
	if err != nil {
		return "", err
	}
	savePath := utils.TencentSever + "/" + name
	return savePath, err
}