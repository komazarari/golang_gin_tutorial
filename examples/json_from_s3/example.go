package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

type Account struct {
	AccountId string `json:"accountId"`
	IamAlias  string `json:"iamAlias"`
}

type Metadata struct {
	Accounts []Account `json:"accounts"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	bucketName := "YOUR-BUCKET-NAME"
	objectKey := "accounts.json"

	r.GET("/accounts", func(c *gin.Context) {

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("ap-northeast-1"),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		svc := s3.New(sess)
		res, err := svc.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var metadata Metadata
		err = json.NewDecoder(res.Body).Decode(&metadata)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, metadata.Accounts)

	})

	r.Run()
}
