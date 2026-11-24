//go:build unix

package files

import (
	"path/filepath"
	"testing"

	"github.com/xyproto/env/v2"
)

func TestBinDirectory(t *testing.T) {
	homeDir := env.HomeDir()
	tests := []struct {
		dir  string
		want bool
	}{
		{"/bin", true},
		{"/sbin", true},
		{"/usr/bin", true},
		{"/usr/sbin", true},
		{"/usr/local/bin", true},
		{"/usr/local/sbin", true},
		{filepath.Join(homeDir, ".bin"), true},
		{filepath.Join(homeDir, "bin"), true},
		{filepath.Join(homeDir, ".local", "bin"), true},
		{filepath.Join(homeDir, "local", "bin"), false},
		{filepath.Join(homeDir, ".local", "binary"), false},
		{filepath.Join(homeDir, ".local", "bin", "subdir"), false},
		{"/usr/local/bin-extra", false},
	}
	for _, tt := range tests {
		t.Run(tt.dir, func(t *testing.T) {
			filename := filepath.Join(tt.dir, "tool")
			if got := BinDirectory(filename); got != tt.want {
				t.Errorf("BinDirectory(%q) = %v, want %v", filename, got, tt.want)
			}
		})
	}
}
