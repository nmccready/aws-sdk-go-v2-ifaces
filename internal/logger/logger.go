package logger

import (
	"github.com/nmccready/go-debug"
)

var RootDebug = debug.Debug("aws-sdk-go-v2-ifaces")
var Spawn = RootDebug.Spawn
