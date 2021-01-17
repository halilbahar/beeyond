package models

import (
	"fmt"
	"github.com/instrumenta/kubeval/kubeval"
)

type Content struct {
	Content string `form:"content" json:"content"`
}

func Validate(content string) string {
	var config *kubeval.Config = kubeval.NewDefaultConfig()

	contentBytes := []byte(content)
	_, err := kubeval.Validate(contentBytes, config)

	if err != nil {
		fmt.Printf("error -- %s \n", err)
		return err.Error()
	} else {
		return "no errors"
	}

}

func PrintError(err error) {
	if err != nil {
		panic(err)
	}
}
