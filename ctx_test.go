package glog

import (
	"testing"
)

func TestCtx(t *testing.T) {
	//logging.toStderr = true

	ctx := NewCtx("Fred")
	ctx.Info("Boo")
	ctx.Infof("Baz %d", 1)

	ctx.V(1).Info("Verbose")
	logging.verbosity = 1
	ctx.V(1).Info("Verbose Now")
}
