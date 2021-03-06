package osser

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"grouper/common/conf"
)

type alioss struct {
	bucket *oss.Bucket
}

// AliyunGetBucket 获取一个桶子
func AliyunGetBucket(aly conf.AliyunOss) *alioss {
	// 创建OSSClient实例。
	client, err := oss.New(aly.Endpoint, aly.KeyID, aly.KeySecret)
	if err != nil {
		fmt.Println("报错:", err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(aly.BucketName)
	if err != nil {
		fmt.Println("报错:", err)
	}
	return &alioss{bucket: bucket}
}

// AliyunGoUpload 执行上传
func (ali *alioss) AliyunGoUpload(obj, locPth string) (ok bool) {
	// 上传文件，第一个参数是云文件，第二个参数是本地路径
	err := ali.bucket.PutObjectFromFile(obj, locPth)
	if err != nil {
		fmt.Println("报错:", err)
		return false
	} else {
		return true
	}
}
