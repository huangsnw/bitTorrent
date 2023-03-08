package main

import (
	"bitTorrent/torrentfile"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("程序开始了。")
	torrentFilePath := os.Args[0]
	log.Printf("file path: %v", torrentFilePath)
	tf := torrentfile.OpenFile(torrentFilePath)
	fmt.Println(tf.Announce)
}
