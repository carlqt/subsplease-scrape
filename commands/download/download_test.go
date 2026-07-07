package download

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDownloadCommandParsesResolutionFlag(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want any
	}{
		{
			name: "valid resolution flag",
			args: []string{"-r", "480", "https://example.com"},
			want: 480,
		},
		{
			name: "valid resolution flag",
			args: []string{"-r", "720", "https://example.com"},
			want: 720,
		},
		{
			name: "valid resolution flag",
			args: []string{"-r", "1080", "https://example.com"},
			want: 1080,
		},
		{
			name: "invalid resolution flag",
			args: []string{"-r", "7200", "https://example.com"},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			command := NewDownloadCommand()
			err := command.Parse(tt.args)
			if err == nil {
				assert.Equal(t, tt.want, command.Flags.Resolution)
			} else {
				assert.EqualError(t, err, "Invalid resolution: 7200. Choose between 1080, 720 and 480.")
			}
		})
	}
}
