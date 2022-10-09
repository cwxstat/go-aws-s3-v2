# go-aws-s3-v2

main.go and main_test.go actively creates buckets. common_test.go 
mocks the test and does not actively create buckets, which is probably
what you'll use in your own code.


