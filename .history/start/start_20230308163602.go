package main

import (
	"bitTorrent/torrentfile"
	"fmt"
	"os"
)

func main() {
	fmt.Println("程序开始了。")
	torrentFilePath := os.Args[1]
	tf := torrentfile.OpenFile(torrentFilePath)
	fmt.Println(tf.Announce)
	
}
