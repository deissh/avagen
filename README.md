# Avagen

```bash
$ avagen g "Fist Second" > image.png
$ avagen g "Fist Second" --type jpeg > image.jpeg
$ avagen g "Fist Second" --plugin identicon> image.png

# simple avatar server
$ avagen serve 
$ avagen serve --addr 0.0.0.0:8080
# request
$ curl --output avatar.png http://0.0.0.0:8080?name=Avagen
$ curl --output avatar.png http://0.0.0.0:8080?name=Best&size=512&fsize=200
$ curl --output avatar.png http://0.0.0.0:8080?name=Example%20With%20Type&type=jpeg
```

```go
package main

import (
	"github.com/deissh/avagen/app"
	// load plugins
	_ "github.com/deissh/avagen/plugins/identicon"
	"log"
	"os"
)

func main()  {
	plugin, _ := app.GetPlugin("identicon")

	bytes, err := plugin.Generate(app.ParsedArg{"name": "Я R", "type": "png"})
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("hello-go.png")
	f.Write(bytes)
	f.Close()
}

```
