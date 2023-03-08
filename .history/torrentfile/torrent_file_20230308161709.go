package torrentfile

import (
	"bitTorrent/utils"
	"os"

	"gorm.io/gorm/utils"
)

type s


func UnmarshalTorrentFile(filePath string) {
	f, err := os.Open(filePath)
	utils.Check(err)

}
