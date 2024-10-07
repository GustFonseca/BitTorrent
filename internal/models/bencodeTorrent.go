package internal.models

import (
	"github.com/jackpal/bencode-go"
	"BitTorrent/internal/models/bencodeInfo"
)

type bencodeTorrent struct {
	Announce 		string `bencode:"announce"`
	Info 			bencodeInfo `bencode:"info"`
}