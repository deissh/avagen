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
	ctx.Logger().Printf("")
	name := string(ctx.QueryArgs().Peek("name"))
	if name == "" {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	enc := string(ctx.QueryArgs().Peek("type"))
	if enc == "" {
		enc = "png"
	}

	size := ctx.QueryArgs().GetUintOrZero("size")
	if size == 0 {
		size = 128
	}

	length := ctx.QueryArgs().GetUintOrZero("length")
	if length == 0 {
		length = 2
	}

	b, err := ah.avatar.DrawToBytes(name, size, length, enc)
	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		_, _ = ctx.WriteString(err.Error())
		return
	}

	ctx.SetContentType("image/" + enc)
	ctx.Response.Header.Set("Cache-Control", "max-age=600")
	ctx.SetBody(b)
	ctx.SetStatusCode(http.StatusOK)
}

func main() {
	h := newAvatarHandler("Cousine-Bold.ttf")

	_ = fasthttp.ListenAndServe(":8080", h.fastHTTPHandler)
}
