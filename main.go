package main

import (
	"flag"
	"github.com/deissh/avagen/pkg/avatar"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type avatarHandler struct {
	avatar *avatar.InitialsAvatar
}

var (
	serverAddress string
	fontPath string
)
func init() {
	flag.StringVar(&serverAddress, "address", ":8080", "example value :8080")
	flag.StringVar(&fontPath, "font", "Cousine-Bold.ttf", "font file path")
	flag.Parse()
}

func newAvatarHandler(fontFile string) *avatarHandler {
	return &avatarHandler{
		avatar.NewWithConfig(avatar.Config{
			FontFile: fontFile,
			FontSize: 50,
		}),
	}
}

func (ah avatarHandler) handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.String())

	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	enc := r.URL.Query().Get("type")
	if enc == "" {
		enc = "png"
	}

	//size := r.URL.Query().Get("size")
	//if size <= 0 || size >= 1024 {
	//	size = 128
	//}
	//
	//length := r.URL.Query().Get("length")
	//if length == 0 {
	//	length = 2
	//}

	b, err := ah.avatar.DrawToBytes(name, 128, 2, enc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "image/" + enc)
	w.Header().Set("Cache-Control", "max-age=600")
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	if _, err := w.Write(b); err != nil {
		log.Println("unable to write image.")
	}
}

func main() {
	h := newAvatarHandler(fontPath)
	server := fasthttp.Server{}

	listenErr := make(chan error, 1)

	r := http.NewServeMux()
	r.HandleFunc("/", h.handler)

	server.Handler = fasthttpadaptor.NewFastHTTPHandler(r)

	go func() {
		log.Printf("Server starting on %s", serverAddress)
		log.Printf("Press Ctrl+C to stop")
		listenErr <- server.ListenAndServe(serverAddress)
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)


	for {
		select {
		case err := <-listenErr:
			if err != nil {
				log.Fatalf("listener error: %s", err)
			}
			os.Exit(0)

		case <-osSignals:
			log.Print("Shutdown signal received")

			// Attempt the graceful shutdown by closing the listener
			// and completing all inflight requests.
			if err := server.Shutdown(); err != nil {
				log.Fatalf("error with graceful close: %s", err)
			}
		}
	}
}
