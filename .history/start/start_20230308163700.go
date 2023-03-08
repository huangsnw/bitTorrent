package main

import (
	"bitTorrent/torrentfile"
	"log"
	"os"
)

func main() {
	torrentFilePath := os.Args[1]
	tf := torrentfile.OpenFile(torrentFilePath)

	log.Println(tf.Announce)
	log.Println(tf.InfoHash)
	log.Println(tf.Length)
	log.Println(tf.Name)
	log.Println(tf.PieceHashes)
	log.Println(tf.PieceLength)
}
