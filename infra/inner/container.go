package inner

import (
	"janus-portal/infra/inner/interfaces"
)

type Container struct {
	wxwork interfaces.Wxwork
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) SetWxwork(wxwork interfaces.Wxwork) {
	c.wxwork = wxwork
}

func (c *Container) GetWxwork() interfaces.Wxwork {
	return c.wxwork
}
