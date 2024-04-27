package container

import (
	"fmt"
	"go-micro-services/src/auth-service/config"
	httpdelivery "go-micro-services/src/auth-service/delivery/http"
	"go-micro-services/src/auth-service/delivery/http/route"
	"go-micro-services/src/auth-service/repository"
	"go-micro-services/src/auth-service/use_case"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type WebContainer struct {
	Env             *config.EnvConfig
	AuthDatabase    *config.DatabaseConfig
	ProductDatabase *config.DatabaseConfig
	OrderDatabase   *config.DatabaseConfig
	Repository      *RepositoryContainer
	UseCase         *UseCaseContainer
	Controller      *ControllerContainer
	Route           *route.RootRoute
}

func NewWebContainer() *WebContainer {
	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		panic(fmt.Errorf("error loading .env file: %w", errEnvLoad))
	}

	envConfig := config.NewEnvConfig()
	authDBConfig := config.NewAuthDBConfig(envConfig)

	authRepository := repository.NewAuthRepository()
	repositoryContainer := NewRepositoryContainer(authRepository)

	authUseCase := use_case.NewAuthUseCase(authDBConfig, authRepository, envConfig)
	exposeUseCase := use_case.NewExposeUseCase(envConfig)

	useCaseContainer := NewUseCaseContainer(authUseCase, exposeUseCase)

	authController := httpdelivery.NewAuthController(authUseCase)
	exposeController := httpdelivery.NewExposeController(exposeUseCase)

	controllerContainer := NewControllerContainer(authController, exposeController)

	router := mux.NewRouter()
	authRoute := route.NewAuthRoute(router, authController)
	// expose route
	userRoute := route.NewUserRoute(router, exposeController)
	productRoute := route.NewProductRoute(router, exposeController)
	categoryRoute := route.NewCategoryRoute(router, exposeController)
	orderRoute := route.NewOrderRoute(router, exposeController)

	rootRoute := route.NewRootRoute(
		router,
		authRoute,
	)
	exposeRoute := route.NewExposeRoute(
		router,
		userRoute,
		productRoute,
		categoryRoute,
		orderRoute,
	)

	rootRoute.Register()
	exposeRoute.Register()

	webContainer := &WebContainer{
		Env:          envConfig,
		AuthDatabase: authDBConfig,
		Repository:   repositoryContainer,
		UseCase:      useCaseContainer,
		Controller:   controllerContainer,
		Route:        rootRoute,
	}

	return webContainer
}
