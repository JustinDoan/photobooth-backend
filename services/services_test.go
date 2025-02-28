package services_test

import (
	"testing"

	"github.com/justindoan/photobooth-backend/services"
	"github.com/stretchr/testify/assert"
)

func TestNewServices(t *testing.T) {
	services := services.NewServices()
	assert.NotNil(t, services.Forge, "Expected Forge to be initialized")
}
