package bencode

import (
	"io"

	"github.com/jackpal/bencode-go"
)

type info struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

type torrent struct {
	Announce string `bencode:"announce"`
	Info     info   `bencode:"info"`
}

// Open parses a torrent file and returns a torrent struct
func Open(r io.Reader) (*torrent, error) {
	t := &torrent{}
	err := bencode.Unmarshal(r, t)

	if err != nil {
		return nil, err
	}

	return t, err
}
