package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	data := encoder("GET", 1, 1, "Hello, World!")
	fmt.Println(data)
	decoder(data)
}

func decoder(data []byte) {
	if len(data) <= 16 {
		fmt.Println("data not body")
		return
	}

	packetLen := binary.BigEndian.Uint32(data[:4])
	headerLen := binary.BigEndian.Uint16(data[4:6])
	version := binary.BigEndian.Uint16(data[6:8])
	operation := binary.BigEndian.Uint32(data[8:12])
	sequence := binary.BigEndian.Uint32(data[12:16])
	body := string(data[16:])
	fmt.Println("packetLen:", packetLen)
	fmt.Println("headerLen:", headerLen)
	fmt.Println("version:", version)
	fmt.Println("operation:", operation)
	fmt.Println("sequence:", sequence)
	fmt.Println("body:", body)
}

var opt = map[string]int{
	"GET":    1,
	"POST":   2,
	"PUT":    3,
	"DELETE": 4,
}

func encoder(operation string, sequence, version int, body string) []byte {
	headerLen := 16
	packetLen := len(body) + headerLen
	ret := make([]byte, packetLen)

	binary.BigEndian.PutUint32(ret[:4], uint32(packetLen))
	binary.BigEndian.PutUint16(ret[4:6], uint16(headerLen))
	binary.BigEndian.PutUint16(ret[6:8], uint16(version))
	operation1 := opt[operation]
	binary.BigEndian.PutUint32(ret[8:12], uint32(operation1))
	binary.BigEndian.PutUint32(ret[12:16], uint32(sequence))

	byteBody := []byte(body)
	copy(ret[16:], byteBody)

	return ret
}
