package main

import (
	"bitTorrent/torrentfile"
	"flag"
	"fmt"
	"log"
)

func main() {
	fmt.Println("程序开始了。")
	torrentFilePath := flag.String("p", "null", "Torrent文件路径")
	log.Printf("file path: %v", torrentFilePath)
	tf := torrentfile.OpenFile(*torrentFilePath)
	fmt.Println(tf.Announce)
}
