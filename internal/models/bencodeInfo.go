package internal.models

import (
	"github.com/jackpal/bencode-go"
)

type bencodeInfo struct {
	Pieces 			string `bencode:"pieces"`
	PieceLength 	int `bencode:"piece length"`
	Length 			int `bencode:"length"`
	Name 			string `bencode:"name"`
}
