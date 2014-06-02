package glog

import (
	"fmt"
	"strings"
)

type Ctx struct {
	Context       string
	printPrefix   []interface{}
	printlnPrefix []interface{}
	printfPrefix  string
}

func NewCtx(context string) *Ctx {
	if strings.Contains(context, "%") {
		panic("Context must not contain '%'")
	}

	s := fmt.Sprintf("[%s] ", context)
	return &Ctx{
		Context: context, 
		printPrefix: []interface{}{s}, 
		printlnPrefix: []interface{}{fmt.Sprintf("[%s]", context)}, 
		printfPrefix: s}
}

func NewCtxf(format string, args ...interface{}) *Ctx {
	return NewCtx(fmt.Sprintf(format, args...))
}

func (c *Ctx) Info(args ...interface{}) {
	if c != nil {
		logging.print(infoLog, append(c.printPrefix, args...)...)
	}
}
func (c *Ctx) Infoln(args ...interface{}) {
	if c != nil {
		logging.println(infoLog, append(c.printlnPrefix, args...)...)
	}
}

func (c *Ctx) Infof(format string, args ...interface{}) {
	if c != nil {
		logging.printf(infoLog, c.printfPrefix+format, args...)
	}
}

func (c *Ctx) Warning(args ...interface{}) {
	logging.print(warningLog, append(c.printPrefix, args...)...)
}

func (c *Ctx) Warningln(args ...interface{}) {
	logging.println(warningLog, append(c.printlnPrefix, args...)...)
}

func (c *Ctx) Warningf(format string, args ...interface{}) {
	logging.printf(warningLog, c.printfPrefix+format, args...)
}

func (c *Ctx) Error(args ...interface{}) {
	logging.print(errorLog, append(c.printPrefix, args...)...)
}

func (c *Ctx) Errorln(args ...interface{}) {
	logging.println(errorLog, append(c.printlnPrefix, args...)...)
}

func (c *Ctx) Errorf(format string, args ...interface{}) {
	logging.printf(errorLog, c.printfPrefix+format, args...)
}

func (c *Ctx) Fatal(args ...interface{}) {
	logging.print(fatalLog, append(c.printPrefix, args...)...)
}

func (c *Ctx) Fatalln(args ...interface{}) {
	logging.print(fatalLog, append(c.printlnPrefix, args...)...)
}

func (c *Ctx) Fatalf(format string, args ...interface{}) {
	logging.printf(fatalLog, c.printfPrefix+format, args...)
}

func (c *Ctx) V(level Level) *Ctx {
	if V(level) {
		return c
	}
	return nil
}

func (c *Ctx) VB(level Level) bool {
	return c.V(level) != nil
}
