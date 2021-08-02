package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/rs/xid"
)

func main() {
	x := uint64(8817236)
	b := make([]byte, 8) // 64bit

	// binary.PutUvarint(b, x)
	binary.BigEndian.PutUint64(b, x)
	// binary.LittleEndian.PutUint64(b, x)
	s := hex.EncodeToString(b)

	fmt.Println("x\t:", x)
	fmt.Println("b\t:", b)
	fmt.Println("s\t:", s)
	fmt.Println("s\t:", fmt.Sprintf("%.16x", x))
	fmt.Println("xid\t:", xid.New().String())

}
