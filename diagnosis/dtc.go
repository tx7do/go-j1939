package diagnosis

import (
	"errors"
	"fmt"
)

const (
	DtcFrameSize = 4

	cmMask  = 0x80
	ocMask  = 0x7F
	fmiMask = 0x1F
)

// DTC
// 诊断故障代码 Diagnostic Trouble Code (DTC)
// 文档可以查看: https://elearning.vector.com/mod/page/view.php?id=5421
type DTC struct {
	// 可疑参数编号 Suspect Parameter number (SPN) (0 to 524,287)
	// 表示错误的SPN。所有的已定义SPN均可在DTC中使用。
	SPN uint32

	// 故障模式标识符 Failure Mode Identifier (FMI)  (0 to 31)
	// 表示已发生错误的性质和类型，例如超出值范围（过高或过低）、传感器短路、更新速率错误、标定错误。
	FMI uint8

	// 发生计数器 Occurrence Counter (OC) (0 to 127)
	// 该计数器计算每个SPN发生错误的次数并存储该数值（即使错误失效）。
	OC uint8

	// SPN转化方法 SPN Conversion Method (CM)
	// 定义DTC对齐方式。值为“0”表示“Structure of a DTC”图中所示的对齐方式；如果值为“1”，则必须区分是标准曾经定义过的3种方式之中的哪一种。系统必须了解这一点。
	CM uint8
}

// NewDTC 创建一个dtc
func NewDTC(spn uint32, fmi, cm, oc uint8) *DTC {
	c := &DTC{spn, fmi, oc, cm}
	return c
}

// NewDTCAndDecode 创建一个dtc并且解码
func NewDTCAndDecode(buffer []byte) *DTC {
	c := &DTC{0, 0, 0, 0}
	err := c.Decode(buffer)
	if err != nil {
		return nil
	}
	return c
}

// ToString 转换为字符串
func (c *DTC) ToString() string {
	return fmt.Sprintf("Diagnosis Trouble Code -> Spn: %d Failure Mode Identifier: %d Occurrence Count: %d",
		c.SPN, c.FMI, c.OC)
}

// GetFmi 获取FMI值
func (c *DTC) GetFmi() uint8 {
	return c.FMI
}

// SetFmi 设置FMI值
func (c *DTC) SetFmi(fmi uint8) {
	c.FMI = fmi
}

// GetSpn 获取SPN值
func (c *DTC) GetSpn() uint32 {
	return c.SPN
}

// SetSpn 设置SPN值
func (c *DTC) SetSpn(spn uint32) {
	c.SPN = spn
}

// GetOc 获取OC值
func (c *DTC) GetOc() uint8 {
	return c.OC
}

// SetOc 设置OC值
func (c *DTC) SetOc(oc uint8) {
	c.OC = oc
}

// GetCm 获取CM值
func (c *DTC) GetCm() uint8 {
	return c.CM
}

// SetCm 设置CM值
func (c *DTC) SetCm(cm uint8) {
	c.CM = cm
}

// Encode 编码DTC
func (c *DTC) Encode(buffer []byte) error {
	if len(buffer) < DtcFrameSize {
		return errors.New("[DTC::Encode] Length smaller than expected")
	}

	//buffer := []byte{0, 0, 0, 0}

	buffer[0] = byte(c.SPN & 0xFF)
	buffer[1] = byte((c.SPN >> 8) & 0xFF)
	buffer[2] = byte(int(c.SPN>>11) & ^fmiMask)

	buffer[2] |= c.FMI & fmiMask

	buffer[3] = c.OC & ocMask
	buffer[3] |= c.CM & cmMask

	return nil
}

// Decode 解码DTC
func (c *DTC) Decode(buffer []byte) error {
	if len(buffer) < DtcFrameSize {
		return errors.New("[DTC::Decode] Length smaller than expected")
	}

	bit1 := int(buffer[0])
	bit2 := int(buffer[1])
	bit3 := int(buffer[2])
	bit4 := int(buffer[3])

	c.SPN = uint32(bit1)
	c.SPN |= uint32(bit2 << 8)
	c.SPN |= uint32((bit3 & (^fmiMask)) << 11)

	c.FMI = uint8(bit3 & fmiMask)
	c.CM = uint8(bit4 & cmMask)
	c.OC = uint8(bit4 & ocMask)

	return nil
}
