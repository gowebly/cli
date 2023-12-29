package injectors

import (
	"testing"

	"github.com/gowebly/gowebly/internal/attachments"
	"github.com/gowebly/gowebly/internal/config"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	_, err := New(&config.Config{Backend: &config.Backend{}, Frontend: &config.Frontend{}}, &attachments.Attachments{})
	require.NoError(t, err)
}
