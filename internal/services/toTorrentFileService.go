package internal.services

import(
	"BitTorrent/internal/internal/models"
)

func (bto *bencodeTorrent) ToTorrentFile() (TorrentFile, error){
	infoHash, err := bto.infoHash.hash()
	
}