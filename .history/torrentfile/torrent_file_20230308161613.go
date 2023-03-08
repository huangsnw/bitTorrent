package torrentfile

import "os"

func UnmarshalTorrentFile(filePath string) {
	f, err := os.Open(filePath)
	
}
