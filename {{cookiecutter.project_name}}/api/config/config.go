package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"{{cookiecutter.repo_location}}/{{cookiecutter.project_name}}/api/errors"
	"log"
)

type AppConfig struct {
	Session      *session.Session
}

func NewAppConfig() (*AppConfig, *errors.ApiError) {
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		log.Println(err)
		return nil, errors.ServerError(err)
	}
	return &AppConfig{
		Session: sess,
	},nil
}
