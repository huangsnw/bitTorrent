package torrentfile

import (
	"bitTorrent/utils"
	"os"
)

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

// 
func UnmarshalTorrentFile(filePath string) {
	f, err := os.Open(filePath)
	utils.Check(err)

}
