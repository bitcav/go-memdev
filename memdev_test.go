//Package memdev
//A Go package that access to Memory Devices information.

package memdev

import (
	"encoding/binary"
	"errors"
	"os"
	"testing"

	"github.com/digitalocean/go-smbios/smbios"
)

//dev builds a Memory Device structure with an n-byte formatted section.
func dev(n int, size uint16, strs ...string) *smbios.Structure {
	f := make([]byte, n)
	binary.LittleEndian.PutUint16(f[8:10], size)
	return &smbios.Structure{Header: smbios.Header{Type: 17, Length: uint8(n + 4)}, Formatted: f, Strings: strs}
}

func TestParseMemory(t *testing.T) {
	// Hyper-V (GitHub Actions Windows runner): only 3 strings, serial/asset/part
	// are index 0. Positional Strings[3] used to panic here.
	hv := dev(0x24, 4096, "M0001", "BANK 0", "Microsoft Corporation")
	hv.Formatted[10], hv.Formatted[14] = 0x02, 0x01 // Unknown, Other
	hv.Formatted[12], hv.Formatted[13], hv.Formatted[19] = 1, 2, 3

	// DDR4 DIMM, strings out of field order to prove lookup is by index.
	ddr4 := dev(0x24, 8192, "SN123", "DIMM A", "Kingston", "KF432C16", "TAG1")
	binary.LittleEndian.PutUint16(ddr4.Formatted[4:6], 64)
	binary.LittleEndian.PutUint16(ddr4.Formatted[6:8], 64)
	binary.LittleEndian.PutUint16(ddr4.Formatted[17:19], 3200)
	ddr4.Formatted[10], ddr4.Formatted[14] = 0x09, 0x1a // DIMM, DDR4
	ddr4.Formatted[12], ddr4.Formatted[19], ddr4.Formatted[20], ddr4.Formatted[21], ddr4.Formatted[22] = 2, 3, 1, 5, 4

	// DDR5 (0x22) and form factor 0x10 used to index past the tables.
	ddr5 := dev(0x24, 0x7fff)
	binary.LittleEndian.PutUint32(ddr5.Formatted[24:28], 65536) // Extended Size, MB
	ddr5.Formatted[10], ddr5.Formatted[14] = 0x10, 0x22

	// Bit 15 set: size is in KB.
	kb := dev(0x24, 0x8000|512)
	kb.Formatted[14] = 0xff // newer than the table

	// SMBIOS 2.1 structure: ends before Speed and Manufacturer.
	old := dev(0x11, 256, "DIMM0")
	old.Formatted[12], old.Formatted[14] = 1, 0x0f

	tests := []struct {
		name string
		s    *smbios.Structure
		want Memory
	}{
		{"hyper-v", hv, Memory{Bank: "M0001", Size: 4096, Unit: "MB", Type: "Other", FormFactor: "Unknown", Manufacturer: "Microsoft Corporation"}},
		{"ddr4", ddr4, Memory{Bank: "DIMM A", Size: 8192, Unit: "MB", Type: "DDR4", FormFactor: "DIMM", Manufacturer: "Kingston",
			Serial: "SN123", AssetTag: "TAG1", PartNumber: "KF432C16", Speed: 3200, DataWidth: 64, TotalWidth: 64}},
		{"ddr5 extended size", ddr5, Memory{Size: 65536, Unit: "MB", Type: "DDR5", FormFactor: "Die"}},
		{"kb unit", kb, Memory{Size: 512, Unit: "KB", Type: "Unknown"}},
		{"unknown size", dev(0x24, 0xffff), Memory{}},
		{"smbios 2.1", old, Memory{Bank: "DIMM0", Size: 256, Unit: "MB", Type: "SDRAM"}},
	}
	for _, tt := range tests {
		got, ok := parseMemory(tt.s)
		if !ok || got != tt.want {
			t.Errorf("%s: got %+v, %v; want %+v", tt.name, got, ok, tt.want)
		}
	}

	if _, ok := parseMemory(dev(0x24, 0)); ok {
		t.Error("empty slot: got ok, want skipped")
	}
}

//TestInfo and TestSlots read the real SMBIOS table, which needs root on Linux.
func TestInfo(t *testing.T) {
	_, err := Info()
	if errors.Is(err, os.ErrPermission) {
		t.Skip("needs elevated privileges:", err)
	}
	if err != nil {
		t.Fatalf("Expected nil error but got %v", err)
	}
}

func TestSlots(t *testing.T) {
	_, err := Slots()
	if errors.Is(err, os.ErrPermission) {
		t.Skip("needs elevated privileges:", err)
	}
	if err != nil {
		t.Fatalf("Expected nil error but got %v", err)
	}
}
