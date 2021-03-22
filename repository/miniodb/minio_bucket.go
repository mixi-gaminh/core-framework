package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
	logger "github.com/mixi-gaminh/core-framework/logs"
)

// BucketIsExist - BucketIsExist
func (c *FileStorage) BucketIsExist(ctx context.Context, bucketID string) bool {
	exists, err := MinioClient.BucketExists(ctx, bucketID)
	if err != nil || !exists {
		logger.ERROR(err)
		return false
	}
	return true
}

// MakeBucket - MakeBucket
func (c *FileStorage) MakeBucket(ctx context.Context, bucketID string) error {
	// Make a new bucket if bucket is not exist.
	if err := MinioClient.MakeBucket(ctx, bucketID, minio.MakeBucketOptions{Region: minioLocation}); err != nil {
		logger.ERROR(err)
		return err
	}
	return nil
}

// RemoveBucket - RemoveBucket
func (c *FileStorage) RemoveBucket(ctx context.Context, userID, deviceID, bucketID string) ([]string, error) {
	listObjectsRemove, err := c.RemoveAllObjects(ctx, bucketID)
	if err != nil {
		return nil, err
	}
	if err := MinioClient.RemoveBucket(ctx, bucketID); err != nil {
		return nil, err
	}
	return listObjectsRemove, nil
}

// GetBucketPolicy - GetBucketPolicy
func (c *FileStorage) GetBucketPolicy(ctx context.Context, bucketID string) (string, error) {
	policy, err := MinioClient.GetBucketPolicy(ctx, bucketID)
	if err != nil {
		logger.ERROR(err)
		return "", err
	}
	return policy, nil
}

// SetBucketPolicy - SetBucketPolicy
func (c *FileStorage) SetBucketPolicy(ctx context.Context, bucketID, policy string) error {
	err := MinioClient.SetBucketPolicy(ctx, bucketID, policy)
	if err != nil {
		logger.ERROR(err)
		return err
	}
	return nil
}
