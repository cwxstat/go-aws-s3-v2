package mock

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cwxstat/go-aws-s3-v2/pkg"
)

type MockGetObject func(ctx context.Context, params *s3.GetObjectInput,
	optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)

func (m MockGetObject) GetObject(ctx context.Context, params *s3.GetObjectInput,
	optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return m(ctx, params, optFns...)
}

type MockClientGetObject struct {
	bucket          *string
	getObjectOutput *s3.GetObjectOutput
	getObjectInput  *s3.GetObjectInput
	err             error
}

func NewMockGetObjectClient() *MockClientGetObject {
	return &MockClientGetObject{}
}
func (m *MockClientGetObject) GetBucket() *string {
	return m.bucket
}

func (m *MockClientGetObject) ObjectOutput(output *s3.GetObjectOutput) {
	m.getObjectOutput = output
}

func (m *MockClientGetObject) ObjectInput(input *s3.GetObjectInput) {
	m.getObjectInput = input
}

func (m *MockClientGetObject) GetPutObjectOutput() *s3.GetObjectOutput {
	return m.getObjectOutput
}

func (m *MockClientGetObject) GetPutObjectInput() *s3.GetObjectInput {
	return m.getObjectInput
}

func (m *MockClientGetObject) GetData(b []byte) ([]byte, error) {
	n, err := m.getObjectOutput.Body.Read(b)
	return b[:n], err
}

func (m *MockClientGetObject) ClientS3GetObject() pkg.S3GetObjectAPI {
	return MockGetObject(func(ctx context.Context, params *s3.GetObjectInput,
		optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
		m.bucket = params.Bucket
		m.getObjectInput = params
		return m.getObjectOutput, m.err
	})

}
