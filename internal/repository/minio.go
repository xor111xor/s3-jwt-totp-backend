package repository

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/encrypt"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/domain"
)

type StorageMinio struct {
	conn          *minio.Client
	encryptPasswd string
	tmp_path      string
}

func NewStorageMinio(config domain.SysConfig) (*StorageMinio, error) {
	s3Client, err := minio.New(config.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.S3AccessKeyID, config.S3SecretAccessKey, ""),
		Secure: config.S3EnableSSL,
	})
	if err != nil {
		return nil, err
	}
	return &StorageMinio{
		conn:          s3Client,
		encryptPasswd: config.S3EncryptPasswd,
		tmp_path:      config.FileTmpPath,
	}, nil
}

func (s *StorageMinio) Upload(file domain.File) error {
	// Open a local file that we will upload
	filePath := fmt.Sprintf("%s/%s", s.tmp_path, file.Name)
	upload, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer upload.Close()

	// Get file stats.
	fstat, err := upload.Stat()
	if err != nil {
		return err
	}

	// New SSE-C where the cryptographic key is derived from a password and the objectname + bucketname as salt
	encryption := encrypt.DefaultPBKDF([]byte(s.encryptPasswd), []byte(file.Bucket+file.Name))

	// Encrypt file content and upload to the server
	_, err = s.conn.PutObject(context.Background(), file.Bucket, file.UuidFile.String(), upload, fstat.Size(), minio.PutObjectOptions{ServerSideEncryption: encryption})
	if err != nil {
		return err
	}

	return nil
}

func (s *StorageMinio) Download(file domain.File) error {
	// New SSE-C where the cryptographic key is derived from a password and the objectname + bucketname as salt
	encryption := encrypt.DefaultPBKDF([]byte(s.encryptPasswd), []byte(file.Bucket+file.Name))

	// Get the encrypted object
	reader, err := s.conn.GetObject(context.Background(), file.Bucket, file.UuidFile.String(), minio.GetObjectOptions{ServerSideEncryption: encryption})
	if err != nil {
		return err
	}
	defer reader.Close()

	// Local file which holds plain data
	filePath := fmt.Sprintf("%s/%s", s.tmp_path, file.Name)
	localFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer localFile.Close()

	if _, err := io.Copy(localFile, reader); err != nil {
		return err
	}
	return nil
}

// Remove object from bucket
func (s *StorageMinio) Remove(file domain.File) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err := s.conn.RemoveObject(context.Background(), file.Bucket, file.UuidFile.String(), opts)
	if err != nil {
		return err
	}
	return nil
}
