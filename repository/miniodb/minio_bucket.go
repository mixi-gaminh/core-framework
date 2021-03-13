package minio

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
)

// BucketIsExist - BucketIsExist
func (c *FileStorage) BucketIsExist(ctx context.Context, bucketID string) bool {
	exists, err := minioClient.BucketExists(ctx, bucketID)
	if err != nil || !exists {
		log.Println(err)
		return false
	}
	return true
}

// MakeBucket - MakeBucket
func (c *FileStorage) MakeBucket(ctx context.Context, bucketID string) error {
	// Make a new bucket if bucket is not exist.
	if err := minioClient.MakeBucket(ctx, bucketID, minio.MakeBucketOptions{Region: minioLocation}); err != nil {
		log.Println(err)
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
	if err := minioClient.RemoveBucket(ctx, bucketID); err != nil {
		return nil, err
	}
	return listObjectsRemove, nil
}

// GetBucketPolicy - GetBucketPolicy
func (c *FileStorage) GetBucketPolicy(ctx context.Context, bucketID string) (string, error) {
	policy, err := minioClient.GetBucketPolicy(ctx, bucketID)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return policy, nil
}

// SetBucketPolicy - SetBucketPolicy
func (c *FileStorage) SetBucketPolicy(ctx context.Context, bucketID, policy string) error {
	err := minioClient.SetBucketPolicy(ctx, bucketID, policy)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
