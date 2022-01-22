package spn

// Sample 样本
type Sample struct {
	// 时间戳-微秒
	timestampMicro int64
	// 数值
	numeric float64
	// 状态
	status uint8
}

// NewSampleWithNumeric 创建一个样本 - 数值
func NewSampleWithNumeric(timestamp int64, numeric float64) *Sample {
	c := &Sample{}
	c.timestampMicro = timestamp
	c.numeric = numeric
	c.status = 0
	return c
}

// NewSampleWithStatus 创建一个样本 - 状态
func NewSampleWithStatus(timestamp int64, status uint8) *Sample {
	c := &Sample{}
	c.timestampMicro = timestamp
	c.numeric = 0
	c.status = status
	return c
}

// SetTimeStamp 设置时间戳
func (c *Sample) SetTimeStamp(timestamp int64) {
	c.timestampMicro = timestamp
}

// AddTimeStampByMillisecond 增量时间戳
func (c *Sample) AddTimeStampByMillisecond(millis int64) {
	c.timestampMicro += (millis % 1000) * 1000
}

// GetTimeStamp 获取时间戳
func (c *Sample) GetTimeStamp() int64 {
	return c.timestampMicro
}

// GetNumeric 获取数值
func (c *Sample) GetNumeric() float64 {
	return c.numeric
}

// GetStatus 获取状态值
func (c *Sample) GetStatus() uint8 {
	return c.status
}
