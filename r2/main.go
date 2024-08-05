package r2

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type R2 struct {
	bucketName   string
	accountId    string
	accessKeyId  string
	accessSecret string
	publicUrl    string
	clientID     string
	symphonyUuid string
	dev          bool
}

var ctx = context.TODO()
var R2Client *s3.Client

func Init(licence map[string]interface{}) *R2 {
	// Cast the licence to the R2 struct
	// This is necessary to access the fields of the licence
	r2Licence := &R2{
		bucketName:   licence["BucketName"].(string),
		accountId:    licence["AccountId"].(string),
		accessKeyId:  licence["AccessKeyId"].(string),
		accessSecret: licence["AccessSecret"].(string),
		publicUrl:    licence["PublicUrl"].(string),
		clientID:     licence["ClientID"].(string),
		symphonyUuid: licence["SymphonyUuid"].(string),
		dev:          licence["Dev"].(bool),
	}
	A := R2{
		bucketName:   r2Licence.bucketName,
		accountId:    r2Licence.accountId,
		accessKeyId:  r2Licence.accessKeyId,
		accessSecret: r2Licence.accessSecret,
		publicUrl:    r2Licence.publicUrl,
		clientID:     r2Licence.clientID,
		symphonyUuid: r2Licence.symphonyUuid,
		dev:          r2Licence.dev,
	}
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", A.accountId),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(A.accessKeyId, A.accessSecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}

	R2Client = s3.NewFromConfig(cfg)
	return &A
}

func storeoath(ClientID string, StudyUuid string) string {
	return ClientID + "/" + StudyUuid
}

func (r *R2) upload(filePath string, StudyUuid string, StoreFolder string, ContentType string) (string, error) {
	if R2Client == nil {
		// R2Client is not assigned
		// Handle the error or return an error message
		return "", fmt.Errorf("client is not assigned")
	}

	// R2Client is assigned, continue with the upload logic
	// Add your upload code here
	// Open the file
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	// Get file info
	fileInfo, err := f.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to get file info: %v", err)
	}

	// Create the S3 object key
	key := fmt.Sprintf(storeoath(r.clientID, StudyUuid)+"/%s/%s", StoreFolder, fileInfo.Name())
	key = strings.ReplaceAll(key, " ", "_")
	// Create the S3 input object
	input := &s3.PutObjectInput{
		Bucket:      &r.bucketName,
		Key:         &key,
		Body:        f,
		ContentType: aws.String(ContentType),
	}
	if !r.dev {
		// Upload the file to S3
		_, err = R2Client.PutObject(ctx, input)
		if err != nil {
			return "", fmt.Errorf("failed to upload file: %v", err)
		}
	} else {
		fmt.Println("Dev mode")
		time.Sleep(1 * time.Second)
	}
	return r.publicUrl + "/" + key, nil
}

func (r *R2) UploadDCM(filePath string, StudyUuid string) (string, error) {
	return r.upload(filePath, StudyUuid, "DCM", "application/dicom")
}

func (r *R2) UploadIMG(filePath string, StudyUuid string) (string, error) {
	return r.upload(filePath, StudyUuid, "IMG", "image/jpeg")
}

func (r *R2) UploadZIP(filePath string, StudyUuid string) (string, error) {
	return r.upload(filePath, StudyUuid, "ZIP", "application/zip")
}

// https://pub-d0fcedc6472e46eb9eea69adbec2ea6f.r2.dev/ClientID/StudyUuid/DCM/FileName
