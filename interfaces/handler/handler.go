package handler

import (
	"github.com/Ras96/traq-kinano-cli/usecases/repository"
	"github.com/Ras96/traq-kinano-cli/usecases/service"
)

type Handlers interface {
	Ping(channelID string) error
}

type handlers struct {
	Repo repository.Repositories
	Srv  service.Services
}

func NewHandlers(repo repository.Repositories, srv service.Services) Handlers {
	return &handlers{
		Repo: repo,
		Srv:  srv,
	}
}
