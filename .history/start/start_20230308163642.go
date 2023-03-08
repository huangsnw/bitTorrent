package main

import (
	"bitTorrent/torrentfile"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("程序开始了。")
	torrentFilePath := os.Args[1]
	tf := torrentfile.OpenFile(torrentFilePath)
	log.Println(tf.Announce)
	log.Println(tf.InfoHash)
	log.Println(tf.Length)
	log.Println(tf.Name)
	log.Println(tf.PieceHashes)
	log.Println(tf.)
}
