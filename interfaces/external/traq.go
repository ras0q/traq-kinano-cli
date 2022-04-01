package external

import (
	"github.com/antihax/optional"
	"github.com/gofrs/uuid"
)

type TraqAPI interface {
	PostMessage(channelID uuid.UUID, content string, embed bool) (err error)
}

type SearchMessagesOpts struct {
	// Word           optional.String
	After  optional.Time
	Before optional.Time
	// In             optional.Interface
	// To             optional.Interface
	// From           optional.Interface
	// Citation       optional.Interface
	Bot optional.Bool
	// HasURL         optional.Bool
	// HasAttachments optional.Bool
	// HasImage       optional.Bool
	// HasVideo       optional.Bool
	// HasAudio       optional.Bool
	Limit  optional.Int32
	Offset optional.Int32
	// Sort           optional.String
}
