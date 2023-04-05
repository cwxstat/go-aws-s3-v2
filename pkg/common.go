package pkg

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3CreateBucketAPI interface {
	CreateBucket(ctx context.Context,
		params *s3.CreateBucketInput,
		optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
}

func MakeBucket(ctx context.Context, client S3CreateBucketAPI, location string, name string) error {
	//snippet-start:[s3.go-v2.CreateBucket]
	// Create a bucket: We're going to create a bucket to hold content.
	// Best practice is to use the preset private access control list (ACL).
	// If you are not creating a bucket from us-east-1, you must specify a bucket location constraint.
	// bucket names must conform to several rules; read more at https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html

	if location == "us-east-1" {
		_, err := client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(name),
			ACL:    types.BucketCannedACLPrivate,
		})
		return err
	}

	_, err := client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket:                    aws.String(name),
		ACL:                       types.BucketCannedACLPrivate,
		CreateBucketConfiguration: &types.CreateBucketConfiguration{LocationConstraint: types.BucketLocationConstraint(location)},
	})

	if err != nil {
		return err
	}
	return nil

}

type S3ListBucketsAPI interface {
	ListBuckets(ctx context.Context,
		params *s3.ListBucketsInput,
		optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

type S3ListObjectsAPI interface {
	ListBuckets(ctx context.Context,
		params *s3.ListObjectsV2Input,
		optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

func AccountBucketOps(client s3.Client, name string) {

	fmt.Println("List buckets: ")
	//snippet-start:[s3.go-v2.ListBuckets]
	listBucketsResult, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

	if err != nil {
		panic("Couldn't list buckets")
	}

	for _, bucket := range listBucketsResult.Buckets {
		fmt.Printf("bucket name: %s\t\tcreated at: %v\n", *bucket.Name, bucket.CreationDate)
	}
	//snippet-end:[s3.go-v2.ListBuckets]

	//snippet-start:[s3.go-v2.ListObjects]
	// List objects in the bucket.
	// n.b. object keys in Amazon S3 do not begin with '/'. You do not need to lead your
	// prefix with it.
	fmt.Println("Listing the objects in the bucket:")
	listObjsResponse, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(name),
		Prefix: aws.String(""),
	})

	if err != nil {
		panic("Couldn't list bucket contents")
	}

	for _, object := range listObjsResponse.Contents {
		fmt.Printf("%s (%d bytes, class %v) \n", *object.Key, object.Size, object.StorageClass)
	}
	//snippet-end:[s3.go-v2.ListObjects]
}

type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

type S3GetObjectAPI interface {
	GetObject(ctx context.Context,
		params *s3.GetObjectInput,
		optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

type S3DeleteObjectAPI interface {
	DeleteObject(ctx context.Context,
		params *s3.DeleteObjectInput,
		optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
}

type S3CopyObjectAPI interface {
	CopyObject(ctx context.Context,
		params *s3.CopyObjectInput,
		optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error)
}

type S3ListObjectsV2API interface {
	ListObjectsV2(ctx context.Context,
		params *s3.ListObjectsV2Input,
		optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

type S3DeleteBuckeAPI interface {
	DeleteBucket(ctx context.Context,
		params *s3.DeleteBucketInput,
		optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
}

// NewPresignClient
type S3NewPresignClient interface {
	NewPresignClient(ctx context.Context,
		optFns ...func(*s3.Options)) (*s3.PresignClient, error)
}

func PutObject(ctx context.Context, client S3PutObjectAPI, bucket, key string, body io.Reader) (*s3.PutObjectOutput, error) {

	o, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
	})

	if err != nil {
		return o, err
	}
	return o, err
}

func GetObject(ctx context.Context, client S3GetObjectAPI, bucket, key string) (*s3.GetObjectOutput, error) {

	o, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return o, err
	}
	return o, err
}

func CopyObject(ctx context.Context, client S3CopyObjectAPI, bucket, key string) (*s3.CopyObjectOutput, error) {

	o, err := client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return o, err
	}
	return o, err
}

func DeleteObject(ctx context.Context, client S3DeleteObjectAPI, bucket, key string) (*s3.DeleteObjectOutput, error) {

	o, err := client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return o, err
	}
	return o, err
}

// TODO: (mmc) Need to break this up into smaller functions
func BucketOps(ctx context.Context, client s3.Client, name string) error {

	//snippet-start:[s3.go-v2.PutObject]
	// Place an object in a bucket.
	fmt.Println("Upload an object to the bucket")
	// Get the object body to upload.
	// Image credit: https://unsplash.com/photos/iz58d89q3ss
	stat, err := os.Stat("image.jpg")
	if err != nil {
		panic("Couldn't stat image: " + err.Error())
	}
	file, err := os.Open("image.jpg")

	if err != nil {
		panic("Couldn't open local file")
	}

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(name),
		Key:           aws.String("path/myfile.jpg"),
		Body:          file,
		ContentLength: stat.Size(),
	})

	file.Close()

	if err != nil {
		return err
	}

	//snippet-end:[s3.go-v2.PutObject]

	//snippet-start:[s3.go-v2.generate_presigned_url]
	// Get a presigned URL for the object.
	// In order to get a presigned URL for an object, you must
	// create a Presignclient
	fmt.Println("Create Presign client")
	presignClient := s3.NewPresignClient(&client)

	presignResult, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(name),
		Key:    aws.String("path/myfile.jpg"),
	})

	if err != nil {
		panic("Couldn't get presigned URL for GetObject")
	}

	fmt.Printf("Presigned URL For object: %s\n", presignResult.URL)

	//snippet-end:[s3.go-v2.generate_presigned_url]
	// Download the file.

	//snippet-start:[s3.go-v2.GetObject]
	fmt.Println("Download a file")
	getObjectResponse, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(name),
		Key:    aws.String("path/myfile.jpg"),
	})

	if err == nil {
		file, err = os.Create("download.jpg")

		if err != nil {
			return err
		}
		written, err := io.Copy(file, getObjectResponse.Body)
		if err != nil {
			panic("Failed to write file contents! " + err.Error())
		} else if written != getObjectResponse.ContentLength {
			panic("wrote a different size than was given to us")
		}
		fmt.Println("Done pulling file")
		file.Close()

	} else {
		panic("Couldn't download object")
	}
	//snippet-end:[s3.go-v2.GetObject]

	//snippet-start:[s3.go-v2.CopyObject]
	// Copy an object to another name.

	// CopyObject is "Pull an object from the source bucket + path".
	// The semantics of CopySource varies depending on whether you're using Amazon S3 on Outposts,
	// or through access points.
	// See https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html#API_CopyObject_RequestSyntax
	fmt.Println("Copy an object from another bucket to our bucket.")
	_, err = client.CopyObject(context.TODO(), &s3.CopyObjectInput{
		Bucket:     aws.String(name),
		CopySource: aws.String(name + "/path/myfile.jpg"),
		Key:        aws.String("other/file.jpg"),
	})

	if err != nil {
		return err
	}
	return nil
}

