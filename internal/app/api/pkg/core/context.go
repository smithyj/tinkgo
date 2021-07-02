package core

type Context struct {
	Config *Config
}

func NewContext(config *Config) (ctx *Context, err error) {
	ctx = &Context{
		Config: config,
	}
	return
}