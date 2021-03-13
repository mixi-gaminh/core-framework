package minio

import (
	"context"
	"io"
)

// IMinioAdmin - IMinioAdmin
type IMinioAdmin interface {
	AdminContruction() error

	GetAdminInfo() (string, error)

	PolicyList() (string, error)

	// RemovePolicy(policyName string) error
	RemovePolicy(string) error

	// ShowPolicy(cannedPolicy string) (string, error)
	ShowPolicy(string) (string, error)

	// AddNewUser(username, password string) error
	AddNewUser(string, string) error

	// DisableUser(username string) error
	DisableUser(string) error

	// EnableUser( string) error
	EnableUser(username string) error

	// RemoveUser(username string) error
	RemoveUser(string) error

	ListAllUser() (string, error)

	// InfoUser(username string) (string, error)
	InfoUser(string) (string, error)

	// AddUsersToGroup(groupName string, users []string) error
	AddUsersToGroup(string, []string) error

	// RemoveUsersInGroup(groupName string, users []string) error
	RemoveUsersInGroup(string, []string) error

	// InfoGroup(groupName string) (string, error)
	InfoGroup(string) (string, error)

	ListGroup() (string, error)

	// RemoveGroup(groupName string) error
	RemoveGroup(string) error

	// EnableGroup(groupName string) error
	EnableGroup(string) error

	// DisableGroup(groupName string) error
	DisableGroup(string) error

	// ListBucketQuota(bucketName string) (string, error)
	ListBucketQuota(string) (string, error)

	// SetBucketQuota(bucketName, quota string) error
	SetBucketQuota(string, string) error

	// ResetBucketQuota(bucketName string) error
	ResetBucketQuota(string) error
}

// IMinioStorage - IMinioStorage
type IMinioStorage interface {
	FileStorageConstructor()

	// BucketIsExist(ctx context.Context, bucketID string) bool
	BucketIsExist(context.Context, string) bool

	// MakeBucket(ctx context.Context, bucketID string) error
	MakeBucket(context.Context, string) error

	// RemoveBucket(ctx context.Context, userID, deviceID, bucketID string) ([]string, error)
	RemoveBucket(context.Context, string, string, string) ([]string, error)

	// GetBucketPolicy(ctx context.Context, bucketID string) (string, error)
	GetBucketPolicy(context.Context, string) (string, error)

	// SetBucketPolicy(ctx context.Context, bucketID, policy string) error
	SetBucketPolicy(context.Context, string, string) error

	// PutObject(ctx context.Context, bucketID, contentType, path string, size int64, src io.Reader) (map[string]interface{}, error)
	PutObject(context.Context, string, string, string, int64, io.Reader) (map[string]interface{}, error)

	// RemoveObject(ctx context.Context, bucketID, objectName string) error
	RemoveObject(context.Context, string, string) error

	// RemoveAllObjects(ctx context.Context, bucketID string) ([]string, error)
	RemoveAllObjects(context.Context, string) ([]string, error)

	// CopyObject(ctx context.Context, bucketSrc, objectSrc, bucketDst, objectDst string) error
	CopyObject(context.Context, string, string, string, string) error

	// StatObject(ctx context.Context, bucketID, objectName string) (interface{}, error)
	StatObject(context.Context, string, string) (interface{}, error)
}
