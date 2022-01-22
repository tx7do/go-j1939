package spec

type Spec struct {
	number uint32
	offset uint32
	name   string
}

func NewSpec(number uint32, name string, offset uint32) *Spec {
	c := &Spec{
		number: number,
		name:   name,
		offset: offset,
	}
	return c
}

func (c *Spec) SetSpnNumber(number uint32) {
	c.number = number
}
func (c *Spec) GetSpnNumber() uint32 {
	return c.number
}

func (c *Spec) SetOffset(offset uint32) {
	c.offset = offset
}
func (c *Spec) GetOffset() uint32 {
	return c.offset
}

func (c *Spec) SetName(name string) {
	c.name = name
}
func (c *Spec) GetName() string {
	return c.name
}
