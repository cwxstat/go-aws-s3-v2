package mock

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/middleware"
	"github.com/cwxstat/go-aws-s3-v2/pkg"
	"io"
	"testing"
)

func TestMockClient_ClientS3GetObject(t *testing.T) {

	m := NewMockGetObjectClient()

	dataToWrite := []byte("Hi!")
	m.ObjectOutput(&s3.GetObjectOutput{
		BucketKeyEnabled: false,
		ETag:             aws.String("etag"),
		Body:             io.NopCloser(bytes.NewReader(dataToWrite)),
		ResultMetadata:   middleware.Metadata{},
	})

	output, err := pkg.GetObject(context.TODO(), m.ClientS3GetObject(), "bucketName", "key")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if e, a := "etag", *output.ETag; e != a {
		t.Errorf("expected %v, got %v", e, a)
	}
	b := make([]byte, 20)
	data, err := m.GetData(b)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if e, a := "Hi!", string(data); e != a {
		t.Errorf("expected %v, got %v", e, a)
	}

}
