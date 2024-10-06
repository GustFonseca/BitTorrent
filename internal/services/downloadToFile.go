package services

import {
	"os"
	"internal/models"
}


func (t *torrentFile) DownloadToFile(path string) error{
	
	var peerID [20]byte
	_, err := rand.Read(peerID[:])

	if err != nil{
		return err
	}

	peers, err := t.requestPeers(peerID, Port)

	if err != nil{
		return err
	}

	torrent := torrent{
		peers: peers,
		per
	}
}