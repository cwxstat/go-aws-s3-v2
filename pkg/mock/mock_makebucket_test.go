package mock

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/middleware"
	"github.com/cwxstat/go-aws-s3-v2/pkg"
	"testing"
)

func TestMockClient_CreateBucket(t *testing.T) {

	m := NewMockCreateBucketClient()
	m.CreateBucketOutput(&s3.CreateBucketOutput{
		Location:       aws.String("us-east-2"),
		ResultMetadata: middleware.Metadata{},
	})

	err := pkg.MakeBucket(context.TODO(), m.ClientCreateBucket(),
		"us-east-2", "bucketName")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	output := m.GetCreateBucketOutput()
	if e, a := "us-east-2", *output.Location; e != a {
		t.Errorf("expected %v, got %v", e, a)
	}
	input := m.GetCreateBucketInput()
	if e, a := "bucketName", *input.Bucket; e != a {
		t.Errorf("expected %v, got %v", e, a)
	}

}
