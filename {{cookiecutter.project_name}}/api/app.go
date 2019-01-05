package main


import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	config2 "github.com/basketsavings/{{cookiecutter.project_name}}/api/config"
	"github.com/basketsavings/{{cookiecutter.project_name}}/api/errors"
	"log"
)

type LambdaHandler func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

type AppHandler func(*gin.Context) *errors.ApiError

func (f AppHandler) WrapWithError(context *gin.Context) {
	err := f(context)
	if err != nil {
		context.AbortWithStatusJSON(err.HttpCode, err)
	}
}


type AppDependencies struct {
	AppConfig       *config2.AppConfig
}

type App struct {
	*AppDependencies
	router *ginadapter.GinLambda
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"message": "healthy"})
}

func (a *App) CreateRouter() {
	log.Printf("Gin cold start")
	
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("healthcheck", healthCheck)

	a.router = ginadapter.New(r)
}
func NewApp() (*App, error) {
	deps, err := CreateAppDependencies()
	if err != nil {
		return nil, err
	}
	app := &App{AppDependencies:deps}
	app.CreateRouter()
	return app, nil
}

func AppProxy() LambdaHandler {
	var app *App

	return func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		if app == nil {
			app, _ = NewApp()
		}
		return app.router.Proxy(req)
	}
}

func CreateAppDependencies() (*AppDependencies, error) {
	config, err := config2.NewAppConfig()

	if err != nil {
		return nil, err
	}

	return &AppDependencies{
		AppConfig:       config,
	}, nil
}