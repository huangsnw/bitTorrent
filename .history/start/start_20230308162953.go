package main

import (
	"bitTorrent/torrentfile"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("程序开始了。")
	torrentFilePath := flag.String("path", "null", "Torrent文件路径")
	torrentfile.OpenFile(*torrentFilePath)

}
