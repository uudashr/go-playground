package timemarshal_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type Tick struct {
	Time time.Time `json:"time"`
}

func TestTime_JSON(t *testing.T) {
	time1 := time.Now().Truncate(time.Second)
	encodedTime, err := time1.MarshalText()
	require.NoError(t, err)

	b := `
	{
		"time": "` + string(encodedTime) + `"
	}
	`

	var tick Tick
	err = json.Unmarshal([]byte(b), &tick)
	require.NoError(t, err)

	require.Equal(t, time1, tick.Time)
}

func TestTime_Marshaling(t *testing.T) {
	time1 := time.Now().Truncate(time.Second)
	encodedTime1, err := time1.MarshalText()
	require.NoError(t, err)
	fmt.Printf("Time1: %v\n", time1)
	fmt.Printf("Time1 location: %v\n", time1.Location())
	fmt.Printf("EncodedTime1: %s\n", string(encodedTime1))

	jsonEncodedTime1, err := json.Marshal(time1)
	require.NoError(t, err)
	fmt.Printf("JSON Encoded Time1: %s\n", string(jsonEncodedTime1))

	var time2 time.Time
	err = time2.UnmarshalText(encodedTime1)
	require.NoError(t, err)
	fmt.Printf("Time2: %v\n", time2)
	fmt.Printf("Time2 location: %v\n", time2.Location())

	require.Equal(t, time1, time2)
}
