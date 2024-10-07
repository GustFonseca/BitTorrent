package main

import (
	"log"
	"os"
	"BitTorrent/internal/services"
)

func main() {

	inPath := os.Args[1]	
	outPath := os.Args[2]

	inFile, err := Open(inPath)

	if err != nil{
		log.Fatal(err)
	}

	err = DownloadToFile(outPath)

	if err != nil{
		log.Fatal(err)
	}
}