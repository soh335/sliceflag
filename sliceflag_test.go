package sliceflag

import (
	"bytes"
	"flag"
	"reflect"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	flagset := flag.NewFlagSet("test", flag.ContinueOnError)

	durationFlag1 := Duration(flagset, "duration1", []time.Duration{}, "duration1 value")
	durationFlag2 := Duration(flagset, "duration2", []time.Duration{time.Second}, "duration2 value")
	durationFlag3 := Duration(flagset, "duration3", []time.Duration{time.Minute}, "duration3 value")

	floatFlag1 := Float64(flagset, "float1", []float64{}, "float1 value")
	floatFlag2 := Float64(flagset, "float2", []float64{10}, "float2 value")
	floatFlag3 := Float64(flagset, "float3", []float64{20}, "float3 value")

	intFlag1 := Int(flagset, "int1", []int{}, "int1 value")
	intFlag2 := Int(flagset, "int2", []int{10}, "int2 value")
	intFlag3 := Int(flagset, "int3", []int{20}, "int3 value")

	int64Flag1 := Int64(flagset, "int64_1", []int64{}, "int64_1 value")
	int64Flag2 := Int64(flagset, "int64_2", []int64{10}, "int64_2 value")
	int64Flag3 := Int64(flagset, "int64_3", []int64{20}, "int64_3 value")

	stringFlag1 := String(flagset, "string1", []string{}, "string1 value")
	stringFlag2 := String(flagset, "string2", []string{"ddd"}, "string2 value")
	stringFlag3 := String(flagset, "string3", []string{"eee"}, "string3 value")

	args := []string{
		"-string1", "aaa",
		"-string1", "bbb",
		"-string2", "ccc",

		"-duration1", "1s",
		"-duration1", "10s",
		"-duration2", "10m",

		"-float1", "10",
		"-float1", "20",
		"-float2", "30",

		"-int1", "10",
		"-int1", "20",
		"-int2", "30",

		"-int64_1", "10",
		"-int64_1", "20",
		"-int64_2", "30",
	}

	if err := flagset.Parse(args); err != nil {
		t.Fatal(err)
	}
	if !flagset.Parsed() {
		t.Error("flagset.Parsed() = false after Parse")
	}

	if e, g := []string{"aaa", "bbb"}, *stringFlag1; !reflect.DeepEqual(e, g) {
		t.Errorf("stringFlag1 expected %v got %v", e, g)
	}
	if e, g := []string{"ccc"}, *stringFlag2; !reflect.DeepEqual(e, g) {
		t.Errorf("stringFlag2 expected %v got %v", e, g)
	}
	if e, g := []string{"eee"}, *stringFlag3; !reflect.DeepEqual(e, g) {
		t.Errorf("stringFlag3 expected %v got %v", e, g)
	}

	if e, g := []time.Duration{time.Second, time.Second * 10}, *durationFlag1; !reflect.DeepEqual(e, g) {
		t.Errorf("durationFlag1 expected %v got %v", e, g)
	}
	if e, g := []time.Duration{time.Minute * 10}, *durationFlag2; !reflect.DeepEqual(e, g) {
		t.Errorf("durationFlag2 expected %v got %v", e, g)
	}
	if e, g := []time.Duration{time.Minute}, *durationFlag3; !reflect.DeepEqual(e, g) {
		t.Errorf("durationFlag3 expected %v got %v", e, g)
	}

	if e, g := []float64{10, 20}, *floatFlag1; !reflect.DeepEqual(e, g) {
		t.Errorf("floatFlag1 expected %v got %v", e, g)
	}
	if e, g := []float64{30}, *floatFlag2; !reflect.DeepEqual(e, g) {
		t.Errorf("floatFlag2 expected %v got %v", e, g)
	}
	if e, g := []float64{20}, *floatFlag3; !reflect.DeepEqual(e, g) {
		t.Errorf("floatFlag3 expected %v got %v", e, g)
	}

	if e, g := []int{10, 20}, *intFlag1; !reflect.DeepEqual(e, g) {
		t.Errorf("intFlag1 expected %v got %v", e, g)
	}
	if e, g := []int{30}, *intFlag2; !reflect.DeepEqual(e, g) {
		t.Errorf("intFlag2 expected %v got %v", e, g)
	}
	if e, g := []int{20}, *intFlag3; !reflect.DeepEqual(e, g) {
		t.Errorf("intFlag3 expected %v got %v", e, g)
	}

	if e, g := []int64{10, 20}, *int64Flag1; !reflect.DeepEqual(e, g) {
		t.Errorf("int64Flag1 expected %v got %v", e, g)
	}
	if e, g := []int64{30}, *int64Flag2; !reflect.DeepEqual(e, g) {
		t.Errorf("int64Flag2 expected %v got %v", e, g)
	}
	if e, g := []int64{20}, *int64Flag3; !reflect.DeepEqual(e, g) {
		t.Errorf("int64Flag3 expected %v got %v", e, g)
	}

	var b bytes.Buffer
	flagset.SetOutput(&b)
	flagset.PrintDefaults()

	if e, g := `  -duration1 value
    	duration1 value (default [])
  -duration2 value
    	duration2 value (default [1s])
  -duration3 value
    	duration3 value (default [1m0s])
  -float1 value
    	float1 value (default [])
  -float2 value
    	float2 value (default [10])
  -float3 value
    	float3 value (default [20])
  -int1 value
    	int1 value (default [])
  -int2 value
    	int2 value (default [10])
  -int3 value
    	int3 value (default [20])
  -int64_1 value
    	int64_1 value (default [])
  -int64_2 value
    	int64_2 value (default [10])
  -int64_3 value
    	int64_3 value (default [20])
  -string1 value
    	string1 value (default [])
  -string2 value
    	string2 value (default [ddd])
  -string3 value
    	string3 value (default [eee])
`, b.String(); e != g {
		t.Errorf("defaults expected %v got %v", e, g)
	}

	if e, g := []time.Duration{time.Second, time.Second * 10}, flagset.Lookup("duration1").Value.(flag.Getter).Get(); !reflect.DeepEqual(e, g) {
		t.Errorf("duration1 lookup expected %v got %v", e, g)
	}
	if e, g := []string{"aaa", "bbb"}, flagset.Lookup("string1").Value.(flag.Getter).Get(); !reflect.DeepEqual(e, g) {
		t.Errorf("string1 lookup expected %v got %v", e, g)
	}
}
