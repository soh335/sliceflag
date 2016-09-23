package sliceflag

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// -- duration slice value
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
	return sliceString(d.p)
}

func (d *durationSliceValue) Get() interface{} {
	return *d.p
}

// -- float64 slice value
type float64SliceValue struct {
	p     *[]float64
	seted bool
}

func newFloat64SliceValue(val []float64, p *[]float64) *float64SliceValue {
	*p = append(*p, val...)
	return &float64SliceValue{p, false}
}

func (f *float64SliceValue) Set(value string) error {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	// has some default values and clear its.
	if f.seted == false && len(*f.p) > 0 {
		*f.p = (*f.p)[:0]
	}
	*f.p = append(*f.p, v)
	f.seted = true
	return nil
}

func (f *float64SliceValue) String() string {
	return sliceString(f.p)
}

func (f *float64SliceValue) Get() interface{} {
	return *f.p
}

// -- int slice value
type intSliceValue struct {
	p     *[]int
	seted bool
}

func newIntSliceVlaue(val []int, p *[]int) *intSliceValue {
	*p = append(*p, val...)
	return &intSliceValue{p, false}
}

func (i *intSliceValue) Set(value string) error {
	v, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return err
	}
	// has some default values and clear its.
	if i.seted == false && len(*i.p) > 0 {
		*i.p = (*i.p)[:0]
	}
	*i.p = append(*i.p, int(v))
	i.seted = true
	return nil
}

func (i *intSliceValue) String() string {
	return sliceString(i.p)
}

func (i *intSliceValue) Get() interface{} {
	return *i.p
}

// -- int64 slice value
type int64SliceValue struct {
	p     *[]int64
	seted bool
}

func newInt64SliceVlaue(val []int64, p *[]int64) *int64SliceValue {
	*p = append(*p, val...)
	return &int64SliceValue{p, false}
}

func (i *int64SliceValue) Set(value string) error {
	v, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return err
	}
	// has some default values and clear its.
	if i.seted == false && len(*i.p) > 0 {
		*i.p = (*i.p)[:0]
	}
	*i.p = append(*i.p, v)
	i.seted = true
	return nil
}

func (i *int64SliceValue) String() string {
	return sliceString(i.p)
}

func (i *int64SliceValue) Get() interface{} {
	return *i.p
}

// -- uint slice value
type uintSliceValue struct {
	p     *[]uint
	seted bool
}

func newUintSliceVlaue(val []uint, p *[]uint) *uintSliceValue {
	*p = append(*p, val...)
	return &uintSliceValue{p, false}
}

func (u *uintSliceValue) Set(value string) error {
	v, err := strconv.ParseUint(value, 0, 64)
	if err != nil {
		return err
	}
	// has some default values and clear its.
	if u.seted == false && len(*u.p) > 0 {
		*u.p = (*u.p)[:0]
	}
	*u.p = append(*u.p, uint(v))
	u.seted = true
	return nil
}

func (u *uintSliceValue) String() string {
	return sliceString(u.p)
}

func (u *uintSliceValue) Get() interface{} {
	return *u.p
}

// -- uint64 slice value
type uint64SliceValue struct {
	p     *[]uint64
	seted bool
}

func newUint64SliceVlaue(val []uint64, p *[]uint64) *uint64SliceValue {
	*p = append(*p, val...)
	return &uint64SliceValue{p, false}
}

func (u *uint64SliceValue) Set(value string) error {
	v, err := strconv.ParseUint(value, 0, 64)
	if err != nil {
		return err
	}
	// has some default values and clear its.
	if u.seted == false && len(*u.p) > 0 {
		*u.p = (*u.p)[:0]
	}
	*u.p = append(*u.p, v)
	u.seted = true
	return nil
}

func (u *uint64SliceValue) String() string {
	return sliceString(u.p)
}

func (u *uint64SliceValue) Get() interface{} {
	return *u.p
}

// -- string slice value
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
	return sliceString(s.p)
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

func Float64(flagset *flag.FlagSet, name string, value []float64, usage string) *[]float64 {
	p := new([]float64)
	Float64Var(flagset, p, name, value, usage)
	return p
}

func Float64Var(flagset *flag.FlagSet, p *[]float64, name string, value []float64, usage string) {
	flagset.Var(newFloat64SliceValue(value, p), name, usage)
}

func Int(flagset *flag.FlagSet, name string, value []int, usage string) *[]int {
	p := new([]int)
	IntVar(flagset, p, name, value, usage)
	return p
}

func IntVar(flagset *flag.FlagSet, p *[]int, name string, value []int, usage string) {
	flagset.Var(newIntSliceVlaue(value, p), name, usage)
}

func Int64(flagset *flag.FlagSet, name string, value []int64, usage string) *[]int64 {
	p := new([]int64)
	Int64Var(flagset, p, name, value, usage)
	return p
}

func Int64Var(flagset *flag.FlagSet, p *[]int64, name string, value []int64, usage string) {
	flagset.Var(newInt64SliceVlaue(value, p), name, usage)
}

func Uint(flagset *flag.FlagSet, name string, value []uint, usage string) *[]uint {
	p := new([]uint)
	UintVar(flagset, p, name, value, usage)
	return p
}

func UintVar(flagset *flag.FlagSet, p *[]uint, name string, value []uint, usage string) {
	flagset.Var(newUintSliceVlaue(value, p), name, usage)
}

func Uint64(flagset *flag.FlagSet, name string, value []uint64, usage string) *[]uint64 {
	p := new([]uint64)
	Uint64Var(flagset, p, name, value, usage)
	return p
}

func Uint64Var(flagset *flag.FlagSet, p *[]uint64, name string, value []uint64, usage string) {
	flagset.Var(newUint64SliceVlaue(value, p), name, usage)
}

func String(flagset *flag.FlagSet, name string, value []string, usage string) *[]string {
	p := new([]string)
	StringVar(flagset, p, name, value, usage)
	return p
}

func StringVar(flagset *flag.FlagSet, p *[]string, name string, value []string, usage string) {
	flagset.Var(newStringSliceValue(value, p), name, usage)
}

func sliceString(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "[]"
	} else {
		return fmt.Sprintf("%v", rv.Elem())
	}
}
