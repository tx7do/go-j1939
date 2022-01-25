package can

type Filter struct {
	StdFrame bool
	ExtFrame bool
	Id       uint32
	Mask     uint32
}

func NewCanFilter() Filter {
	c := Filter{}
	c.Id = 0
	c.Mask = 0
	c.StdFrame = false
	c.ExtFrame = false
	return c
}

func NewCanFilterWithData(id, mask uint32, filterExt, filterStd bool) Filter {
	c := Filter{}
	c.Id = id
	c.Mask = mask
	c.StdFrame = filterStd
	c.ExtFrame = filterExt
	return c
}

func (c *Filter) GetId() uint32 {
	return c.Id
}
func (c *Filter) SetId(id uint32) {
	c.Id = id
}

func (c *Filter) GetMask() uint32 {
	return c.Mask
}
func (c *Filter) SetMask(mask uint32) {
	c.Mask = mask
}

func (c *Filter) FilterExtFrame() bool {
	return c.ExtFrame
}
func (c *Filter) SetExtFrame(extFrame bool) {
	c.ExtFrame = extFrame
}

func (c *Filter) FilterStdFrame() bool {
	return c.StdFrame
}
func (c *Filter) SetStdFrame(stdFrame bool) {
	c.StdFrame = stdFrame
}
