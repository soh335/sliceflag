package sliceflag

import (
	"flag"
	"fmt"
	"time"
)

type durationSliceValue struct {
	p     *[]time.Duration
	seted bool
}

func newDurationSliceValue(val []time.Duration, p *[]time.Duration) *durationSliceValue {
	*p = append(*p, val...)
	return &durationSliceValue{p, false}
}

func (d *durationSliceValue) Set(value string) error {
	v, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	if d.seted == false && len(*d.p) > 0 {
		*d.p = (*d.p)[:0]
	}
	*d.p = append(*d.p, v)
	d.seted = true
	return nil
}

func (d *durationSliceValue) String() string {
	return fmt.Sprintf("%v", *d.p)
}

func (d *durationSliceValue) Get() interface{} {
	return *d.p
}

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

func Duration(flagset *flag.FlagSet, name string, value []time.Duration, usage string) *[]time.Duration {
	p := new([]time.Duration)
	DurationVar(flagset, p, name, value, usage)
	return p
}

func DurationVar(flagset *flag.FlagSet, p *[]time.Duration, name string, value []time.Duration, usage string) {
	flagset.Var(newDurationSliceValue(value, p), name, usage)
}

func String(flagset *flag.FlagSet, name string, value []string, usage string) *[]string {
	p := new([]string)
	StringVar(flagset, p, name, value, usage)
	return p
}

func StringVar(flagset *flag.FlagSet, p *[]string, name string, value []string, usage string) {
	flagset.Var(newStringSliceValue(value, p), name, usage)
}
