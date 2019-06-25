package slot

import (
	"dmidecode/smbios"
	"fmt"
)

type SystemSlot struct {
	smbios.Header
	Designation          string
	Type                 Type
	DataBusWidth         DataBusWidth
	CurrentUsage         Usage
	Length               Length
	ID                   ID
	Characteristics1     Characteristics1
	Characteristics2     Characteristics2
	SegmentGroupNumber   SegmengGroupNumber
	BusNumber            Number
	DeviceFunctionNumber Number
}

func (s SystemSlot) String() string {
	return fmt.Sprintf("System Slot %s\n"+
		"\tSlot Designation: %s\n"+
		"\tSlot Type: %s\n"+
		"\tSlot Data Bus Width: %s\n"+
		"\tCurrent Usage: %s\n"+
		"\tSlot Length: %s\n"+
		"\tSlot ID: %s\n"+
		"\tSlot Characteristics1: %s\n"+
		"\tSlot Characteristics2: %s\n"+
		"\tSegment Group Number: %s\n"+
		"\tBus Number: %s\n"+
		"\tDevice/Function Number: %s",
		s.Designation,
		s.Type,
		s.DataBusWidth,
		s.CurrentUsage,
		s.Length,
		s.ID,
		s.Characteristics1,
		s.Characteristics2,
		s.SegmentGroupNumber,
		s.BusNumber,
		s.DeviceFunctionNumber)
}
