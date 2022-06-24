package processor

import (
	"github.com/wangyihit/thrift_idl/go/thrift_gen/mt/processor"
)

var NewHtmlProcessor = processor.NewHtmlBlock
var NewHtmlEXtractResult_ = processor.NewHtmlEXtractResult_
var NewItemField = processor.NewItemField

type HtmlProcessor struct {
}

func NewHtmlProcesor() *HtmlProcessor {
	p := &HtmlProcessor{}
	return p
}
