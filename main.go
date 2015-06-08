package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	svc := s3.New(nil)
	variableName := os.Args[1]
	bucketName := strings.Split(variableName, "/")[0]
	keyName := strings.Join(strings.Split(variableName, "/")[1:], "/")

	// make sure bucket exists
	params := &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	}
	_, err := svc.HeadBucket(params)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			var errMessage string
			// A service error occurred
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				errMessage = fmt.Sprintf("%v %v %v", reqErr.StatusCode(), variableName, reqErr.Message())
			} else {
				errMessage = fmt.Sprintf("%v %v", awsErr.Code(), awsErr.Message())
			}
			printAndExit(errMessage)
		} else {
			printAndExit(err.Error())
		}
	}

	getParams := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(keyName),
	}

	resp, err := svc.GetObject(getParams)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			errMessage := fmt.Sprintf("%s %s %s", awsErr.Code(), awsErr.Message(), variableName)
			printAndExit(errMessage)
		} else {
			printAndExit(err.Error())
		}
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		printAndExit(err.Error())
	}

	fmt.Print(string(contents))

}

func printAndExit(err string) {
	os.Stderr.Write([]byte(err))
	os.Exit(1)
}
