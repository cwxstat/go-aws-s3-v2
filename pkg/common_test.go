package pkg

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/middleware"
	"os"
	"testing"
)

type mockCreateBucket func(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)

func (m mockCreateBucket) CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
	return m(ctx, params, optFns...)
}

func TestMakeBucket(t *testing.T) {
	cases := []struct {
		client       func(t *testing.T) S3CreateBucketAPI
		name         string
		bucket       string
		description  string
		secretString string
		expect       []byte
	}{
		{
			client: func(t *testing.T) S3CreateBucketAPI {
				return mockCreateBucket(func(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
					t.Helper()
					if params.Bucket == nil {
						t.Errorf("expected name to be set")
					}
					if e, a := "bucketName", *params.Bucket; e != a {
						t.Errorf("expected %v, got %v", e, a)
					}
					return &s3.CreateBucketOutput{
						Location: aws.String("us-west-2"),
					}, nil
				})
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.TODO()
			err := MakeBucket(ctx, c.client(t), "bucketName")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}

type mockPutObject func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)

func (m mockPutObject) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	return m(ctx, params, optFns...)
}

func TestPutObject(t *testing.T) {
	cases := []struct {
		client        func(t *testing.T) S3PutObjectAPI
		name          string
		bucket        string
		key           string
		body          *os.File
		contentLength int64
		secretString  string
		expect        []byte
	}{
		{
			client: func(t *testing.T) S3PutObjectAPI {
				return mockPutObject(func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
					t.Helper()
					if params.Bucket == nil {
						t.Errorf("expected name to be set")
					}
					if e, a := "bucketName", *params.Bucket; e != a {
						t.Errorf("expected %v, got %v", e, a)
					}
					return &s3.PutObjectOutput{
						ETag: aws.String("etag"),
					}, nil
				})
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.TODO()
			_, err := PutObject(ctx, c.client(t), "bucketName", "key", nil)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}

}

type mockGetObject func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)

func (m mockGetObject) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return m(ctx, params, optFns...)
}
func TestGetObject(t *testing.T) {
	cases := []struct {
		client        func(t *testing.T) S3GetObjectAPI
		name          string
		bucket        string
		key           string
		body          *os.File
		contentLength int64
		secretString  string
		expect        []byte
	}{
		{
			client: func(t *testing.T) S3GetObjectAPI {
				return mockGetObject(func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
					t.Helper()
					if params.Bucket == nil {
						t.Errorf("expected name to be set")
					}
					if e, a := "bucketName", *params.Bucket; e != a {
						t.Errorf("expected %v, got %v", e, a)
					}
					return &s3.GetObjectOutput{
						ETag: aws.String("etag"),
					}, nil
				})
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.TODO()
			_, err := GetObject(ctx, c.client(t), "bucketName", "key")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}

}

type mockCopyObject func(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error)

func (m mockCopyObject) CopyObject(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error) {
	return m(ctx, params, optFns...)
}
func TestCopyObject(t *testing.T) {
	cases := []struct {
		client        func(t *testing.T) S3CopyObjectAPI
		name          string
		bucket        string
		key           string
		body          *os.File
		contentLength int64
		secretString  string
		expect        []byte
	}{
		{
			client: func(t *testing.T) S3CopyObjectAPI {
				return mockCopyObject(func(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error) {
					t.Helper()
					if params.Bucket == nil {
						t.Errorf("expected name to be set")
					}
					if e, a := "bucketName", *params.Bucket; e != a {
						t.Errorf("expected %v, got %v", e, a)
					}
					return &s3.CopyObjectOutput{
						BucketKeyEnabled:        false,
						CopyObjectResult:        nil,
						CopySourceVersionId:     nil,
						Expiration:              nil,
						RequestCharged:          "",
						SSECustomerAlgorithm:    nil,
						SSECustomerKeyMD5:       nil,
						SSEKMSEncryptionContext: nil,
						SSEKMSKeyId:             nil,
						ServerSideEncryption:    "",
						VersionId:               nil,
						ResultMetadata:          middleware.Metadata{},
					}, nil
				})
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.TODO()
			_, err := CopyObject(ctx, c.client(t), "bucketName", "key")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}

}

type mockDeleteObject func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)

func (m mockDeleteObject) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	return m(ctx, params, optFns...)
}

func TestDeleteObject(t *testing.T) {
	cases := []struct {
		client        func(t *testing.T) S3DeleteObjectAPI
		name          string
		bucket        string
		key           string
		body          *os.File
		contentLength int64
		secretString  string
		expect        []byte
	}{
		{
			client: func(t *testing.T) S3DeleteObjectAPI {
				return mockDeleteObject(func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
					t.Helper()
					if params.Bucket == nil {
						t.Errorf("expected name to be set")
					}
					if e, a := "bucketName", *params.Bucket; e != a {
						t.Errorf("expected %v, got %v", e, a)
					}
					return &s3.DeleteObjectOutput{
						DeleteMarker:   false,
						RequestCharged: "",
						VersionId:      nil,
						ResultMetadata: middleware.Metadata{},
					}, nil
				})
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.TODO()
			_, err := DeleteObject(ctx, c.client(t), "bucketName", "key")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}

}