func BucketDelOps(client s3.Client, name string) {

	//snippet-start:[s3.go-v2.DeleteObject]
	// Delete a single object.
	fmt.Println("Delete an object from a bucket")
	_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(name),
		Key:    aws.String("other/file.jpg"),
	})
	if err != nil {
		panic("Couldn't delete object!")
	}

	//snippet-end:[s3.go-v2.DeleteObject]

	//snippet-start:[s3.go-v2.EmptyBucket]
	// Delete all objects in a bucket.

	fmt.Println("Delete the objects in a bucket")
	// Note: For versioned buckets, you must also delete all versions of
	// all objects within the bucket with ListVersions and DeleteVersion.
	listObjectsV2Response, err := client.ListObjectsV2(context.TODO(),
		&s3.ListObjectsV2Input{
			Bucket: aws.String(name),
		})

	for {

		if err != nil {
			panic("Couldn't list objects...")
		}
		for _, item := range listObjectsV2Response.Contents {
			fmt.Printf("- Deleting object %s\n", *item.Key)
			_, err = client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
				Bucket: aws.String(name),
				Key:    item.Key,
			})

			if err != nil {
				panic("Couldn't delete items")
			}
		}

		if listObjectsV2Response.IsTruncated {
			listObjectsV2Response, err = client.ListObjectsV2(context.TODO(),
				&s3.ListObjectsV2Input{
					Bucket:            aws.String(name),
					ContinuationToken: listObjectsV2Response.ContinuationToken,
				})
		} else {
			break
		}

	}
	//snippet-end:[s3.go-v2.EmptyBucket]

	// snippet-start:[s3.go-v2.DeleteBucket]
	fmt.Println("Delete a bucket")
	// Delete the bucket.

	_, err = client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String(name),
	})
	if err != nil {
		panic("Couldn't delete bucket: " + err.Error())
	}
	// snippet-end:[s3.go-v2.DeleteBucket]
}
