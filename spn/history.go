package spn

import (
	"go-j1939/spn/spec"
	"sort"
)

const (
	MaxNumberOfSamples = 50000
)

type SpecificSpecSet struct {
	status  *spec.StatusSpec
	numeric *spec.NumericSpec
}

type SampleSet []*Sample

type History struct {
	generalSpec  *spec.Spec
	specificSpec SpecificSpecSet

	samples SampleSet
}

func NewHistory() {

}

func (c *History) AddNumericSample(timestamp int64, spn *NumericSPN) {
	length := len(c.samples)
	if length == MaxNumberOfSamples {
		return
	}

	if c.generalSpec == nil {
		c.generalSpec = spn.GetSpec()
	}

	c.specificSpec.numeric = spn.GetNumericSpec()

	if c.generalSpec != spn.GetSpec() {
		//Not the same spnImpl
		return
	}

	if length != 0 && c.samples[length-1].GetNumeric() == spn.GetFormattedValue() {
		//Value did not change
		return
	}

	sample := NewSampleWithNumeric(timestamp, spn.GetFormattedValue())

	c.samples = append(c.samples, sample)
}

func (c *History) AddStatusSample(timestamp int64, spn *StatusSPN) {
	length := len(c.samples)
	if length == MaxNumberOfSamples {
		return
	}

	if c.generalSpec == nil {
		c.generalSpec = spn.GetSpec()
	}

	c.specificSpec.status = spn.GetStatusSpec()

	if c.generalSpec != spn.GetSpec() {
		//Not the same spnImpl
		return
	}

	if length != 0 && c.samples[length-1].GetStatus() == spn.GetStatusValue() {
		//Value did not change
		return
	}

	sample := NewSampleWithStatus(timestamp, spn.GetStatusValue())

	c.samples = append(c.samples, sample)
}

func (c *History) GetWindow(timestamp int64, milliseconds uint32, samples uint32) SampleSet {
	var window SampleSet
	length := len(c.samples)
	if length < 2 {
		return window
	}

	period := int64(milliseconds / samples)

	current := timestamp - (int64(milliseconds)%1000)*1000

	//Get the beginning
	pos := sort.Search(length,
		func(i int) bool {
			return c.samples[i].GetTimeStamp() >= 0
		},
	)

	if pos != 0 {
		//Start with the sample whose timestamp is lower or equals to the beginning of the window
		pos--
	}

	for current <= timestamp {
		mySample := c.samples[pos]
		window = append(window, mySample)

		current = current + period

		for pos+1 < length && current >= c.samples[pos+1].GetTimeStamp() {
			//TODO check variation of more of 10% and maximums and minimums.
			pos++
		}
	}

	return window
}

func (c *History) GetGeneralSpec() *spec.Spec {
	return c.generalSpec
}

func (c *History) GetNumericSpec() *spec.NumericSpec {
	return c.specificSpec.numeric
}
func (c *History) GetStatusSpec() *spec.StatusSpec {
	return c.specificSpec.status
}
