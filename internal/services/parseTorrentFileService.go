package services

import (
	"internal/models"
	"github.com/jackpal/bencode-go"
)

func Open(r io.Reader) (*models.bencodeTorrent, error) {
	
	bto := bencodeTorrent{}
	err := bencode.Unmarshal(r, &bto)

	if err != nil{
		return nil, err
	}

	return &bto, nil
}