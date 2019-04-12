package avatar

import (
	"io"
	"testing"
)

func Test_parseInitials(t *testing.T) {
	type args struct {
		src io.Reader
		o   opts
	}
	var tests []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInitials(tt.args.src, tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInitials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseInitials() = %v, want %v", got, tt.want)
			}
		})
	}
}
