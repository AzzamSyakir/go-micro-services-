package container

import (
	"go-micro-services/src/user-service/delivery/http"
)

type ControllerContainer struct {
	User *http.UserController
}

func NewControllerContainer(
	user *http.UserController,

) *ControllerContainer {
	controllerContainer := &ControllerContainer{
		User: user,
	}
	return controllerContainer
}