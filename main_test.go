package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/smithy-go/middleware"
	"github.com/cwxstat/go-aws-s3-v2/pkg/mock"
	"io"
	"log"

	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cwxstat/go-aws-s3-v2/pkg"
	"github.com/rs/xid"
)

var client *s3.Client
var bucketName string

var runLiveTests = false

func init() {
	log.Println("Setting up suite")

	bucketName = "mybucket-" + (xid.New().String())
	fmt.Printf("bucket name: %v\n", bucketName)

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("Failed to load configuration")
	}

	client = s3.NewFromConfig(cfg)

}

func TestOps(t *testing.T) {
	if !runLiveTests {
		t.Skip("Skipping live test. Change variable runLiveTests to true to run.")
	}
	ctx := context.TODO()
	t.Log("Creating bucket...")
	pkg.MakeBucket(ctx, client, "us-east-1", bucketName)
	t.Log("Doing things to the bucket...")
	pkg.BucketOps(ctx, *client, bucketName)
	t.Log("list and such things being done to the bucket...")
	pkg.AccountBucketOps(*client, bucketName)
	t.Log("Cleaning up the bucket...")
	pkg.BucketDelOps(*client, bucketName)

}

func TestMock_PutObject(t *testing.T) {
	m := mock.NewMockPutObjectClient()
	dataToWrite := []byte("Hi!")

	m.ObjectOutput(&s3.PutObjectOutput{
		BucketKeyEnabled:     false,
		ETag:                 aws.String("etag"),
		RequestCharged:       "",
		ServerSideEncryption: "",
		ResultMetadata:       middleware.Metadata{},
	})

	result, err := pkg.PutObject(context.TODO(),
		m.ClientS3PutObject(),
		"bucketName", "key",
		bytes.NewReader(dataToWrite))
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Error("Expected a result")
	}

	data, err := m.GetData(make([]byte, 20))
	if string(data) != "Hi!" {
		t.Error("Expected data to be Hi!")
	}

}

func TestMock_GetObject(t *testing.T) {
	m := mock.NewMockGetObjectClient()

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

func TestMock_MakeBucket(t *testing.T) {
	m := mock.NewMockCreateBucketClient()
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
