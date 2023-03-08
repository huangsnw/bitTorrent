package main

import (
	"bitTorrent/torrentfile"
	"flag"
	"fmt"
	"log"
)

func main() {
	fmt.Println("程序开始了。")
	torrentFilePath := flag.String("path", "null", "Torrent文件路径")
	log.Println("file path: %v",)
	tf := torrentfile.OpenFile(*torrentFilePath)
	fmt.Println(tf.Announce)
}
