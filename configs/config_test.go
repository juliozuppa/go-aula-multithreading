package configs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestLoadConfigSuccessfully tests that LoadConfig can successfully load a configuration from a file.
func TestLoadConfigSuccessfully(t *testing.T) {
	cfg, err := LoadConfig("../cmd/desafio_busca_cep")
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	brasilApi := cfg.BrasilApi
	fmt.Printf("%s: %s\n", brasilApi.Name, brasilApi.URL)
	assert.NotNil(t, brasilApi.Name)
	assert.NotNil(t, brasilApi.URL)
	assert.Contains(t, brasilApi.URL, "https://")
	assert.Contains(t, brasilApi.URL, "%s")

	viaCep := cfg.ViaCep
	fmt.Printf("%s: %s\n", viaCep.Name, viaCep.URL)
	assert.NotNil(t, viaCep.Name)
	assert.NotNil(t, viaCep.URL)
	assert.Contains(t, viaCep.URL, "https://")
	assert.Contains(t, viaCep.URL, "%s")

}
