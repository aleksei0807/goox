package goox

type GooxComponent struct {
	Tagname string
	Text string
	Attributes map[string]string
	Styles map[string]string
	Childs []GooxComponent
}
