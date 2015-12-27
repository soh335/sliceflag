package sliceflag

import (
	"flag"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	flagset := flag.NewFlagSet("test", flag.ContinueOnError)
	stringFlag1 := String(flagset, "string1", []string{}, "string1 value")
	stringFlag2 := String(flagset, "string2", []string{"ddd"}, "string2 value")
	stringFlag3 := String(flagset, "string3", []string{"eee"}, "string3 value")
	args := []string{
		"-string1", "aaa",
		"-string2", "ccc",
		"-string1", "bbb",
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
}
