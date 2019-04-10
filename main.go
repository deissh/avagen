package main

import (
	"github.com/deissh/avagen/pkg/avatar"
	"github.com/valyala/fasthttp"
	"net/http"
)

type avatarHandler struct {
	avatar *avatar.InitialsAvatar
}

func newAvatarHandler(fontFile string) *avatarHandler {
	return &avatarHandler{
		avatar.NewWithConfig(avatar.Config{
			FontFile: fontFile,
			FontSize: 50,
		}),
	}
}

func (ah avatarHandler) fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	name := ctx.QueryArgs().Peek("name")
	if name == nil || string(name) == "" {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	b, _ := ah.avatar.DrawToBytes(string(name), 128, "png")

	ctx.SetContentType("image/png")
	ctx.Response.Header.Set("Cache-Control", "max-age=600")
	ctx.SetBody(b)
	ctx.SetStatusCode(http.StatusOK)
}

func main() {
	h := newAvatarHandler("wqy-zenhei.ttf")

	_ = fasthttp.ListenAndServe(":8080", h.fastHTTPHandler)
}
