package sliceflag

import (
	"flag"
	"fmt"
)

type stringSliceValue struct {
	p     *[]string
	seted bool
}

func newStringSliceValue(val []string, p *[]string) *stringSliceValue {
	*p = append(*p, val...)
	return &stringSliceValue{p, false}
}

func (s *stringSliceValue) Set(value string) error {
	// has some default values and clear its.
	if s.seted == false && len(*s.p) > 0 {
		*s.p = (*s.p)[:0]
	}
	*s.p = append(*s.p, value)
	s.seted = true
	return nil
}

func (s *stringSliceValue) String() string {
	return fmt.Sprintf("%v", *s.p)
}

func (s *stringSliceValue) Get() interface{} {
	return *s.p
}

func String(flagset *flag.FlagSet, name string, value []string, usage string) *[]string {
	p := new([]string)
	StringVar(flagset, p, name, value, usage)
	return p
}

func StringVar(flagset *flag.FlagSet, p *[]string, name string, value []string, usage string) {
	flagset.Var(newStringSliceValue(value, p), name, usage)
}
