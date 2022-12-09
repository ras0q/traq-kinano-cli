package infrastructure

import (
	"context"
	"fmt"

	"github.com/ras0q/traq-kinano-cli/ent"
	"github.com/ras0q/traq-kinano-cli/util/config"
)

func NewEntClient() (*ent.Client, error) {
	// Setup ent client
	client, err := ent.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&collation=utf8mb4_general_ci",
		config.SQL.User,
		config.SQL.Pass,
		config.SQL.Host,
		config.SQL.Port,
		config.SQL.DBName,
	))
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
