package nfl

import (
	"strings"

	bridge "github.com/IAvellanedaGit/nuevo-platform-mobile-bridge"
	afc "github.com/IAvellanedaGit/sandbox-avellaneda-pm-nfl-afc/pkg"
	nfc "github.com/IAvellanedaGit/sandbox-avellaneda-pm-nfl-nfc/pkg"
)

type Service struct {
	b      *bridge.Bridge
	logger bridge.Logger

	cAFC *afc.Client
	cNFC *nfc.Client
}

func NewService(b *bridge.Bridge, logger bridge.Logger) *Service {

	cAFC, err := afc.NewClient(b)
	if err != nil {
		logger.Error("couldNotGetAFCClient", err)
		return nil
	}

	cNFC, err := nfc.NewClient(b)
	if err != nil {
		logger.Error("couldNotGetNFCClient", err)
		return nil
	}

	svc := &Service{
		b:      b,
		logger: logger,
		cAFC:   cAFC,
		cNFC:   cNFC,
	}

	b.RegisterService("nfl", svc)

	return svc
}

func (s *Service) ValidateTeam(conference string, team string) bool {

	if conference == "afc" {
		return s.cAFC.ValidateTeam(team)
	}

	return s.cNFC.ValidateTeam(team)
}

func (s *Service) PredictScore(afc string, nfc string) *Score {

	afcScore := s.cAFC.PredictScore(afc)
	nfcScore := s.cNFC.PredictScore(nfc)

	if strings.ToLower(afc) == "chiefs" && nfcScore > afcScore {
		afcScore = nfcScore + 1
	}

	return &Score{
		AFC: afcScore,
		NFC: nfcScore,
	}
}
