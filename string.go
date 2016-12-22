package goox

import (
	"github.com/kirillDanshin/myutils"
)

func (gooxComponent GooxComponent) String() string {
	if gooxComponent.Text != "" && gooxComponent.Tagname == "" {
		return gooxComponent.Text
	}
	if gooxComponent.Tagname == "" {
		return ""
	}
	result := "<"
	result = myutils.Concat(result, gooxComponent.Tagname)
	if len(gooxComponent.Attributes) > 0 {
		for key, v := range gooxComponent.Attributes {
			result = myutils.Concat(result, " ", key,"=\"", v, "\"")
		}
	}
	if len(gooxComponent.Styles) > 0 {
		result = myutils.Concat(result, " style=\"")
		for key, v := range gooxComponent.Styles {
			result = myutils.Concat(result, key, ":", v, ";")
		}
		result = myutils.Concat(result, "\"")
	}

	/* TODO add list single tags */
	if (gooxComponent.Tagname == "meta" || gooxComponent.Tagname == "br") {
		result = myutils.Concat(result, "/>")
		return result
	} else {
		result = myutils.Concat(result, ">")
	}

	if (gooxComponent.Text != "") {
		result = myutils.Concat(result, gooxComponent.Text)
	}

	if len(gooxComponent.Childs) > 0 {
		for _, v := range gooxComponent.Childs {
			result = myutils.Concat(result, v.String())
		}
	}
	result = myutils.Concat(result, "</", gooxComponent.Tagname, ">")
	return result
}
