package minio

import (
	"context"
	"io"
	"log"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	logger "github.com/mixi-gaminh/core-framework/logs"
)

// PutObject - PutObject
func (c *FileStorage) PutObject(ctx context.Context, bucketID, contentType, path string, size int64, src io.Reader) (map[string]interface{}, error) {
	currentTime := time.Now().Local()
	timeUpload := currentTime.Format("02-01-2006 15-03-04")

	// size := fileStorage.Size
	path = strings.ReplaceAll(path, " ", "_")
	var opts minio.PutObjectOptions
	opts.ContentType = contentType
	// ioReader := bytes.NewReader(fileStorage.Source)
	// Upload the file with PutObject
	//go func(src io.Reader, bucketID, path string, size int64, opts minio.PutObjectOptions) {
	n, err := MinioClient.PutObject(ctx, bucketID, path, src, size, opts)
	if err != nil {
		logger.ERROR("Failed Upload the file with PutObject - err:", err)
	} else {
		log.Printf("Successfully uploaded %s of size %d\n", path, n.Size)
	}
	//}(src, bucketID, path, size, opts)
	infoMediaUpload := map[string]interface{}{
		"id":           path,
		"content_type": strings.Split(contentType, ";")[0],
		"path":         Domain + "/" + bucketID + "/" + path,
		"upload_time":  timeUpload,
	}
	return infoMediaUpload, nil
}

func (c *FileStorage) ListObjects(ctx context.Context, bucketID, prefix string) ([]string, error) {
	var listObjects []string
	// List all Object in Bucket
	objectCh := MinioClient.ListObjects(ctx, bucketID, minio.ListObjectsOptions{Prefix: prefix, Recursive: true})
	for object := range objectCh {
		if object.Err != nil {
			logger.ERROR(object.Err)
			return nil, object.Err
		}
		listObjects = append(listObjects, object.Key)
	}
	return listObjects, nil
}

// RemoveObject - RemoveObject
func (c *FileStorage) RemoveObject(ctx context.Context, bucketID, objectName string) error {
	if err := MinioClient.RemoveObject(ctx, bucketID, objectName, minio.RemoveObjectOptions{}); err != nil {
		logger.ERROR(err)
		return err
	}
	return nil
}

// RemoveAllObjects - RemoveAllObjects
func (c *FileStorage) RemoveAllObjects(ctx context.Context, bucketID string) ([]string, error) {
	var listObjectsRemove []string
	// Remove all Object in Bucket
	objectCh := MinioClient.ListObjects(ctx, bucketID, minio.ListObjectsOptions{Recursive: true})
	for object := range objectCh {
		if object.Err != nil {
			logger.ERROR(object.Err)
			return nil, object.Err
		}
		if err := c.RemoveObject(ctx, bucketID, object.Key); err != nil {
			return nil, err
		}
		listObjectsRemove = append(listObjectsRemove, object.Key)
	}
	return listObjectsRemove, nil
}

// CopyObject - CopyObject
func (c *FileStorage) CopyObject(ctx context.Context, bucketSrc, objectSrc, bucketDst, objectDst string) error {
	// Source object
	src := minio.CopySrcOptions{
		Bucket:             bucketSrc,
		Object:             objectSrc,
		MatchModifiedSince: time.Date(2014, time.April, 0, 0, 0, 0, 0, time.UTC),
	}

	// Destination object
	dst := minio.CopyDestOptions{
		Bucket: bucketDst,
		Object: objectDst,
	}

	// Initiate copy object.
	ui, err := MinioClient.CopyObject(ctx, dst, src)
	if err != nil {
		logger.ERROR(err)
		return err
	}

	log.Printf("Copied Object: %s sucessfully from Bucket: %s, to %s - UploadInfo %v\n", dst.Object, src.Bucket, dst.Bucket, ui)
	return nil
}

// StatObject - StatObject
func (c *FileStorage) StatObject(ctx context.Context, bucketID, objectName string) (interface{}, error) {
	stat, err := MinioClient.StatObject(ctx, bucketID, objectName, minio.StatObjectOptions{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return stat, nil
}
