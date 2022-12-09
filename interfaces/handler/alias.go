package handler

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/ras0q/traq-kinano-cli/ent"
	"github.com/ras0q/traq-kinano-cli/usecases/repository"
)

func (h *handlers) CallAlias(ctx context.Context, short string) (*ent.Alias, error) {
	alias, err := h.Repo.CallAlias(ctx, short)
	if err != nil {
		return nil, err
	}

	return alias, nil
}

func (h *handlers) AddAlias(ctx context.Context, userID uuid.UUID, short string, long string) error {
	args := repository.NewAddAliasArgs(userID, short, long)

	if _, err := h.Repo.AddAlias(ctx, args); err != nil {
		return err
	}

	return nil
}
