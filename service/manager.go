package nfl

import (
	bridge "github.com/IAvellanedaGit/nuevo-platform-mobile-bridge"
	model "github.com/IAvellanedaGit/sandbox-avellaneda-pm-nfl/service/model"
	flatbuffers "github.com/google/flatbuffers/go"
)

type Manager struct {
	b      *bridge.Bridge
	logger bridge.Logger
	bus    *bridge.MessageBus
	svc    *Service
}

const NFL_VERSION = "0.0.1"

func NewManager(b *bridge.Bridge) *Manager {
	m := &Manager{
		b:      b,
		logger: b,
		bus:    b.MessageBusInstance(),
	}

	m.bus.Subscribe("nfl:initialize", m.initialize)

	return m
}

func (mgr *Manager) initialize(m *bridge.Message) ([]byte, error) {
	request := model.GetRootAsInitializeRequest(m.Data, 0)

	mgr.logger.Debug("Initializing NFL Manager [%s] LogLevel [%d]",
		NFL_VERSION,
		request.LogLevel())

	mgr.logger = bridge.NewLogger("nfl", bridge.LogLevel(request.LogLevel()), mgr.logger)

	mgr.svc = NewService(mgr.b, mgr.logger)

	mgr.bus.Subscribe("nfl:validate-team", mgr.validateTeam)
	mgr.bus.Subscribe("nfl:predict-score", mgr.predictScore)

	builder := flatbuffers.NewBuilder(0)

	vOffset := builder.CreateString(NFL_VERSION)

	model.StringResponseStart(builder)
	model.StringResponseAddOutput(builder, vOffset)
	builder.Finish(model.StringResponseEnd(builder))

	return builder.FinishedBytes(), nil
}

func (mgr *Manager) validateTeam(m *bridge.Message) ([]byte, error) {
	request := model.GetRootAsTeamRequest(m.Data, 0)

	valid := mgr.svc.ValidateTeam(string(request.Conference()), string(request.Team()))

	builder := flatbuffers.NewBuilder(0)
	model.BoolResponseStart(builder)
	model.BoolResponseAddOutput(builder, valid)
	builder.Finish(model.BoolResponseEnd(builder))

	return builder.FinishedBytes(), nil
}

func (mgr *Manager) predictScore(m *bridge.Message) ([]byte, error) {

	request := model.GetRootAsPredictScoreRequest(m.Data, 0)

	score := mgr.svc.PredictScore(string(request.Afc()), string(request.Nfc()))

	builder := flatbuffers.NewBuilder(0)

	model.PredictScoreResponseStart(builder)
	model.PredictScoreResponseAddAfc(builder, int32(score.AFC))
	model.PredictScoreResponseAddNfc(builder, int32(score.NFC))
	builder.Finish(model.PredictScoreResponseEnd(builder))

	return builder.FinishedBytes(), nil
}
