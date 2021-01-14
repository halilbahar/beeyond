package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/instrumenta/kubeval/kubeval"
	"net/http"
)

type Content struct  {
	Content string `form:"content" json:"content"`
}

func main() {


}

func GetValidationResult(c *gin.Context) {
	var json Content
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"results": validate(json.Content),
	})
}

func validate(content string) string {
	var config* kubeval.Config = kubeval.NewDefaultConfig()

	contentBytes := []byte(content)
	_, err := kubeval.Validate(contentBytes, config)

	if err != nil{
		fmt.Printf("error -- %s \n",  err)
		return err.Error()
	} else {
		return "no errors"
	}

}

func PrintError(err error) {
	if err != nil{
		panic(err)
	}
}