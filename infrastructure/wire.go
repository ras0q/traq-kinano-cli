//go:generate go run github.com/google/wire/cmd/wire@latest
//go:build wireinject
// +build wireinject

package infrastructure

import (
	"context"

	"github.com/google/wire"
	"github.com/ras0q/traq-kinano-cli/cmd"
	"github.com/ras0q/traq-kinano-cli/ent"
	"github.com/ras0q/traq-kinano-cli/interfaces/external"
	"github.com/ras0q/traq-kinano-cli/interfaces/handler"
	"github.com/ras0q/traq-kinano-cli/interfaces/repository"
	"github.com/ras0q/traq-kinano-cli/usecases/service"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func InjectCmds(ctx context.Context, client *ent.Client, pl *payload.MessageCreated, q external.TraqAPI) *cmd.Cmds {
	wire.Build(
		cmd.NewCmds,
		handler.NewHandlers,
		repository.NewRepositories,
		service.NewServices,
	)

	return nil
}
