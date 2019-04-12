package avatar

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/dchest/lru"
)

func TestNew(t *testing.T) {
	type args struct {
		fontFile string
	}
	var tests []struct {
		name string
		args args
		want *InitialsAvatar
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.fontFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithConfig(t *testing.T) {
	type args struct {
		cfg Config
	}
	var tests []struct {
		name string
		args args
		want *InitialsAvatar
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWithConfig(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitialsAvatar_DrawToBytes(t *testing.T) {
	type fields struct {
		drawer *drawer
		cache  *lru.Cache
	}
	type args struct {
		name     string
		size     int
		count    int
		encoding []string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &InitialsAvatar{
				drawer: tt.fields.drawer,
				cache:  tt.fields.cache,
			}
			got, err := a.DrawToBytes(tt.args.name, tt.args.size, tt.args.count, tt.args.encoding...)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitialsAvatar.DrawToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitialsAvatar.DrawToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHan(t *testing.T) {
	type args struct {
		r rune
	}
	var tests []struct {
		name string
		args args
		want bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHan(tt.args.r); got != tt.want {
				t.Errorf("isHan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getColorByName(t *testing.T) {
	type args struct {
		name string
	}
	var tests []struct {
		name string
		args args
		want *color.RGBA
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getColorByName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getColorByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getInitials(t *testing.T) {
	type args struct {
		name  string
		count int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one symbol",
			args: args{
				name:  "deissh",
				count: 1,
			},
			want: "d",
		},
		{
			name: "one unicode symbol",
			args: args{
				name:  "что-то",
				count: 1,
			},
			want: "ч",
		},
		{
			name: "two symbols from sentence",
			args: args{
				name:  "some sentence",
				count: 2,
			},
			want: "ss",
		},
		{
			name: "two symbols from unicode sentence",
			args: args{
				name:  "какое то",
				count: 2,
			},
			want: "кт",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInitials(tt.args.name, tt.args.count); got != tt.want {
				t.Errorf("getInitials() = %v, want %v", got, tt.want)
			}
		})
	}
}
