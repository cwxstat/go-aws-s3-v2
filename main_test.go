package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/xid"
	"s3-v2/pkg"
)

var client *s3.Client
var bucketName string

var runLiveTests = false

func init() {
	log.Println("Setting up suite")

	bucketName = "mybucket-" + (xid.New().String())
	fmt.Printf("Bucket name: %v\n", bucketName)

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
