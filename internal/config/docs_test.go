package config

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/services/chainlink/cfgtest"
)

func TestDoc(t *testing.T) {
	var c chainlink.Config
	d := toml.NewDecoder(strings.NewReader(docsTOML))
	d.DisallowUnknownFields() // Ensure no extra fields
	err := d.Decode(&c)
	if err != nil && strings.Contains(err.Error(), "undecoded keys: ") {
		t.Errorf("Docs contain extra fields: %v", err)
	} else {
		require.NoError(t, err)
	}

	cfgtest.AssertFieldsNotNil(t, c)

	//TODO validate defaults?
}

var (
	//go:embed testdata/example.toml
	exampleTOML string
	//go:embed testdata/example.md
	exampleMarkdown string
)

func Test_generateDocs(t *testing.T) {
	got, err := generateDocs(exampleTOML)
	require.NoError(t, err)
	assert.Equal(t, exampleMarkdown, got)
}
