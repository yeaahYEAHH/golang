package vlc

import (
	"archivator/lib/compression/vlc/chunk"
	"archivator/lib/compression/vlc/table"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"strings"
)

type EncoderDecoder struct {
	tblGenerator table.Generator
}

func NewEncoderDecoder(tblGenerator table.Generator) EncoderDecoder {
	return EncoderDecoder{tblGenerator: tblGenerator}
}

func (ed EncoderDecoder) Encode(str string) []byte {
	tbl := ed.tblGenerator.NewTable(str)

	// encode to binary: some text -> 10010101
	encoded := encodeBin(str, tbl)

	return buildEncodedFile(tbl, encoded)
}

func buildEncodedFile(tbl table.EncodingTable, data string) []byte {
	encodedTbl := encodeTable(tbl)

	var buf bytes.Buffer

	buf.Write(encodeInt(len(encodedTbl)))
	buf.Write(encodeInt(len(data)))
	buf.Write(encodedTbl)
	buf.Write(chunk.SplitByChunks(data).Bytes())

	return buf.Bytes()
}

func encodeInt(num int) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))

	return res
}

func decodeTable(data []byte) table.EncodingTable {
	var tbl table.EncodingTable

	r := bytes.NewReader(data)
	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		log.Fatal(err)
	}

	return tbl
}

func encodeTable(tbl table.EncodingTable) []byte {
	var tableBuf bytes.Buffer

	if err := gob.NewEncoder(&tableBuf).Encode(tbl); err != nil {
		log.Fatal("can`t serialize table: ", err)
	}

	return tableBuf.Bytes()
}

func (ed EncoderDecoder) Decode(encodingData []byte) string {
	tbl, data := parseFile(encodingData)

	return tbl.Decode(data)
}

func parseFile(data []byte) (table.EncodingTable, string) {
	const (
		tableSizeBytesCount = 4
		dataSizeBytesCount
	)
	tableSizeBinary, data := data[:tableSizeBytesCount], data[tableSizeBytesCount:]
	dataSizeBinary, data := data[:dataSizeBytesCount], data[dataSizeBytesCount:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(dataSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]

	tbl := decodeTable(tblBinary)

	body := chunk.NewBinChunks(data).Join()

	return tbl, body[:dataSize]
}

func encodeBin(str string, table table.EncodingTable) string {
	var buf strings.Builder

	for _, char := range str {
		buf.WriteString(bin(char, table))
	}

	return buf.String()
}

func bin(char rune, table table.EncodingTable) string {
	res, ok := table[char]
	if !ok {
		panic("unknowa character: " + string(char))
	}

	return res
}
