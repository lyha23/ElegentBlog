package model

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/wejectchen/ginblog/utils"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func Check() (string, error) {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(utils.TencentSever)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  utils.TencentAccessKey, // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: utils.TencentSecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	opt := &cos.BucketGetOptions{
		Prefix:    "acg-img/", // prefix表示要查询的文件夹
		Delimiter: "/",        // deliter表示分隔符, 设置为/表示列出当前目录下的object, 设置为空表示列出所有的object
		//MaxKeys:   1000,       // 设置最大遍历出多少个对象, 一次listobject最大支持1000
	}
	v, _, err := client.Bucket.Get(context.Background(), opt)
	rand.Seed(time.Now().UnixNano())
	_url := fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, v.Contents[rand.Intn(len(v.Contents))].Key)
	return _url, err
}