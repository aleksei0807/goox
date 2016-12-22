package main

import (
	"github.com/aleksei0807/goox"
	"log"
)

func main() {
	myComponent := goox.GooxComponent{
		TagName: "div",
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
				TagName: "span",
				Attributes: map[string]string{
					"class": "my-span",
				},
				Childs: []goox.GooxComponent{
					goox.GooxComponent{
						Text: "просто текст",
					},
				},
			},
			goox.GooxComponent{
				TagName: "span",
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
	}
	log.Printf("%v", myComponent)
}
