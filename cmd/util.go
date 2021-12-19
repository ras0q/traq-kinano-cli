package cmd

import "io"

type Writer interface {
	io.Writer
	SetChannelID(channelID string) Writer
	SetEmbed(embed bool) Writer
}
