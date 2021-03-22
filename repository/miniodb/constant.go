package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	logger "github.com/mixi-gaminh/core-framework/logs"
)

// MinioDB - exported as symbol named "MinioDB"
// var MinioDB FileStorage

var MinioClient *minio.Client

var minioAccessKeyID string
var minioSecretAccessKey string
var minioLocation string

// MinIOHost - MinIOHost
var MinIOHost string

// Domain - Domain
var Domain string

// FileStorage - FileStorage
type FileStorage struct{}

// Admin - Admin
type Admin struct{}

// FileStorageConstructor -  FileStorageConstructor
func (c *FileStorage) FileStorageConstructor(_minIOHost, _minioEndpoint, _minioAccessKeyID, _minioSecretAccessKey, _minioLocation, _domain string, _minioUseSSL bool) {
	MinIOHost = _minIOHost
	minioAccessKeyID = _minioAccessKeyID
	minioSecretAccessKey = _minioSecretAccessKey
	minioLocation = _minioLocation
	Domain = _domain

	// Initialize minio client object.
	_minioClient, err := minio.New(_minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKeyID, minioSecretAccessKey, ""),
		Secure: _minioUseSSL,
	})
	if err != nil {
		panic(err)
	}
	MinioClient = _minioClient
	logger.NewLogger()
}

const minioAlias string = "vnptminio"
const (
	// SetAlias - MinIO server displays URL, access and secret keys.
	// Eg: mc alias set vnptminio http://203.162.141.37:9000 minioadmin minioadmin
	SetAlias string = "mc alias set " + minioAlias

	// GetAdminInfo - Display MinIO server information.
	// Eg: mc admin info vnptminio
	GetAdminInfo string = "mc admin info " + minioAlias

	// PolicyList - List all canned policies on MinIO.
	// Eg: mc admin policy list vnptminio
	PolicyList string = "mc admin policy list " + minioAlias

	// RemovePolicy - Remove policy 'listbucketsonly' on MinIO.
	// Eg: mc admin policy remove listbucketsonly
	RemovePolicy string = "mc admin policy remove"

	// ShowPolicy - Show info on a canned policy, 'writeonly'
	// Eg: mc admin policy info vnptminio writeinfo
	ShowPolicy string = "mc admin policy info " + minioAlias

	// AddNewUser - Add a new user 'newuser' on MinIO.
	// Eg: mc admin user add vnptminio username password
	AddNewUser string = "mc admin user add " + minioAlias

	// DisableUser - Disable a user 'newuser' on MinIO.
	// Eg: mc admin user disable vnptminio someuser1
	DisableUser string = "mc admin user disable " + minioAlias

	// EnableUser - Enable a user 'newuser' on MinIO.
	// Eg: mc admin user enable vnptminio someuser1
	EnableUser string = "mc admin user enable " + minioAlias

	// RemoveUser - Remove user 'newuser' on MinIO.
	// Eg: mc admin user remove vnptminio someuser1
	RemoveUser string = "mc admin user remove " + minioAlias

	// ListAllUser - List all users on MinIO.
	// Eg: mc admin user list --json vnptminio
	ListAllUser string = "mc admin user list --json " + minioAlias

	// InfoUser - Display info of a user
	// Eg: mc admin user info vnptminio someuser1
	InfoUser string = "mc admin user info " + minioAlias

	// AddUsersToGroup - Add a pair of users to a group 'somegroup' on MinIO.
	// Eg: mc admin group add vnptminio somegroup someuser1 someuser2
	AddUsersToGroup string = "mc admin group add " + minioAlias

	// RemoveUsersInGroup - Remove a pair of users from a group 'somegroup' on MinIO.
	// Eg: mc admin group remove vnptminio somegroup someuser1 someuser2
	RemoveUsersInGroup string = "mc admin group remove " + minioAlias

	// RemoveGroup - Remove a group 'somegroup' on MinIO. Only works if the given group is empty.
	// Eg: mc admin group remove vnptminio somegroup
	RemoveGroup string = "mc admin group remove " + minioAlias

	// InfoGroup - Get info on a group 'somegroup' on MinIO.
	// Eg: mc admin group info vnptminio somegroup
	InfoGroup string = "mc admin group info " + minioAlias

	// ListGroup - List all groups on MinIO.
	// Eg: mc admin group list myminio
	ListGroup string = "mc admin group list " + minioAlias

	// EnableGroup - Enable a group 'somegroup' on MinIO.
	// Eg: mc admin group enable myminio somegroup
	EnableGroup string = "mc admin group enable " + minioAlias

	// DisableGroup - Disable a group 'somegroup' on MinIO.
	// Eg: mc admin group disable myminio somegroup
	DisableGroup string = "mc admin group disable " + minioAlias

	// ListBucketQuota - List bucket quota on bucket 'mybucket' on MinIO.
	// mc admin bucket quota myminio/mybucket
	ListBucketQuota string = "mc admin bucket quota"

	// SetBucketQuota - Set a hard bucket quota of 64Mb for bucket 'mybucket' on MinIO.
	// Eg: mc admin bucket quota myminio/mybucket --hard 64MB
	SetBucketQuota string = "mc admin bucket quota"

	// ResetBucketQuota - Reset bucket quota configured for bucket 'mybucket' on MinIO.
	// Eg: mc admin bucket quota myminio/mybucket --clear
	ResetBucketQuota string = "mc admin bucket quota"
)
