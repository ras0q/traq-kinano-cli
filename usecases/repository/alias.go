//go:generate go run github.com/golang/mock/mockgen@latest -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE

package repository

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/ras0q/traq-kinano-cli/ent"
)

type AliasRepository interface {
	CallAlias(ctx context.Context, short string) (*ent.Alias, error)
	AddAlias(ctx context.Context, args *AddAliasArgs) (*ent.Alias, error)
}

type AddAliasArgs struct {
	UserID uuid.UUID
	Short  string
	Long   string
}

func NewAddAliasArgs(userID uuid.UUID, short string, long string) *AddAliasArgs {
	return &AddAliasArgs{
		UserID: userID,
		Short:  short,
		Long:   long,
	}
}
