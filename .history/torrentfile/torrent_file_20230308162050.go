package torrentfile

import (
	"bitTorrent/utils"
	"os"
)

// .torrent文件结构
type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

// 反序列化
func UnmarshalTorrentFile(filePath string) {
	f, err := os.Open(filePath)
	utils.Check(err)

}
