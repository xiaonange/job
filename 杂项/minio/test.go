package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"log"
	"time"
)

func main() {
	endpoint := "voice.clk.center"
	accessKeyID := "XXU6KXP7Y83A4RSKQK26"
	secretAccessKey := "vvgGJgHNX+hoh0+JidHixbBMEGPZO5sOrKvExMAg"
	//useSSL := true

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, true)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%#v\n", minioClient) // minioClient初使化成功
	// 创建一个叫mymusic的存储桶。
	bucketName := "opplus"
	location := "us-east-1"
	err = minioClient.MakeBucket(bucketName, location)
	fmt.Print("111")
	fmt.Print(err)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			fmt.Print("错误")
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	// 上传一个zip文件。
	objectName := "/test/" + time.Now().String() + ".mp4"
	filePath := "D:/video/test.mp4"
	contentType := "application/mp4"

	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
