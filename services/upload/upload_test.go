package upload

import (
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/joho/godotenv"
)

var config *aws.Config

func setup() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config = &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("BACKBLAZE_KEY_ID"),
			os.Getenv("BACKBLAZE_APPLICATION_KEY"),
			"",
		),
		Endpoint:         aws.String(os.Getenv("AWS_ENDPOINT_URL")),
		Region:           aws.String(os.Getenv("AWS_REGION")),
		S3ForcePathStyle: aws.Bool(true),
	}
}

func shutdown() {
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestService(t *testing.T) {
	// s3Service := NewService(config)
	s3Service := service{
		config: config,
	}

	fileid := "1"
	bucketName := "auto-grader"
	filepath := "../../uploaded-files/31"

	err := s3Service.UploadFile(bucketName, fileid, filepath)
	if err != nil {
		t.Errorf("Uploading Failed: %v", err)
	}

	// Chekc if uploaded file has the same key
	files, err := s3Service.ListFiles(bucketName)
	if err != nil {
		t.Errorf("Could not list files of bucket: %v", err)
	}
	found := false
	for _, filename := range files {
		if filename == fileid {
			found = true
		}
	}
	if !found {
		t.Errorf("Could not find the file in the bucket")
	}
	temp := "temp"
	s3Service.DownloadFile(bucketName, fileid, temp)

	// Check if uploaded file is the same as the local file
	f1, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatal(err)
	}
	f2, err := os.ReadFile(temp)
	if err != nil {
		t.Fatal(err)
	}
	if string(f1) != string(f2) {
		t.Errorf("The uploaded file doesn't match the file in the s3")
	}
	os.Remove(temp)
}
