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

	stringFlag1 := String(flagset, "string1", []string{}, "string1 value")
	stringFlag2 := String(flagset, "string2", []string{"ddd"}, "string2 value")
	stringFlag3 := String(flagset, "string3", []string{"eee"}, "string3 value")

	durationFlag1 := Duration(flagset, "duration1", []time.Duration{}, "duration1 value")
	durationFlag2 := Duration(flagset, "duration2", []time.Duration{time.Second}, "duration2 value")
	durationFlag3 := Duration(flagset, "duration3", []time.Duration{time.Minute}, "duration3 value")
	args := []string{
		"-string1", "aaa",
		"-string1", "bbb",
		"-string2", "ccc",
		"-duration1", "1s",
		"-duration1", "10s",
		"-duration2", "10m",
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

	var b bytes.Buffer
	flagset.SetOutput(&b)
	flagset.PrintDefaults()

	if e, g := `  -duration1 value
    	duration1 value (default [])
  -duration2 value
    	duration2 value (default [1s])
  -duration3 value
    	duration3 value (default [1m0s])
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
