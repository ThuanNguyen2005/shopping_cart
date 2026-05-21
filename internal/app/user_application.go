package app

import (
	v1handler "23.usersapi/internal/handler/v1"
	"23.usersapi/internal/repository"
	"23.usersapi/internal/routes"
	v1routes "23.usersapi/internal/routes/v1"
	v1service "23.usersapi/internal/service/v1"
)

type UserModule struct {
	routes routes.Routers
}

func NewUserModule() Module {
	//initial repository
	repoUser := repository.NewInMemoryRepository()

	//initial service
	serviceUser := v1service.NewServiceUser(repoUser)

	//initial handler
	handlerUser := v1handler.NewHandlerUser(serviceUser)

	//initial routes
	routesUser := v1routes.NewRoutesUser(handlerUser)

	return &UserModule{
		routes: routesUser,
	}
}

func (m *UserModule) Routes() routes.Routers {
	return m.routes
}
