package container

import (
	"go-micro-services/services/Order/repository"
)

type RepositoryContainer struct {
	Order *repository.OrderRepository
}

func NewRepositoryContainer(
	order *repository.OrderRepository,

) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		Order: order,
	}
	return repositoryContainer
}
