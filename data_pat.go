package astits

// PATData represents a PAT data
// https://en.wikipedia.org/wiki/Program-specific_information
type PATData struct {
	Programs          []*PATProgram
	TransportStreamID uint16
}

// PATProgram represents a PAT program
type PATProgram struct {
	ProgramMapID  uint16 // The packet identifier that contains the associated PMT
	ProgramNumber uint16 // Relates to the Table ID extension in the associated PMT. A value of 0 is reserved for a NIT packet identifier.
}

// parsePATSection parses a PAT section
func parsePATSection(i []byte, offset *int, offsetSectionsEnd int, tableIDExtension uint16) (d *PATData) {

	//fmt.Println("parsePATSection: ", *offset, "offsetSectionEnd: ", offsetSectionsEnd, "byteLen: ", len(i), "bytes: ", i)
	// Init
	d = &PATData{TransportStreamID: tableIDExtension}

	// Loop until end of section data is reached
	for *offset < offsetSectionsEnd && *offset+4 < len(i) {
		d.Programs = append(d.Programs, &PATProgram{
			ProgramMapID:  uint16(i[*offset+2]&0x1f)<<8 | uint16(i[*offset+3]),
			ProgramNumber: uint16(i[*offset])<<8 | uint16(i[*offset+1]),
		})
		*offset += 4
	}
	return
}
