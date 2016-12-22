package main

import (
	"github.com/aleksei0807/goox"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	"log"
	"fmt"
)

func main() {
	myComponent := goox.GooxComponent{
		Tagname: "html",
		Childs: []goox.GooxComponent{
			goox.GooxComponent{
				Tagname: "head",
				Childs: []goox.GooxComponent{
					goox.GooxComponent{
						Tagname: "meta",
						Attributes: map[string]string{
							"charset": "utf-8",
						},
					},
				},
			},
			goox.GooxComponent{
				Tagname: "body",
				Childs: []goox.GooxComponent{
					goox.GooxComponent{
						Tagname: "div",
						Attributes: map[string]string{
							"class": "root",
							"id": "root",
							"data-my": "my",
						},
						Styles: map[string]string{
							"background-color": "#fafafa",
							"color": "#3333bf",
							"font-size": "15px",
							"font-weight": "300",
							"line-height": "1",
						},
						Childs: []goox.GooxComponent{
							goox.GooxComponent{
								Tagname: "span",
								Attributes: map[string]string{
									"class": "my-span",
								},
								Text: "текст в спане",
							},
							goox.GooxComponent{
								Tagname: "br",
							},
							goox.GooxComponent{
								Tagname: "span",
								Attributes: map[string]string{
									"class": "my-span2",
								},
								Childs: []goox.GooxComponent{
									goox.GooxComponent{
										Text: "просто текст 2, ",
									},
									goox.GooxComponent{
										Text: "просто текст 3",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	log.Printf("%v", myComponent)

	r := fasthttprouter.New()
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("text/html")
		fmt.Fprintf(ctx, "<!DOCTYPE html>%v", myComponent)
	})
	log.Printf("serve on %s", ":9999")
	fasthttp.ListenAndServe(":9999", r.Handler)
}
