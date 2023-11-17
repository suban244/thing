package upload

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Service interface {
	UploadFile(
		bucketName string,
		objectKey string,
		filepath string,
	) error

	DownloadFile(
		bucketName string,
		objectKey string,
		fileName string,
	) error
}

type service struct {
	config *aws.Config
}

func NewService(config *aws.Config) Service {
	return &service{
		config: config,
	}
}

func (s *service) createClient() (*s3.S3, error) {
	newSession, err := session.NewSession(s.config)
	if err != nil {
		log.Printf("Error while creating a new session: %v", err)
		return nil, err
	}
	client := s3.New(newSession)
	return client, nil
}

func (s *service) ListFiles(bucketName string) ([]string, error) {
	s3Client, err := s.createClient()
	if err != nil {
		return nil, err
	}
	result, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, err
	}

	keys := make([]string, len(result.Contents))
	for i, obj := range result.Contents {
		keys[i] = *obj.Key
	}
	return keys, nil
}

func (s *service) DownloadFile(
	bucketName string,
	objectKey string,
	fileName string,
) error {
	s3Client, err := s.createClient()
	if err != nil {
		return nil
	}
	result, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, objectKey, err)
		return err
	}
	defer result.Body.Close()
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Couldn't create file %v. Here's why: %v\n", fileName, err)
		return err
	}
	defer file.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Printf("Couldn't read object body from %v. Here's why: %v\n", objectKey, err)
	}
	_, err = file.Write(body)
	return err
}

func (s *service) UploadFile(bucket string, fileid string, filepath string) error {
	s3Client, err := s.createClient()
	if err != nil {
		return err
	}
	filedata, err := os.Open(filepath)
	if err != nil {
		log.Printf("Unable to open file: %s. %v", filepath, err)
		return err
	}
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Body:   filedata,
		Bucket: aws.String(bucket),
		Key:    aws.String(fileid),
	})

	if err != nil {
		fmt.Printf("Failed to upload object %s/%s, %s\n", bucket, fileid, err.Error())
		return err
	}
	return nil
}
