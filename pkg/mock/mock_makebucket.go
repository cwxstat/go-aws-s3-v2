package mock

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cwxstat/go-aws-s3-v2/pkg"
)

type MockCreateBucket func(ctx context.Context, params *s3.CreateBucketInput,
	optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)

func (m MockCreateBucket) CreateBucket(ctx context.Context, params *s3.CreateBucketInput,
	optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
	return m(ctx, params, optFns...)
}

type MockClientCreateBucket struct {
	bucket             *string
	createBucketOutput *s3.CreateBucketOutput
	createBucketInput  *s3.CreateBucketInput
	err                error
}

func NewMockCreateBucketClient() *MockClientCreateBucket {
	return &MockClientCreateBucket{}
}
func (m *MockClientCreateBucket) GetBucket() *string {
	return m.bucket
}

func (m *MockClientCreateBucket) CreateBucketOutput(output *s3.CreateBucketOutput) {
	m.createBucketOutput = output
}

func (m *MockClientCreateBucket) CreateBucketInput(input *s3.CreateBucketInput) {
	m.createBucketInput = input
}

func (m *MockClientCreateBucket) GetCreateBucketOutput() *s3.CreateBucketOutput {
	return m.createBucketOutput
}

func (m *MockClientCreateBucket) GetCreateBucketInput() *s3.CreateBucketInput {
	return m.createBucketInput
}

// S3CreateBucketAPI
func (m *MockClientCreateBucket) ClientCreateBucket() pkg.S3CreateBucketAPI {
	return MockCreateBucket(func(ctx context.Context, params *s3.CreateBucketInput,
		optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
		m.bucket = params.Bucket
		m.createBucketInput = params
		return m.createBucketOutput, m.err
	})

}
