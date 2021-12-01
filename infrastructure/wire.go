//go:generate go run github.com/google/wire/cmd/wire@latest
//go:build wireinject
// +build wireinject

package infrastructure

import (
	"github.com/Ras96/traq-kinano-cli/cmd"
	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/Ras96/traq-kinano-cli/interfaces/handler"
	"github.com/Ras96/traq-kinano-cli/interfaces/repository"
	"github.com/Ras96/traq-kinano-cli/usecases/service"
	"github.com/google/wire"
	traqbot "github.com/traPtitech/traq-bot"
)

func injectCmds(client *ent.Client, pl *traqbot.MessageCreatedPayload) *cmd.Cmds {
	wire.Build(
		cmd.NewCmds,
		handler.NewHandlers,
		repository.NewRepositories,
		service.NewServices,
	)

	return nil
}
