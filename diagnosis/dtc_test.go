package diagnosis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	spn uint32 = 5052
	fmi uint8  = 12
	cm  uint8  = 0
	oc  uint8  = 3
)

var dataBuffer = []byte{ /*0x55, 0xFF,*/ 0xBC, 0x13, 0x0C, 0x03}

func TestDTCEncode(t *testing.T) {
	var dtc = NewDTC(spn, fmi, cm, oc)

	buff := make([]byte, DtcFrameSize)
	err := dtc.Encode(buff)
	assert.Nil(t, err)
	assert.Equal(t, len(buff), DtcFrameSize)

	for i := 0; i < DtcFrameSize; i++ {
		assert.Equal(t, dataBuffer[i], buff[i])
	}
}

func TestDTCDecode(t *testing.T) {
	dtc := NewDTCAndDecode(dataBuffer)
	assert.NotNil(t, dtc)

	assert.Equal(t, dtc.GetSpn(), spn)
	assert.Equal(t, dtc.GetFmi(), fmi)
	assert.Equal(t, dtc.GetCm(), cm)
	assert.Equal(t, dtc.GetOc(), oc)
}

func TestDTCToString(t *testing.T) {
	var dtc = NewDTC(spn, fmi, cm, oc)

	assert.Equal(t, dtc.GetSpn(), spn)
	assert.Equal(t, dtc.GetFmi(), fmi)
	assert.Equal(t, dtc.GetCm(), cm)
	assert.Equal(t, dtc.GetOc(), oc)

	buff := make([]byte, DtcFrameSize)
	err := dtc.Encode(buff)
	assert.Nil(t, err)
	assert.Equal(t, len(buff), DtcFrameSize)

	fmt.Println(dtc.ToString())
}
