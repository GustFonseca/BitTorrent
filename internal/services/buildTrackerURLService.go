package internal.services

import (	
	"BitTorrent/internal/models"
)

func (t *TorrentFile) BuildTrackerURL(peerID [20]byte, port uint16) (string, error) {
	// Faz parsing da URL base do announce
	base, err := url.Parse(t.Announce)
	if err != nil {
		return "", err
	}

	// Faz encoding do info_hash e do peer_id como URL-safe
	infoHashEncoded := url.QueryEscape(string(t.InfoHash[:]))
	peerIDEncoded := url.QueryEscape(string(peerID[:]))

	// Constrói os parâmetros da query
	params := url.Values{
		"info_hash":  []string{infoHashEncoded},
		"peer_id":    []string{peerIDEncoded},
		"port":       []string{strconv.Itoa(int(port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(t.Length)},
	}

	// Anexa os parâmetros à URL base
	base.RawQuery = params.Encode()
	return base.String(), nil
}