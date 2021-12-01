package repository

import (
	"context"

	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/Ras96/traq-kinano-cli/ent/alias"
	"github.com/Ras96/traq-kinano-cli/usecases/repository"
	"github.com/Ras96/traq-kinano-cli/util/random"
)

func (r *repositories) CallAlias(short string) (*ent.Alias, error) {
	al, err := r.ent.Alias.
		Query().
		Where(alias.Short(short)).
		First(context.Background()) //TODO: contextを引数にとる
	if err != nil {
		return nil, err
	}

	return al, nil
}

func (r *repositories) AddAlias(args *repository.AddAliasArgs) (*ent.Alias, error) {
	al, err := r.ent.Alias.
		Create().
		SetID(random.UUID()).
		SetUserID(args.UserID).
		SetShort(args.Short).
		SetLong(args.Long).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return al, nil
}
