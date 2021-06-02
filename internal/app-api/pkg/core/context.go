package core

import "context"

type Context struct {
	context.Context
	Config *Config
}

func NewContext(config *Config) (ctx *Context, err error) {
	ctx = &Context{
		Config: config,
	}
	return
}