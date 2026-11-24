package version

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantCode int
		wantOut  string
		wantErr  string
	}{
		{name: "version word", args: []string{"version"}, wantCode: 0, wantOut: "nextar-proxy dev"},
		{name: "version flag", args: []string{"--version"}, wantCode: 0, wantOut: "nextar-proxy dev"},
		// A placeholder must exit non-zero: systemd would otherwise show a
		// binary that does nothing as a healthy, running service.
		{name: "no arguments", args: nil, wantCode: 1, wantErr: "not implemented yet"},
		{name: "unknown command", args: []string{"serve"}, wantCode: 1, wantErr: "not implemented yet"},
		{name: "version with extra argument", args: []string{"version", "x"}, wantCode: 1, wantErr: "not implemented yet"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out, errOut bytes.Buffer
			code := Run("nextar-proxy", tt.args, &out, &errOut)
			if code != tt.wantCode {
				t.Fatalf("exit code = %d, want %d", code, tt.wantCode)
			}
			if !strings.Contains(out.String(), tt.wantOut) {
				t.Errorf("stdout = %q, want it to contain %q", out.String(), tt.wantOut)
			}
			if !strings.Contains(errOut.String(), tt.wantErr) {
				t.Errorf("stderr = %q, want it to contain %q", errOut.String(), tt.wantErr)
			}
		})
	}
}
