package main

import (
	"fmt"
	"time"

	"github.com/davidbanham/human_duration"
	"github.com/dustin/go-humanize"
	"github.com/hako/durafmt"
)

func main() {
	now := time.Now()

	// It will expiry in the next 2 hours, 30 minutes, 10 seconds
	expiryTime := now.Add(2 * time.Hour).Add(30 * time.Minute).Add(10 * time.Second)

	expiryDuration := expiryTime.Sub(now)
	fmt.Println("== Simple Duration Format ==")
	fmt.Println("It will expire in the next", expiryDuration)

	fmt.Println("== Humanize Duration Format ==")
	// https://pkg.go.dev/github.com/dustin/go-humanize
	fmt.Println("It will expire in the next", humanize.Time(expiryTime))
	fmt.Println("It will expire in", humanize.RelTime(expiryTime, now, "earlier", "later"))

	fmt.Println("== Human Duration Format ==")
	// https://github.com/davidbanham/human_duration
	fmt.Println("It will expire in the next", human_duration.String(expiryDuration, "second"))
	fmt.Println("It will expire in the next", human_duration.String(expiryDuration, "minute"))

	fmt.Println("== Durafmt Duration Format ==")
	// https://github.com/hako/durafmt
	dfmt := durafmt.Parse(expiryDuration)
	fmt.Println("It will expire in the next", dfmt)

	englishUnits, err := durafmt.DefaultUnitsCoder.Decode("year:years,week:weeks,day:days,hour:hours,minute:minutes,second:seconds,millisecond:milliseconds,microsecond:microseconds")
	if err != nil {
		panic(err)
	}

	fmt.Println("It will expire in the next", dfmt.Format(englishUnits))

	polishUnit, err := durafmt.DefaultUnitsCoder.Decode("rok:lat,tydzień:tygodnie,dzień:dni,godzina:godziny,minuta:minuty,sekunda:sekundy,milisekunda:milisekundy,mikrosekunda:mikrosekundy")
	if err != nil {
		panic(err)
	}

	fmt.Println("Wygaśnie za", dfmt.Format(polishUnit))

	arabicUnit, err := durafmt.DefaultUnitsCoder.Decode("سنة:سنوات,أسبوع:أسابيع,يوم:أيام,ساعة:ساعات,دقيقة:دقائق,ثانية:ثواني,ميلي ثانية:ميلي ثواني,ميكرو ثانية:ميكرو ثواني")
	if err != nil {
		panic(err)
	}

	fmt.Println("ستنتهي في", dfmt.Format(arabicUnit)) // note: this not print the numbers in Arabic scripture
}
