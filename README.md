# go-aws-s3-v2

main.go and main_test.go actively creates buckets. common_test.go 
mocks the test and does not actively create buckets, which is probably
what you'll use in your own code.


## Example Mock PutObject

```go
func TestMock(t *testing.T) {
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



```

## Example Mock GetObject

```go
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
```