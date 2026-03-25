// Package minio provides repository layer implementation for S3 storage using minio library
package minio

import (
	"context"
	"fmt"
	"net/url"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	*minio.Client
}

// NewStorage accepts s3URL in the form:
//
//	s3://accessKey:secretKey@host:port/bucket?ssl=true
func NewStorage(_ context.Context, s3URL string) (*Storage, error) {
	u, err := url.Parse(s3URL)
	if err != nil {
		return nil, fmt.Errorf("minio: parse url: %w", err)
	}

	accessKey := u.User.Username()
	secretKey, _ := u.User.Password()
	endpoint := u.Host
	secure := u.Query().Get("ssl") == "true"

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})
	if err != nil {
		return nil, fmt.Errorf("minio: new client: %w", err)
	}

	return &Storage{Client: client}, nil
}
