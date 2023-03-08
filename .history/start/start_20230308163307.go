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
	log.P("file path: %s", torrentFilePath)
	tf := torrentfile.OpenFile(*torrentFilePath)
	fmt.Println(tf.Announce)
}
