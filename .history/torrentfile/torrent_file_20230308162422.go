package torrentfile

import (
	"bitTorrent/utils"
	"os"

	"github.com/jackpal/bencode-go"
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

// 反序列化文件
func OpenFile(filePath string) (TorrentFile, error) {
	torrentFile := TorrentFile{}
	file, err := os.Open(filePath)
	utils.Check(err)
	defer file.Close()
	err2 := bencode.Unmarshal(file, &torrentFile)
}
