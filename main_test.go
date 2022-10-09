package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/xid"
	"s3-v2/common"
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
	common.MakeBucket(ctx, client, bucketName)
	t.Log("Doing things to the bucket...")
	common.BucketOps(ctx, *client, bucketName)
	t.Log("list and such things being done to the bucket...")
	common.AccountBucketOps(*client, bucketName)
	t.Log("Cleaning up the bucket...")
	common.BucketDelOps(*client, bucketName)

}
