package config

import "testing"

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		args    []string
		want    Config
		wantErr bool
	}{
		{
			name: "flags",
			args: []string{"-port", "/dev/ttyUSB0", "-baud", "115200", "-format", "hex"},
			want: Config{Port: "/dev/ttyUSB0", Baud: 115200, Format: FormatHex},
		},
		{
			name: "positional port",
			args: []string{"COM3"},
			want: Config{Port: "COM3", Baud: 460800, Format: FormatHex},
		},
		{
			name:    "missing port",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "bad format",
			args:    []string{"-port", "COM1", "-format", "json"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Fatalf("Parse() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
