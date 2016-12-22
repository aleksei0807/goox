package goox

type GooxComponent struct {
	TagName string
	Text string
	Attributes map[string]string
	Styles map[string]string
	Childs []GooxComponent
}
