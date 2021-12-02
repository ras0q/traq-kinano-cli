package repository

import (
	"context"

	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/Ras96/traq-kinano-cli/ent/alias"
	"github.com/Ras96/traq-kinano-cli/usecases/repository"
	"github.com/Ras96/traq-kinano-cli/util/random"
)

func (r *repositories) CallAlias(ctx context.Context, short string) (*ent.Alias, error) {
	al, err := r.ent.Alias.
		Query().
		Where(alias.Short(short)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return al, nil
}

func (r *repositories) AddAlias(ctx context.Context, args *repository.AddAliasArgs) (*ent.Alias, error) {
	al, err := r.ent.Alias.
		Create().
		SetID(random.UUID()).
		SetUserID(args.UserID).
		SetShort(args.Short).
		SetLong(args.Long).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return al, nil
}
