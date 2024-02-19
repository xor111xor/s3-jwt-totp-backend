package cmd

import (
	"github.com/xor111xor/s3-jwt-totp-backend/internal/domain"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/repository"
)

// DB persistent
func GetRepoDB(config string) (domain.RepoDB, error) {
	repo, err := repository.NewPgConnect(config)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

// cache of current session
func GetCache() domain.Cache {
	cache := repository.NewInMemoryCache()
	return cache
}

// S3 compatible
func GetStorageS3(config domain.SysConfig) (domain.StorageS3, error) {
	repo, err := repository.NewStorageMinio(config)
	if err != nil {
		return nil, err
	}
	return repo, nil
}
