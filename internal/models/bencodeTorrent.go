package models

import (
	"github.com/jackpal/bencode-go"
	"models/bencodeInfo"
)

type bencodeTorrent struct {
	Announce 		string `bencode:"announce"`
	Info 			bencodeInfo `bencode:"info"`
}