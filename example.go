package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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

	r.GET("/accounts", func(c *gin.Context) {
		data, err := ioutil.ReadFile("./testdata/accounts.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var metadata Metadata
		err = json.Unmarshal(data, &metadata)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, metadata.Accounts)

	})

	r.Run()
}
