package mux

type Context struct {
}

type Handler = func(ctx *Context) error

func (c *Context) Params() map[string]string {
	return map[string]string{"name": "mux"}
}
