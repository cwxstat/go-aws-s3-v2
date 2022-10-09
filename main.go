package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/xid"
	"s3-v2/common"
)

func main() {
	//snippet-start:[s3.go-v2.s3_basics]

	// This bucket name is 100% unique.
	// Remember that bucket names must be globally unique among all buckets.

	myBucketName := "mybucket-" + (xid.New().String())
	fmt.Printf("Bucket name: %v\n", myBucketName)

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("Failed to load configuration")
	}

	s3client := s3.NewFromConfig(cfg)

	ctx := context.TODO()

	err = common.MakeBucket(ctx, s3client, myBucketName)
	if err != nil {
		panic("Failed to create bucket")
	}
	// TODO: (mmc) make generic and add tests
	common.BucketOps(ctx, *s3client, myBucketName)
	common.AccountBucketOps(*s3client, myBucketName)
	common.BucketDelOps(*s3client, myBucketName)

}
