package torrentfile

import (
	"bitTorrent/utils"
	"os"

)

type TorrentFile struct {
}

func UnmarshalTorrentFile(filePath string) {
	f, err := os.Open(filePath)
	utils.Check(err)

}
