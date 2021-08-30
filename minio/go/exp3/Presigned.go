package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	ctx := context.Background()
	endpoint := "127.0.0.1:9000"
	accessKeyID := "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "mymusic"
	// location := "us-east-1"

	// Upload the zip file
	objectName := "test.tar"
	// filePath := "./test.tar"
	// contentType := "application/zip"

	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	reqParams.Set("ContentType", "application/zip")
	reqParams.Set("response-content-disposition", "attachment; filename=\"test.tar\"")

	// Generates a presigned url which expires in a day.
	presignedURL, err := minioClient.PresignedGetObject(ctx, bucketName, objectName, time.Second*24*60*60, reqParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully generated presigned URL", presignedURL)
	// Generates a url which expires in a day.
	expiry := time.Second * 24 * 60 * 60 // 1 day.
	presignedURLPut, err := minioClient.PresignedPutObject(ctx, bucketName, objectName, expiry)
	//presignedURL, err := minioClient.PresignedPutObject(context.Background(), "my-bucketname", "my-objectname", time.Duration(1000)*time.Second)
	//https://github.com/minio/minio-go/blob/master/examples/s3/presignedputobject.go

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully generated presigned URL", presignedURLPut)

}

//The request signature we calculated does not match the signature you provided. Check your key and signing method.</Message>
//http://docs.minio.org.cn/docs/master/golang-client-api-reference
//https://docs.aws.amazon.com/zh_cn/general/latest/gr/signature-v4-troubleshooting.html

/*
http://127.0.0.1:9000/mymusic/test.tar?ContentType=application%2Fzip&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIOSFODNN7EXAMPLE%2F20210830%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210830T045308Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3B%20filename%3D%22test.tar%22&X-Amz-Signature=21683ef38d323623a02846faaec51af6da99b9c2ef055bab965d21d697224034

http://172.17.0.2:9000/mymusic/test.tar?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=8YIT0LWXGS43759D8E2E%2F20210830%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210830T035057Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiI4WUlUMExXWEdTNDM3NTlEOEUyRSIsImV4cCI6MzYwMDAwMDAwMDAwMCwicG9saWN5IjoiY29uc29sZUFkbWluIn0.R2G7niQMuu9vYs9xrSb5GKJHnqIXvRzWMWUFJjD5_ZQVm2tfSEF4w-PAjmVv90MZWbyixHOzh0_e1VQOWN4Wsg&X-Amz-SignedHeaders=host&versionId=null&X-Amz-Signature=53e151d1f1090e1a34018690654abe259d2c3dcefaab5732ac188c4c169b48dd

*/
//http://www.allocmem.com/articles/8
//https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html
//https://www.cnblogs.com/woki/p/10342200.html
