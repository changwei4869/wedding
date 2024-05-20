package minio

import (
	"context"
	"fmt"
	"io"

	"github.com/changwei4869/wedding/utils"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// NewMinIOClient initializes a new MinIO client
func NewMinIOClient() (*minio.Client, error) {
	minioClient, err := minio.New(utils.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(utils.AccessKeyID, utils.SecretAccessKey, ""),
		Secure: utils.UseSSL,
	})
	return minioClient, err
}
func UploadFileToMinIO(client *minio.Client, bucketName, objectName string, file io.Reader, fileSize int64, contentType string) (minio.UploadInfo, error) {
	// Ensure the bucket exists and has the appropriate policies
	if exists, err := client.BucketExists(context.Background(), bucketName); err != nil {
		return minio.UploadInfo{}, err
	} else if !exists {
		if err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "us-east-1"}); err != nil {
			return minio.UploadInfo{}, err
		}

		policy := `{
			"Version": "2012-10-17",
			"Statement": [{
				"Effect": "Allow",
				"Principal": {
					"AWS": ["*"]
				},
				"Action": ["s3:GetBucketLocation", "s3:ListBucket", "s3:ListBucketMultipartUploads"],
				"Resource": ["arn:aws:s3:::%s"]
			}, {
				"Effect": "Allow",
				"Principal": {
					"AWS": ["*"]
				},
				"Action": ["s3:AbortMultipartUpload", "s3:DeleteObject", "s3:GetObject", "s3:ListMultipartUploadParts", "s3:PutObject"],
				"Resource": ["arn:aws:s3:::%s/*"]
			}]
		}`
		if err := client.SetBucketPolicy(context.Background(), bucketName, fmt.Sprintf(policy, bucketName, bucketName)); err != nil {
			return minio.UploadInfo{}, err
		}
	}

	// Upload the file
	uploadInfo, err := client.PutObject(context.Background(), bucketName, objectName, file, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return uploadInfo, nil
}
