//go:generate go run github.com/golang/mock/mockgen@latest -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE

package repository

import (
	"context"

	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/gofrs/uuid"
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
