//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"syscall/js"

	"github.com/alexdogonin/mediastorage_frontend_go-wasm/pkg/router"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func main() {
	log.Println(js.Global().Get(`location`).Get(`href`))
	vecty.SetTitle("go-wasm")

	href := js.Global().Get(`location`).Get(`href`).String()

	r := router.NewRouter()

	r.Handle("/*", func() {
		vecty.RenderBody(&MediaView{})
	})

	r.Handle("/albums", func() {
		vecty.RenderBody(&AlbumsView{})
	})

	err := r.Serve(href)
	if err != nil {
		log.Fatal(err)
	}
}

// PageView is our main page component.
type MediaView struct {
	vecty.Core
	// Input string
}

func (p *MediaView) Render() vecty.ComponentOrHTML {
	resp, err := http.Get("http://localhost:3000/v2/media")
	if err != nil {
		log.Println(err)
		return elem.Body()
	}
	defer resp.Body.Close()

	var mediaResp MediaListResponse
	err = json.NewDecoder(resp.Body).Decode(&mediaResp)
	if err != nil {
		log.Println(err)
		return elem.Body()
	}

	args := make([]vecty.MarkupOrChild, 0, len(mediaResp.Media))
	for _, i := range mediaResp.Media {

		args = append(args, elem.Image(
			vecty.Markup(
				vecty.Property("src", i.Thumb.URL),
				vecty.Style("height", "200px"),
				vecty.Style("margin", "5px"),
			),
		))
	}

	args = append(args, vecty.Markup(
		vecty.Style("background", "rgb(232, 224, 224)"),
	))

	return elem.Body(
		elem.Div(
			args...,
		),
	)
}

type MediaListResponse struct {
	Media  []MediaItem `json:"media"`
	Cursor string      `json:"cursor"`
}

type MediaItem struct {
	UUID     string         `json:"uuid"`
	Thumb    *MediaItemInfo `json:"thumb,omitempty"`
	Detail   *MediaItemInfo `json:"detail,omitempty"`
	Original *MediaItemInfo `json:"original,omitempty"`
}

type MediaItemInfo struct {
	URL    string `json:"url"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
}

type AlbumsView struct {
	vecty.Core
}

func (v *AlbumsView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Text("Albums are going to be here..."),
		),
	)
}
