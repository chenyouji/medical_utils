package utils

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"time"
)

type Oss struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	Bucket          string `json:"bucket"`
	ReturnUrl       string `json:"return_url"`
}

func UploadOss(filename string, fd multipart.File, o *Oss) (string, error) {
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	ossClient, err := oss.New(o.Endpoint, o.AccessKeyID, o.AccessKeySecret)
	if err != nil {
		return "", err
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := ossClient.Bucket(o.Bucket)
	if err != nil {
		return "", err
	}

	now := time.Now()
	filepath := fmt.Sprintf("%d%d/", now.Year(), now.Month())

	// 将文件流上传至exampledir目录下的exampleobject.txt文件。
	err = bucket.PutObject("images/"+filepath, fd)
	if err != nil {
		return "", err
	}
	return o.ReturnUrl + filepath + filename, nil
}
