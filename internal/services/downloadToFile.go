package internal.services

import {
	"os"
	"BitTorrent/internal/models"
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
		Peers 		: peers,
		PeerId		: peerID,
		InfoHash	: t.InfoHash,
		PieceHashes	: t.PieceHashes,
		PieceLength	: t.PieceLength,
		Length 		: t.Length,
		Name 		: t.Name
	}

	buf, err := torrent.Download()

	if err != nil{
		return err
	}

	outFile, err := os.Create(path)

	if err != nil{
		return err
	}

	_, err = outFile.Write(buf)
	
	if err != nil{
		return err
	}

	defer outFile.Close()

	_, err = outFile.Write(buf)
	
	if err != nil{
		return err
	}	

	return nil

}