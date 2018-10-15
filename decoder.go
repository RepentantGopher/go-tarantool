package tarantool

import (
	"io"
	"reflect"

	"gopkg.in/vmihailenco/msgpack.v2"
)

type Decoder interface {
	Decode(v ...interface{}) error
	Reset(r io.Reader) error
	DecodeMapLen() (int, error)
	DecodeMap() (interface{}, error)
	DecodeSliceLen() (int, error)
	DecodeArrayLen() (int, error)
	DecodeSlice() ([]interface{}, error)
	DecodeBytesLen() (int, error)
	DecodeBytes() ([]byte, error)
	DecodeInt() (int, error)
	DecodeUint64() (uint64, error)
	DecodeInt64() (int64, error)
	DecodeFloat32() (float32, error)
	DecodeFloat64() (float64, error)
	DecodeUint() (uint, error)
	DecodeUint8() (uint8, error)
	DecodeUint16() (uint16, error)
	DecodeUint32() (uint32, error)
	DecodeInt8() (int8, error)
	DecodeInt16() (int16, error)
	DecodeInt32() (int32, error)
	Query(query string) ([]interface{}, error)
	DecodeInterface() (interface{}, error)
	DecodeString() (string, error)
	DecodeValue(v reflect.Value) error
	DecodeBool() (bool, error)
	Skip() error
	PeekCode() (code byte, err error)
}

type DecoderFactory func(r io.Reader) Decoder

var DefaultDecoder = func(r io.Reader) Decoder { return msgpack.NewDecoder(r) }
