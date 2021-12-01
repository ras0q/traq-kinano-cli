package handler

import (
	"context"

	"github.com/Ras96/traq-kinano-cli/usecases/repository"
	"github.com/Ras96/traq-kinano-cli/util/traq"
	"github.com/gofrs/uuid"
)

func (h *handlers) CallAlias(ctx context.Context, channelID string, short string) error {
	alias, err := h.Repo.CallAlias(ctx, short)
	if err != nil {
		return err
	}

	traq.MustPostMessage(channelID, alias.Long)

	return nil
}

func (h *handlers) AddAlias(ctx context.Context, channelID string, userID uuid.UUID, short string, long string) error {
	args := repository.NewAddAliasArgs(userID, short, long)

	if _, err := h.Repo.AddAlias(ctx, args); err != nil {
		return err
	}

	return nil
}
