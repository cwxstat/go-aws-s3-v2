package mock

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"s3-v2/pkg"
)

type MockCreateBucket func(ctx context.Context, params *s3.CreateBucketInput,
	optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)

func (m MockCreateBucket) CreateBucket(ctx context.Context, params *s3.CreateBucketInput,
	optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
	return m(ctx, params, optFns...)
}

type MockPutObject func(ctx context.Context, params *s3.PutObjectInput,
	optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)

func (m MockPutObject) PutObject(ctx context.Context, params *s3.PutObjectInput,
	optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	return m(ctx, params, optFns...)
}

type MockClient struct {
	bucket          *string
	putObjectOutput *s3.PutObjectOutput
	putObjectInput  *s3.PutObjectInput
	err             error
}

func NewMockClient() *MockClient {
	return &MockClient{}
}
func (m *MockClient) GetBucket() *string {
	return m.bucket
}

func (m *MockClient) ObjectOutput(output *s3.PutObjectOutput) {
	m.putObjectOutput = output
}

func (m *MockClient) GetPutObjectOutput() *s3.PutObjectOutput {
	return m.putObjectOutput
}

func (m *MockClient) GetPutObjectInput() *s3.PutObjectInput {
	return m.putObjectInput
}

func (m *MockClient) GetData(b []byte) ([]byte, error) {
	n, err := m.putObjectInput.Body.Read(b)
	return b[:n], err
}

func (m *MockClient) ClientS3PutObject() pkg.S3PutObjectAPI {
	return MockPutObject(func(ctx context.Context, params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
		m.bucket = params.Bucket
		m.putObjectInput = params
		return m.putObjectOutput, m.err
	})

}
