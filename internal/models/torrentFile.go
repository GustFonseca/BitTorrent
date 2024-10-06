package models

type torrentFile struct {
	Announce string
	InfoHash [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length int
	Name string
}

type torrent struct {
	Peers []peer.peer
	PeerId [20]byte
	InfoHash [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length int
	Name string
}

type pieceWork struct {
	Index int
	Hash [20]byte
	Length int
}

type pieceResult struct {
	Index int
	Buf []byte
}

type pieceProgress struct {
	Index int
	client *client.client
	buf []byte
	downloaded int
	requested int
	backlog int
}