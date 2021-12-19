//go:generate go run github.com/google/wire/cmd/wire@latest
//go:build wireinject
// +build wireinject

package infrastructure

import (
	"context"

	"github.com/Ras96/traq-kinano-cli/cmd"
	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/Ras96/traq-kinano-cli/interfaces/handler"
	"github.com/Ras96/traq-kinano-cli/interfaces/repository"
	"github.com/Ras96/traq-kinano-cli/usecases/service"
	"github.com/google/wire"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func InjectCmds(ctx context.Context, client *ent.Client, pl *payload.MessageCreated) *cmd.Cmds {
	wire.Build(
		cmd.NewCmds,
		handler.NewHandlers,
		repository.NewRepositories,
		service.NewServices,
	)

	return nil
}
