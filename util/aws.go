package util

import (
	"log"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func PresignURL(s3Object string) string {
	// Separate bucket and key
	bucket, key := bucketAndKey(s3Object)

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal("failed to load config")
	}

	svc := s3.New(cfg)

	req := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	urlStr, err := req.Presign(24 * 7 * time.Hour) // Max aws duration for IAM user

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	return urlStr
}

func bucketAndKey(s3Url string) (string, string) {
	u, _ := url.Parse(s3Url)
	//fmt.Printf("proto: %q, bucket: %q, key: %q", u.Scheme, u.Host, u.Path)
	return u.Host, u.Path
}
