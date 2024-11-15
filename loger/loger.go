package loger

import (
	"log"
)

var L log.Logger

func init() {
	L.SetFlags(log.Ldate | log.Lmsgprefix | log.Ltime)
}
