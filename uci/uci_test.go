package uci

import (
	gomock "go.uber.org/mock/gomock"
	"testing"
)

func TestExecuteUCIcommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockEngine := NewMockEngine(ctrl)
	mockEngine.EXPECT().Reset().Return().AnyTimes()
	mockEngine.EXPECT().Stop().Return().AnyTimes()
	mockEngine.EXPECT().Quit().Return().AnyTimes()

	ExecuteUCIcommand(mockEngine, "")
	ExecuteUCIcommand(mockEngine, "uci")
	ExecuteUCIcommand(mockEngine, "isready")
	ExecuteUCIcommand(mockEngine, "position")
	ExecuteUCIcommand(mockEngine, "go")
	ExecuteUCIcommand(mockEngine, "ucinewgame")
	ExecuteUCIcommand(mockEngine, "stop")
	ExecuteUCIcommand(mockEngine, "setoption")
	ExecuteUCIcommand(mockEngine, "quit")
	ExecuteUCIcommand(mockEngine, "unknown")
}
