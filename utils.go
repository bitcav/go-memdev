//Package memdev
//A Go package that access to Memory Devices information.

package memdev

//formFactors is indexed by the SMBIOS Memory Device Form Factor value.
var formFactors = []string{"",
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
	"FBDIMM",
	"Die"}

//memoryTypes is indexed by the SMBIOS Memory Device Type value.
var memoryTypes = []string{"",
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
	"DDR2",
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
	"LPDDR4",
	"Logical non-volatile device",
	"HBM",
	"HBM2",
	"DDR5",
	"LPDDR5",
	"HBM3"}

//name returns names[v], or "Unknown" for values newer than the table.
func name(names []string, v byte) string {
	if int(v) < len(names) {
		return names[v]
	}
	return "Unknown"
}
