package handler

import (
	"github.com/Ras96/traq-kinano-cli/util/traq"
)

func (h *handlers) Ping(channelID string) error {
	if _, err := traq.PostMessage(channelID, "pong!!!"); err != nil {
		return err
	}

	return nil
}
