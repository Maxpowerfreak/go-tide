package bencode

import (
	"fmt"
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	testTorrentFile, err := os.Open("testdata/test.torrent")
	if err != nil {
		t.Fatal(err)
	}

	torrent, err := Open(testTorrentFile)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(torrent.Announce)
	fmt.Println(torrent.Info.Name)
	fmt.Println(torrent.Info.Length)
	fmt.Println(torrent.Info.PieceLength)
}
