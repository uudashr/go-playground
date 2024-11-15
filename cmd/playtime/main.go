package main

import (
	"fmt"
	"time"
)

func main() {
	unixMillis := int64(1731656957172)
	t := time.UnixMilli(unixMillis)
	utcTime := t.UTC()
	localTime := t.Local()

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	tokyoTime := t.In(tokyo)

	dubai, err := time.LoadLocation("Asia/Dubai")
	if err != nil {
		panic(err)
	}

	dubaiTime := t.In(dubai)

	// Singapore timezone offset is +8
	// Use time.FixedZone to create a timezone with a fixed offset
	singapore := time.FixedZone("SGT", 8*60*60)
	singaporeTime := t.In(singapore)

	fmt.Println(t)
	fmt.Println(utcTime)
	fmt.Println(localTime)
	fmt.Println(tokyoTime)
	fmt.Println(dubaiTime)
	fmt.Println(singaporeTime)
}
