package internal.services

import (	
	"BitTorrent/internal/internal/models"
)

func Unmarshal(peerBin []byte)([]peer, error)
{	
	const peerSize = 6

	numPeers := len(peerBin) / peerSize

	if len(peerBin) % peerSize != 0 {
		err := fmt.Errorf("received malformed peers")
		return nil, err
	}

	peers := make([]peer, numPeers)

	for i := 0; i < numPeers; i++ {
		offset := i * peerSize
		peers[i].IP = net.IP(peerBin[offset : offset+4]).String()
		peers[i].Port = binary.BigEndian.Uint16(peerBin[offset+4 : offset+6])
	}

	return peers, nil
}