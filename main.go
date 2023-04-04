package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/xid"
	"s3-v2/pkg"
)

func main() {
	//snippet-start:[s3.go-v2.s3_basics]

	// This bucket name is 100% unique.
	// Remember that bucket names must be globally unique among all buckets.

	myBucketName := "mybucket-" + (xid.New().String())
	fmt.Printf("bucket name: %v\n", myBucketName)

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("Failed to load configuration")
	}

	s3client := s3.NewFromConfig(cfg)
	ctx := context.TODO()

	err = pkg.MakeBucket(ctx, s3client, cfg.Region, myBucketName)
	if err != nil {
		panic("Make bucket error: " + err.Error())
	}

	_, err = pkg.PutObject(ctx, s3client, myBucketName, "myobject", bytes.NewReader([]byte("Hi!")))
	if err != nil {
		panic("Failed to put object")
	}

	// TODO: (mmc) make generic and add tests
	//pkg.BucketOps(ctx, *s3client, myBucketName)
	//pkg.AccountBucketOps(*s3client, myBucketName)
	pkg.BucketDelOps(*s3client, myBucketName)

}
