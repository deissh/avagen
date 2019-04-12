package avatar

import (
	"image"
	"image/color"
	"reflect"
	"testing"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func Test_newDrawer(t *testing.T) {
	type args struct {
		fontFile string
		fontSize float64
	}
	var tests []struct {
		name    string
		args    args
		want    *drawer
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newDrawer(tt.args.fontFile, tt.args.fontSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("newDrawer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDrawer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_drawer_Draw(t *testing.T) {
	type fields struct {
		fontSize    float64
		dpi         float64
		fontHinting font.Hinting
		face        font.Face
		font        *truetype.Font
	}
	type args struct {
		s    string
		size int
		bg   *color.RGBA
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   image.Image
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &drawer{
				fontSize:    tt.fields.fontSize,
				dpi:         tt.fields.dpi,
				fontHinting: tt.fields.fontHinting,
				face:        tt.fields.face,
				font:        tt.fields.font,
			}
			if got := g.Draw(tt.args.s, tt.args.size, tt.args.bg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("drawer.Draw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFont(t *testing.T) {
	type args struct {
		fontFile string
	}
	var tests []struct {
		name    string
		args    args
		want    *truetype.Font
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFont(tt.args.fontFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFont() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFont() = %v, want %v", got, tt.want)
			}
		})
	}
}
