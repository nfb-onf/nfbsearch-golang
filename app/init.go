package app

import "github.com/robfig/revel"

func init() {
	revel.Filters = []revel.Filter{
		revel.RouterFilter,
		revel.ParamsFilter,
		HeaderFilter,
		revel.CompressFilter,
		revel.ActionInvoker,
	}
}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:])
}
