package minio

import (
	"log"
)

// AdminContruction -  AdminContruction
func (am *Admin) AdminContruction() error {
	args := []string{MinIOHost, minioAccessKeyID, minioSecretAccessKey}
	out, err := ExecCmd(SetAlias, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}

// GetAdminInfo - GetAdminInfo
func (am *Admin) GetAdminInfo() (string, error) {
	out, err := ExecCmd(GetAdminInfo, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(out)
	return out, nil
}

// PolicyList - PolicyList
func (am *Admin) PolicyList() (string, error) {
	out, err := ExecCmd(PolicyList, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(out)
	return out, nil
}

// RemovePolicy - RemovePolicy
func (am *Admin) RemovePolicy(policyName string) error {
	args := []string{policyName}
	out, err := ExecCmd(RemovePolicy, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}

// ShowPolicy - ShowPolicy
func (am *Admin) ShowPolicy(cannedPolicy string) (string, error) {
	args := []string{cannedPolicy}
	out, err := ExecCmd(ShowPolicy, args)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(out)
	return out, nil
}

// AddNewUser - AddNewUser
func (am *Admin) AddNewUser(username, password string) error {
	args := []string{username, password}
	out, err := ExecCmd(AddNewUser, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}

// DisableUser - DisableUser
func (am *Admin) DisableUser(username string) error {
	args := []string{username}
	out, err := ExecCmd(DisableUser, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}

// EnableUser - EnableUser
func (am *Admin) EnableUser(username string) error {
	args := []string{username}
	out, err := ExecCmd(EnableUser, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}

// RemoveUser - RemoveUser
func (am *Admin) RemoveUser(username string) error {
	args := []string{username}
	out, err := ExecCmd(RemoveUser, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}

// ListAllUser - ListAllUser
func (am *Admin) ListAllUser() (string, error) {
	out, err := ExecCmd(ListAllUser, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(out)
	return out, nil
}

// InfoUser - InfoUser
func (am *Admin) InfoUser(username string) (string, error) {
	args := []string{username}
	out, err := ExecCmd(InfoUser, args)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(out)
	return out, nil
}

// AddUsersToGroup - AddUsersToGroup
func (am *Admin) AddUsersToGroup(groupName string, users []string) error {
	args := []string{groupName}
	args = append(args, users...)
	out, err := ExecCmd(AddUsersToGroup, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}

// RemoveUsersInGroup - RemoveUsersInGroup
func (am *Admin) RemoveUsersInGroup(groupName string, users []string) error {
	args := []string{groupName}
	args = append(args, users...)
	out, err := ExecCmd(RemoveUsersInGroup, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}

// InfoGroup - InfoGroup
func (am *Admin) InfoGroup(groupName string) (string, error) {
	args := []string{groupName}
	out, err := ExecCmd(InfoGroup, args)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(out)
	return out, nil
}

// ListGroup - ListGroup
func (am *Admin) ListGroup() (string, error) {
	out, err := ExecCmd(ListGroup, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(out)
	return out, nil
}

// RemoveGroup - RemoveGroup
func (am *Admin) RemoveGroup(groupName string) error {
	args := []string{groupName}
	out, err := ExecCmd(RemoveGroup, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}

// EnableGroup - EnableGroup
func (am *Admin) EnableGroup(groupName string) error {
	args := []string{groupName}
	out, err := ExecCmd(EnableGroup, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}

// DisableGroup - DisableGroup
func (am *Admin) DisableGroup(groupName string) error {
	args := []string{groupName}
	out, err := ExecCmd(DisableGroup, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}

// ListBucketQuota - ListBucketQuota
func (am *Admin) ListBucketQuota(bucketName string) (string, error) {
	args := []string{minioAlias + "/" + bucketName}
	out, err := ExecCmd(ListBucketQuota, args)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(out)
	return out, nil
}

// SetBucketQuota - SetBucketQuota
func (am *Admin) SetBucketQuota(bucketName, quota string) error {
	args := []string{minioAlias + "/" + bucketName, "--hard", quota}
	out, err := ExecCmd(SetBucketQuota, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}

// ResetBucketQuota - ResetBucketQuota
func (am *Admin) ResetBucketQuota(bucketName string) error {
	args := []string{minioAlias + "/" + bucketName, "--clear"}
	out, err := ExecCmd(SetBucketQuota, args)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(out)
	return nil
}
