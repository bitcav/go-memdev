//Package memdev
//A Go package that access to Memory Devices information.

package memdev

import (
	"encoding/binary"

	"github.com/digitalocean/go-smbios/smbios"
)

//Memory Device Structure
type Memory struct {
	Bank         string `json:"bank"`
	Size         int    `json:"size"`
	Unit         string `json:"unit"`
	Type         string `json:"type"`
	FormFactor   string `json:"formFactor"`
	Manufacturer string `json:"manufacturer"`
	Serial       string `json:"serial"`
	AssetTag     string `json:"assetTag"`
	PartNumber   string `json:"partNumber"`
	Speed        int    `json:"speed"`
	DataWidth    int    `json:"dataWidth"`
	TotalWidth   int    `json:"totalWidth"`
}

//Slot availability information
type Slot struct {
	Free int `json:"free"`
	Used int `json:"used"`
}

//Info returns a slice of Memory struct with Memory Modules Information.
func Info() ([]Memory, error) {
	devs, err := memoryDevices()
	if err != nil {
		return nil, err
	}

	var mems []Memory
	for _, s := range devs {
		if m, ok := parseMemory(s); ok {
			mems = append(mems, m)
		}
	}
	return mems, nil
}

//Slots returns a Slot struct with the amount of Free and Used slots of memory.
func Slots() (Slot, error) {
	devs, err := memoryDevices()
	if err != nil {
		return Slot{}, err
	}

	var slots Slot
	for _, s := range devs {
		if binary.LittleEndian.Uint16(s.Formatted[8:10]) == 0 {
			slots.Free++
		} else {
			slots.Used++
		}
	}
	return slots, nil
}

//memoryDevices decodes the SMBIOS table and returns its Memory Device (type 17) structures.
func memoryDevices() ([]*smbios.Structure, error) {
	stream, _, err := smbios.Stream()
	if err != nil {
		return nil, err
	}

	defer stream.Close()

	ss, err := smbios.NewDecoder(stream).Decode()
	if err != nil {
		return nil, err
	}

	var devs []*smbios.Structure
	for _, s := range ss {
		// 15 bytes reach Memory Type, the last field read unconditionally; shorter is malformed.
		if s.Header.Type == 17 && len(s.Formatted) >= 15 {
			devs = append(devs, s)
		}
	}
	return devs, nil
}

//parseMemory decodes a Memory Device structure, ok is false for an empty slot.
//Offsets index Formatted, which excludes the 4-byte header (SMBIOS spec offset - 4).
func parseMemory(s *smbios.Structure) (m Memory, ok bool) {
	f := s.Formatted
	raw := binary.LittleEndian.Uint16(f[8:10])
	if raw == 0 {
		return Memory{}, false
	}

	size, unit := int(raw), "MB"
	switch {
	case raw == 0xffff: // unknown size
		size, unit = 0, ""
	case raw == 0x7fff && len(f) >= 28: // real size is in Extended Size, in MB
		size = int(binary.LittleEndian.Uint32(f[24:28]) & 0x7fffffff)
	case raw&0x8000 != 0:
		size, unit = int(raw&0x7fff), "KB"
	}

	m = Memory{
		Bank:         str(s, 12), // Device Locator
		Size:         size,
		Unit:         unit,
		Type:         name(memoryTypes, f[14]),
		FormFactor:   name(formFactors, f[10]),
		Manufacturer: str(s, 19),
		Serial:       str(s, 20),
		AssetTag:     str(s, 21),
		PartNumber:   str(s, 22),
		DataWidth:    int(binary.LittleEndian.Uint16(f[6:8])),
		TotalWidth:   int(binary.LittleEndian.Uint16(f[4:6])),
	}
	if len(f) >= 19 { // Speed is SMBIOS 2.3+
		m.Speed = int(binary.LittleEndian.Uint16(f[17:19]))
	}
	return m, true
}

//str returns the string referenced by the 1-based index byte at Formatted[off].
//Index 0, or a field missing from an older SMBIOS version, means no string.
func str(s *smbios.Structure, off int) string {
	if off >= len(s.Formatted) {
		return ""
	}
	i := int(s.Formatted[off])
	if i == 0 || i > len(s.Strings) {
		return ""
	}
	return s.Strings[i-1]
}
