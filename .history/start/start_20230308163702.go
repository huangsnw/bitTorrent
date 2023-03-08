package main

import (
	"bitTorrent/torrentfile"
	"os"
)

func main() {
	torrentFilePath := os.Args[1]
	tf := torrentfile.OpenFile(torrentFilePath)
}
