package rss

import (
	"github.com/JLoup/xml/utils"
	"github.com/JLoup/flag"
)

var (
	LeafElementHasChild      = flag.InitFlag(&utils.ErrorFlagCounter, "LeafElementHasChild")
	MissingAttribute         = flag.InitFlag(&utils.ErrorFlagCounter, "MissingAttribute")
	AttributeDuplicated      = flag.InitFlag(&utils.ErrorFlagCounter, "AttributeDuplicated")
	XHTMLEncodeToStringError = flag.InitFlag(&utils.ErrorFlagCounter, "XHTMLEncodeToStringError")
	CannotFlush              = flag.InitFlag(&utils.ErrorFlagCounter, "CannotFlush")
	DateFormat               = flag.InitFlag(&utils.ErrorFlagCounter, "DateFormat")
	IriNotValid              = flag.InitFlag(&utils.ErrorFlagCounter, "IriNotValid")
)