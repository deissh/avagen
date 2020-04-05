package identicon

import (
	"bytes"
	"github.com/deissh/avagen/app"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
)

type Corpus struct {
	fontFile    string
	fontSize    float64
	dpi         float64
	fontHinting font.Hinting
	face        font.Face

	font *truetype.Font
}

func init() {
	corpus := Corpus{
		//todo: load from configuration
		fontFile:    "RobotoMono-Regular.ttf",
		fontSize:    64,
		dpi:         72.0,
		fontHinting: font.HintingFull,
	}

	app.RegisterPlugin(app.Plugin{
		Name:        "identicon",
		Description: "default handler",
		Version:     1,

		Args: []app.Arg{
			{
				Key:      "name",
				Required: true,
			},
			{
				Key:      "type",
				Required: false,
				Default:  "png",
			},
		},

		Preload:  corpus.Preload,
		Generate: corpus.Generate,
	})
}

func (c *Corpus) Preload() error {
	parsed, err := parseFont(c.fontFile)
	if err != nil {
		return errors.New("invalid font")
	}

	c.font = parsed
	c.face = truetype.NewFace(c.font, &truetype.Options{
		Size:    c.fontSize,
		DPI:     c.dpi,
		Hinting: c.fontHinting,
	})

	return nil
}

func (c *Corpus) Generate(args app.ParsedArg) ([]byte, error) {
	name := args["name"]
	imgType := args["type"]

	bg := getColorByName(name)
	initials, err := GetInitials(name, opts{
		limit:      2,
		allCaps:    true,
		allowEmail: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "parseInitials error")
	}

	//todo: setup image size
	dst := image.NewRGBA(image.Rect(0, 0, 128, 128))

	// fill background
	draw.Draw(dst, dst.Bounds(), &image.Uniform{C: bg}, image.ZP, draw.Src)

	// draw text in center
	// since the font is monospaced, you can use any letter to get the size
	bounds, advance, ok := c.face.GlyphBounds('g')
	if !ok {
		return nil, errors.New("load GlyphBounds failed")
	}

	// calculate center
	// work with mono fonts
	// https://www.freetype.org/freetype2/docs/tutorial/metrics.png

	dY := 128/2 + (int(bounds.Max.Y)>>6-int(bounds.Min.Y)>>6)/2
	dX := (128 - (len([]rune(initials)) * int(advance) >> 6)) / 2

	point := fixed.Point26_6{X: fixed.I(dX), Y: fixed.I(dY)}
	drawer := &font.Drawer{
		Dst:  dst,
		Src:  image.White,
		Face: c.face,
		Dot:  point,
	}
	drawer.DrawString(initials)

	// encode result
	var buf bytes.Buffer

	switch imgType {
	case "jpeg":
		err := jpeg.Encode(&buf, dst, nil)
		if err != nil {
			return nil, err
		}
	case "png":
		err := png.Encode(&buf, dst)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("image type not found")
	}

	return buf.Bytes(), nil
}

// parseFont parse the font file as *truetype.Font (TTF)
func parseFont(fontFile string) (*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return nil, err
	}

	return truetype.Parse(fontBytes)
}
