package plutonium_test

import (
	"testing"

	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/ole-larsen/plutonium/internal/plutonium"
	"github.com/ole-larsen/plutonium/internal/plutonium/settings"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewServer(t *testing.T) {
	s := plutonium.NewServer()
	assert.NotNil(t, s, "NewServer should return a non-nil Server instance")
	assert.Nil(t, s.GetSettings(), "NewServer should initialize with nil settings")
	assert.Nil(t, s.GetLogger(), "NewServer should initialize with nil logger")
}

func TestSetSettings(t *testing.T) {
	s := plutonium.NewServer()

	cfg := &settings.Settings{}

	s = s.SetSettings(cfg)

	assert.NotNil(t, s.GetSettings(), "SetSettings should set the settings configuration")
	assert.Equal(t, cfg, s.GetSettings(), "SetSettings should correctly assign the given configuration")
}

func TestSetLogger(t *testing.T) {
	s := plutonium.NewServer()

	logger := log.NewLogger("info", log.DefaultBuildLogger)

	s = s.SetLogger(logger)

	assert.NotNil(t, s.GetLogger(), "SetLogger should set the zap logger")
	assert.Equal(t, logger, s.GetLogger(), "SetLogger should correctly assign the given logger")
}

// TestServerSetStorage tests the SetStorage method.
func TestServerSetStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockDBStorageInterface(ctrl)

	srv := plutonium.NewServer()
	srv.SetStorage(mockStorage)

	assert.Equal(t, mockStorage, srv.GetStorage())
}

// TestServerGetStorage tests the GetStorage method.
func TestServerGetStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockDBStorageInterface(ctrl)
	srv := plutonium.NewServer()
	srv.SetStorage(mockStorage)

	assert.Equal(t, mockStorage, srv.GetStorage())
}
