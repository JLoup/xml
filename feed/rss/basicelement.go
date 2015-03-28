package rss

import (
	"github.com/JLoup/errors"
	"github.com/JLoup/xml/feed/extension"
	"github.com/JLoup/xml/utils"
	"encoding/xml"
	"fmt"
	"strings"
)

type BasicElement struct {
	Content utils.Element
	name    xml.Name

	Extension extension.VisitorExtension
	depth     utils.DepthWatcher
	Parent    utils.Visitor
}

func NewBasicElement() *BasicElement {
	d := utils.NewDepthWatcher()
	d.SetMaxDepth(1)

	return &BasicElement{depth: d, Content: utils.NewElement("", "", utils.Nop)}
}

func NewBasicElementExt(manager extension.Manager) *BasicElement {
	b := NewBasicElement()

	b.Extension = extension.InitExtension("basicelement", manager)

	return b
}

func (b *BasicElement) SetParent(parent utils.Visitor) {
	b.Parent = parent
}

func (b *BasicElement) Name() xml.Name {
	return b.name
}

func (b *BasicElement) ProcessStartElement(el utils.StartElement) (utils.Visitor, utils.ParserError) {
	if b.depth.IsRoot() {
		b.name = el.Name
		b.Extension = extension.InitExtension(b.name.Local, b.Extension.Manager)

		for _, attr := range el.Attr {
			b.Extension.ProcessAttr(attr, b)
		}
	}

	if b.depth.Down() == utils.MaxDepthReached {
		return b, utils.NewError(LeafElementHasChild, fmt.Sprintf("'%s' shoud not have childs", b.Name))
	}

	return b, nil
}

func (b *BasicElement) ProcessEndElement(el xml.EndElement) (utils.Visitor, utils.ParserError) {
	if b.depth.Up() == utils.RootLevel {
		return b.Parent, b.Validate()
	}

	return b, nil
}

func (b *BasicElement) ProcessCharData(el xml.CharData) (utils.Visitor, utils.ParserError) {
	b.Content.Value = strings.TrimSpace(string(el))
	return b, nil
}

func (b *BasicElement) Validate() utils.ParserError {
	error := errors.NewErrorAggregator()

	b.Extension.Validate(&error)

	if err := b.Content.Validate(); err != nil {
		error.NewError(utils.NewError(err.Flag(), fmt.Sprintf("%s's %s", b.name.Local, err.Msg())))
	}

	return error.ErrorObject()
}

func (b *BasicElement) String() string {
	return b.Content.Value
}

func (b *BasicElement) Reset() {
	b.depth.Reset()
}