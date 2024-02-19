package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/domain"
)

type PgConn struct {
	conn *pgxpool.Pool
	ctx  context.Context
}

func NewPgConnect(config string) (*PgConn, error) {
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, config)
	if err != nil {
		return nil, err
	}

	return &PgConn{
		ctx:  ctx,
		conn: conn,
	}, nil
}

func (p *PgConn) UserGetByMail(mail string) (domain.User, error) {
	log.Printf("Get user by mail: %s", mail)

	var user domain.User

	if err := p.conn.QueryRow(p.ctx,
		"SELECT uuid_user, mail, hash, date_reg, otp_cache FROM users WHERE mail = $1", mail).Scan(
		&user.UuidUser, &user.Mail, &user.PasswordHash, &user.CreatedAt, &user.OtpCache); err != nil {
		return user, err
	}

	user.Verified = true
	return user, nil
}

func (p *PgConn) UserCheckExistByMail(mail string) (bool, error) {
	user, err := p.UserGetByMail(mail)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return user.Mail == mail, nil
}

func (p *PgConn) UserAddOnRegistration(user domain.User) error {
	log.Printf("Add user with mail %s\n", user.Mail)
	_, err := p.conn.Exec(p.ctx,
		"INSERT INTO users (uuid_user, mail, hash, date_reg, otp_cache) VALUES ($1, $2, $3, $4, $5)", user.UuidUser, user.Mail, user.PasswordHash, time.Now(), user.OtpCache)
	return err
}

// insert file
func (p *PgConn) FileInsert(file domain.File) error {
	log.Printf("Add file for user %v\n", file.UuidUser)
	_, err := p.conn.Exec(p.ctx,
		"INSERT INTO files (uuid_file, file_name, upload_date, size, bucket_name, uuid_user) VALUES ($1, $2, $3, $4, $5, $6)",
		file.UuidFile, file.Name, file.UploadDate, file.Size, file.Bucket, file.UuidUser)
	return err
}

// delete file
func (p *PgConn) FileDelete(file domain.File) error {
	log.Printf("Remove file %v\n", file.UuidFile)
	_, err := p.conn.Exec(p.ctx,
		"DELETE FROM files WHERE uuid_file = $1", file.UuidFile.String())
	return err
}

// get file by id
func (p *PgConn) FileGetByID(fileID string) (domain.File, error) {
	log.Printf("Get file by id %s\n", fileID)
	var file domain.File

	if err := p.conn.QueryRow(p.ctx,
		"SELECT uuid_file, file_name, upload_date, size,  bucket_name, uuid_user FROM files WHERE uuid_file = $1",
		fileID).Scan(&file); err != nil {
		return file, err
	}
	return file, nil
}

// get all files for user
func (p *PgConn) FilesGetByUserID(uuidUser string) (*[]domain.File, error) {
	log.Printf("Get files from user %s", uuidUser)
	var files []domain.File

	rows, err := p.conn.Query(p.ctx, "SELECT * FROM files WHERE uuid_user = $1", uuidUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iter for each row
	for rows.Next() {
		var f domain.File
		err := rows.Scan(&f.UuidFile, &f.Name, &f.UploadDate, &f.Size, &f.Bucket, &f.UuidUser)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &files, nil
}
