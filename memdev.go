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
type Slots struct {
	Free int `json:"free"`
	Used int `json:"used"`
}

func formFactorType(ff int) string {
	return [...]string{"",
		"Other",
		"Unknown",
		"SIMM",
		"SIP",
		"Chip",
		"DIP",
		"ZIP",
		"Proprietary Card",
		"DIMM",
		"TSOP",
		"Row of Chips",
		"RIMM",
		"SODIMM",
		"SRIMM",
		"FBDIMM"}[ff]
}

func memoryType(mt int) string {
	return [...]string{"",
		"Other",
		"Unknown",
		"DRAM",
		"EDRAM",
		"VRAM",
		"SRAM",
		"RAM",
		"ROM",
		"FLASH",
		"EEPROM",
		"FEPROM",
		"EPROM",
		"CDRAM",
		"3DRAM",
		"SDRAM",
		"SGRAM",
		"RDRAM",
		"DDR",
		"DDR2 FB-DIMM",
		"Reserved",
		"Reserved",
		"Reserved",
		"DDR3",
		"FBD2",
		"DDR4",
		"LPDDR",
		"LPDDR2",
		"LPDDR3",
		"LPDDR4"}[mt]
}

//Info() returns a slice of Memory struct with Memory Modules Information.
func Info() ([]Memory, error) {
	var mems []Memory
	stream, _, err := smbios.Stream()
	if err != nil {
		return []Memory{}, err
	}

	defer stream.Close()

	d := smbios.NewDecoder(stream)
	ss, err := d.Decode()
	if err != nil {
		return []Memory{}, err
	}

	for _, s := range ss {
		if s.Header.Type != 17 {
			continue
		}

		size := int(binary.LittleEndian.Uint16(s.Formatted[8:10]))
		bankLocator := s.Strings[0]

		if size == 0 {
			bankLocator = "empty"
			continue
		}

		if size == 0x7fff {
			size = int(binary.LittleEndian.Uint32(s.Formatted[24:28]))
		}

		manufacturer := s.Strings[1]
		serial := s.Strings[2]
		assetTag := s.Strings[3]
		partNumber := s.Strings[4]
		totalWidth := s.Formatted[4]
		dataWidth := s.Formatted[6]
		formFactorBytes := s.Formatted[10]
		memTypeBytes := s.Formatted[14]
		memType := memoryType(int(memTypeBytes))
		formFactor := formFactorType(int(formFactorBytes))
		memSpeed := int(binary.LittleEndian.Uint16(s.Formatted[17:19]))

		unit := "KB"
		if s.Formatted[9]&0x80 == 0 {
			unit = "MB"
		}

		mems = append(mems, Memory{
			Bank:         bankLocator,
			Size:         size,
			Unit:         unit,
			Type:         memType,
			FormFactor:   formFactor,
			Manufacturer: manufacturer,
			Serial:       serial,
			AssetTag:     assetTag,
			PartNumber:   partNumber,
			Speed:        memSpeed,
			DataWidth:    int(dataWidth),
			TotalWidth:   int(totalWidth),
		})
	}
	return mems, nil
}

//Slots() returns a Slot struct with the amount of Free and Used slots of memory.
func Slots() (Slots, error) {
	var slots Slots
	stream, _, err := smbios.Stream()
	if err != nil {
		return Slots{}, err
	}

	defer stream.Close()

	d := smbios.NewDecoder(stream)
	ss, err := d.Decode()
	if err != nil {
		return Slots{}, err
	}

	for _, s := range ss {
		if s.Header.Type != 17 {
			continue
		}

		size := int(binary.LittleEndian.Uint16(s.Formatted[8:10]))

		if size == 0 {
			slots.Free += 1
			continue
		} else {
			slots.Used += 1
		}
	}
	return slots, nil
}
