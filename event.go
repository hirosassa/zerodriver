package zerodriver

import (
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog"
)

func (e *Event) AnErr(key string, err error) *Event {
	e.event = e.event.AnErr(key, err)
	return e
}

func (e *Event) Array(key string, arr zerolog.LogArrayMarshaler) *Event {
	e.event = e.event.Array(key, arr)
	return e
}

func (e *Event) Bool(key string, b bool) *Event {
	e.event = e.event.Bool(key, b)
	return e
}

func (e *Event) Bools(key string, b []bool) *Event {
	e.event = e.event.Bools(key, b)
	return e
}

func (e *Event) Bytes(key string, val []byte) *Event {
	e.event = e.event.Bytes(key, val)
	return e
}

func (e *Event) Caller(skip ...int) *Event {
	e.event = e.event.Caller(skip...)
	return e
}

func (e *Event) Dict(key string, dict *zerolog.Event) *Event {
	e.event = e.event.Dict(key, dict)
	return e
}

func (e *Event) Discard() *Event {
	e.event = e.event.Discard()
	return e
}

func (e *Event) Dur(key string, d time.Duration) *Event {
	e.event = e.event.Dur(key, d)
	return e
}

func (e *Event) Durs(key string, d []time.Duration) *Event {
	e.event = e.event.Durs(key, d)
	return e
}

func (e *Event) EmbedObject(obj zerolog.LogObjectMarshaler) *Event {
	e.event = e.event.EmbedObject(obj)
	return e
}

func (e *Event) Enabled() bool {
	return e.Enabled()
}

func (e *Event) Err(err error) *Event {
	e.event = e.event.Err(err)
	return e
}

func (e *Event) Errs(key string, errs []error) *Event {
	e.event = e.event.Errs(key, errs)
	return e
}

func (e *Event) Fields(fields map[string]interface{}) *Event {
	e.event = e.event.Fields(fields)
	return e
}

func (e *Event) Float32(key string, f float32) *Event {
	e.event = e.event.Float32(key, f)
	return e
}

func (e *Event) Float64(key string, f float64) *Event {
	e.event = e.event.Float64(key, f)
	return e
}

func (e *Event) Floats32(key string, f []float32) *Event {
	e.event = e.event.Floats32(key, f)
	return e
}

func (e *Event) Floats64(key string, f []float64) *Event {
	e.event = e.event.Floats64(key, f)
	return e
}

func (e *Event) Hex(key string, val []byte) *Event {
	e.event = e.event.Hex(key, val)
	return e
}

func (e *Event) IPAddr(key string, ip net.IP) *Event {
	e.event = e.event.IPAddr(key, ip)
	return e
}

func (e *Event) IPPrefix(key string, pfx net.IPNet) *Event {
	e.event = e.event.IPPrefix(key, pfx)
	return e
}

func (e *Event) Int(key string, i int) *Event {
	e.event = e.event.Int(key, i)
	return e
}

func (e *Event) Int16(key string, i int16) *Event {
	e.event = e.event.Int16(key, i)
	return e
}

func (e *Event) Int32(key string, i int32) *Event {
	e.event = e.event.Int32(key, i)
	return e
}

func (e *Event) Int64(key string, i int64) *Event {
	e.event = e.event.Int64(key, i)
	return e
}

func (e *Event) Int8(key string, i int8) *Event {
	e.event = e.event.Int8(key, i)
	return e
}

func (e *Event) Interface(key string, i interface{}) *Event {
	e.event = e.event.Interface(key, i)
	return e
}

func (e *Event) Ints(key string, i []int) *Event {
	e.event = e.event.Ints(key, i)
	return e
}

func (e *Event) Ints16(key string, i []int16) *Event {
	e.event = e.event.Ints16(key, i)
	return e
}

func (e *Event) Ints32(key string, i []int32) *Event {
	e.event = e.event.Ints32(key, i)
	return e
}

func (e *Event) Ints64(key string, i []int64) *Event {
	e.event = e.event.Ints64(key, i)
	return e
}

func (e *Event) Ints8(key string, i []int8) *Event {
	e.event = e.event.Ints8(key, i)
	return e
}

func (e *Event) MACAddr(key string, ha net.HardwareAddr) *Event {
	e.event = e.event.MACAddr(key, ha)
	return e
}

func (e *Event) Msg(msg string) {
	e.setLabels()
	e.event.Msg(msg)
}

func (e *Event) Msgf(format string, v ...interface{}) {
	e.setLabels()
	e.event.Msgf(format, v...)
}

func (e *Event) Object(key string, obj zerolog.LogObjectMarshaler) *Event {
	e.event = e.event.Object(key, obj)
	return e
}

func (e *Event) RawJSON(key string, b []byte) *Event {
	e.event = e.event.RawJSON(key, b)
	return e
}

func (e *Event) Send() {
	e.setLabels()
	e.event.Send()
}

func (e *Event) Stack() *Event {
	e.event = e.event.Stack()
	return e
}

func (e *Event) Str(key, val string) *Event {
	e.event = e.event.Str(key, val)
	return e
}

func (e *Event) Stringer(key string, val fmt.Stringer) *Event {
	e.event = e.event.Stringer(key, val)
	return e
}

func (e *Event) Strs(key string, vals []string) *Event {
	e.event = e.event.Strs(key, vals)
	return e
}

func (e *Event) Time(key string, t time.Time) *Event {
	e.event = e.event.Time(key, t)
	return e
}

func (e *Event) TimeDiff(key string, t time.Time, start time.Time) *Event {
	e.event = e.event.TimeDiff(key, t, start)
	return e
}

func (e *Event) Times(key string, t []time.Time) *Event {
	e.event = e.event.Times(key, t)
	return e
}

func (e *Event) Timestamp() *Event {
	e.event = e.event.Timestamp()
	return e
}

func (e *Event) Uint(key string, i uint) *Event {
	e.event = e.event.Uint(key, i)
	return e
}

func (e *Event) Uint16(key string, i uint16) *Event {
	e.event = e.event.Uint16(key, i)
	return e
}

func (e *Event) Uint32(key string, i uint32) *Event {
	e.event = e.event.Uint32(key, i)
	return e
}

func (e *Event) Uint64(key string, i uint64) *Event {
	e.event = e.event.Uint64(key, i)
	return e
}

func (e *Event) Uint8(key string, i uint8) *Event {
	e.event = e.event.Uint8(key, i)
	return e
}

func (e *Event) Uints(key string, i []uint) *Event {
	e.event = e.event.Uints(key, i)
	return e
}

func (e *Event) Uints16(key string, i []uint16) *Event {
	e.event = e.event.Uints16(key, i)
	return e
}

func (e *Event) Uints32(key string, i []uint32) *Event {
	e.event = e.event.Uints32(key, i)
	return e
}

func (e *Event) Uints64(key string, i []uint64) *Event {
	e.event = e.event.Uints64(key, i)
	return e
}

func (e *Event) Uints8(key string, i []uint8) *Event {
	e.event = e.event.Uints8(key, i)
	return e
}
