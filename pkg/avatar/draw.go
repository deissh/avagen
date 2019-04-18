package avatar

import (
	"errors"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
)

var (
	errFontRequired = errors.New("font file is required")
	errInvalidFont  = errors.New("invalid font")
)

// drawer draws an image.Image
type drawer struct {
	fontSize    float64
	dpi         float64
	fontHinting font.Hinting
	face        font.Face
	font        *truetype.Font
}

func newDrawer(fontFile string, fontSize float64) (*drawer, error) {
	if fontFile == "" {
		return nil, errFontRequired
	}
	g := new(drawer)
	g.fontSize = fontSize
	g.dpi = 72.0
	g.fontHinting = font.HintingNone

	parsedfont, err := parseFont(fontFile)
	if err != nil {
		return nil, errInvalidFont
	}
	g.face = truetype.NewFace(parsedfont, &truetype.Options{
		Size:    g.fontSize,
		DPI:     g.dpi,
		Hinting: g.fontHinting,
	})

	g.font = parsedfont
	return g, nil
}

// our avatar image is square
func (g *drawer) Draw(s string, size int, bg *color.RGBA) image.Image {
	// draw the background
	dst := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(dst, dst.Bounds(), &image.Uniform{bg}, image.ZP, draw.Src)

	// draw the text
	drawer := &font.Drawer{
		Dst:  dst,
		Src:  image.White,
		Face: g.face,
	}

	// font index
	fi := g.font.Index([]rune(s)[0])

	// glyph example: http://www.freetype.org/freetype2/docs/tutorial/metrics.png
	var gbuf truetype.GlyphBuf
	fsize := fixed.Int26_6(g.fontSize * g.dpi * (64.0 / 72.0))
	_ = gbuf.Load(g.font, fsize, fi, font.HintingFull)

	// center
	dY := int((size - int(gbuf.Bounds.Max.Y-gbuf.Bounds.Min.Y)>>6) / 2)
	dX := ((size - (len([]rune(s)) * int(gbuf.AdvanceWidth) >> 6)) + ((int(gbuf.AdvanceWidth)>>6)-(int(gbuf.Bounds.Max.X-gbuf.Bounds.Min.X)>>6))*len([]rune(s))) / 2
	//dX := (int(size - len([]rune(s)) * (int(gbuf.AdvanceWidth) >> 6) - int(gbuf.AdvanceWidth) >> 6) - int(gbuf.Bounds.Max.X + gbuf.Bounds.Min.X) >> 6 ) / 2
	y := int(gbuf.Bounds.Max.Y>>6) + dY

	drawer.Dot = fixed.Point26_6{
		X: fixed.I(dX),
		Y: fixed.I(y),
	}
	drawer.DrawString(s)

	return dst
}

// parseFont parse the font file as *truetype.Font (TTF)
func parseFont(fontFile string) (*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return nil, err
	}

	parsedfont, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	return parsedfont, nil
}
