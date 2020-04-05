package identicon

import (
	"image/color"
	"log"
	"stathat.com/c/consistent"
)

var (
	avatarBgColors = map[string]*color.RGBA{
		"45BDF3": {R: 69, G: 189, B: 243, A: 255},
		"E08F70": {R: 224, G: 143, B: 112, A: 255},
		"4DB6AC": {R: 77, G: 182, B: 172, A: 255},
		"9575CD": {R: 149, G: 117, B: 205, A: 255},
		"B0855E": {R: 176, G: 133, B: 94, A: 255},
		"F06292": {R: 240, G: 98, B: 146, A: 255},
		"A3D36C": {R: 163, G: 211, B: 108, A: 255},
		"7986CB": {R: 121, G: 134, B: 203, A: 255},
		"F1B91D": {R: 241, G: 185, B: 29, A: 255},
	}

	defaultColorKey = "45BDF3"

	// simple hash map
	c = consistent.New()
)

func init() {
	for key := range avatarBgColors {
		c.Add(key)
	}
}

func getColorByName(name string) *color.RGBA {
	key, err := c.Get(name)
	if err != nil {
		log.Print(err)
		key = defaultColorKey
	}
	return avatarBgColors[key]
}
