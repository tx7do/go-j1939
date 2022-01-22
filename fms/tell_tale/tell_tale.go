package tell_tale

import "fmt"

type TellTale struct {
	Number uint8
	Status TtsStatusType
}

func NewTellTaleWithValue(number uint8, status TtsStatusType) *TellTale {
	c := &TellTale{}
	c.Number = number
	c.Status = status
	return c
}

func (c *TellTale) GetNumber() uint8 {
	return c.Number
}

func (c *TellTale) SetNumber(number uint8) {
	c.Number = number
}

func (c *TellTale) GetStatus() TtsStatusType {
	return c.Status
}

func (c *TellTale) SetStatus(status TtsStatusType) {
	c.Status = status
}

func (c *TellTale) ToString() string {
	return fmt.Sprintf("TTS %d: %s -> Status: %s(%d)\n",
		c.Number, GetNameForTTSNumber(c.Number),
		GetStatusName(c.Status), c.Status)
}
