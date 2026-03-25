package minio

import (
	"context"
	"fmt"

	"github.com/4nd3rs0n/oapi-test/pkg/errs"
	"github.com/minio/minio-go/v7"
)

func (s *Storage) SetupBucket(bucket string) error {
	if bucket == "" {
		return fmt.Errorf("%w: bucket", errs.ErrMissingArgument)
	}

	ctx := context.Background()

	exists, err := s.BucketExists(ctx, bucket)
	if err != nil {
		return mapErr(err)
	}
	if exists {
		return nil
	}

	return mapErr(
		s.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}),
	)
}

func (s *Storage) ClearBucket(ctx context.Context, bucket string) []error {
	if bucket == "" {
		return []error{fmt.Errorf("%w: bucket", errs.ErrMissingArgument)}
	}
	var out []error

	ch := s.ListObjects(ctx, bucket, minio.ListObjectsOptions{
		Recursive: true,
	})

	for obj := range ch {
		if obj.Err != nil {
			out = append(out, mapErr(obj.Err))
			continue
		}
		if err := s.RemoveObject(ctx, bucket, obj.Key, minio.RemoveObjectOptions{}); err != nil {
			out = append(out, mapErr(err))
		}
	}

	return out
}
