package container

import (
	"fmt"
	"go-micro-services/src/user-service/config"
	httpdelivery "go-micro-services/src/user-service/delivery/http"
	"go-micro-services/src/user-service/delivery/http/route"
	"go-micro-services/src/user-service/repository"
	"go-micro-services/src/user-service/use_case"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type WebContainer struct {
	Env        *config.EnvConfig
	UserDB     *config.DatabaseConfig
	Repository *RepositoryContainer
	UseCase    *UseCaseContainer
	Controller *ControllerContainer
	Route      *route.RootRoute
}

func NewWebContainer() *WebContainer {
	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		panic(fmt.Errorf("error loading .env file: %w", errEnvLoad))
	}

	envConfig := config.NewEnvConfig()
	userDBConfig := config.NewUserDBConfig(envConfig)

	userRepository := repository.NewUserRepository()
	repositoryContainer := NewRepositoryContainer(userRepository)

	userUseCase := use_case.NewUserUseCase(userDBConfig, userRepository)

	useCaseContainer := NewUseCaseContainer(userUseCase)

	userController := httpdelivery.NewUserController(userUseCase)

	controllerContainer := NewControllerContainer(userController)

	router := mux.NewRouter()
	userRoute := route.NewUserRoute(router, userController)

	rootRoute := route.NewRootRoute(
		router,
		userRoute,
	)

	rootRoute.Register()

	webContainer := &WebContainer{
		Env:        envConfig,
		UserDB:     userDBConfig,
		Repository: repositoryContainer,
		UseCase:    useCaseContainer,
		Controller: controllerContainer,
		Route:      rootRoute,
	}

	return webContainer
}
