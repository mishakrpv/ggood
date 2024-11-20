package di

import "github.com/ggood/ggood/pkg/db"

type ServiceContainer struct {
	DBContext db.Context
}

func NewServiceContainer(dbContext db.Context) *ServiceContainer {
	return &ServiceContainer{
		DBContext: dbContext,
	}
}
