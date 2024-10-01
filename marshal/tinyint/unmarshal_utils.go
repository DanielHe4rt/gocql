package tinyint

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"
)

var errWrongDataLen = fmt.Errorf("failed to unmarshal tinyint: the length of the data should less or equal then 1")

func DecInt8(p []byte, v *int8) error {
	switch len(p) {
	case 0:
		*v = 0
	case 1:
		*v = int8(p[0])
	default:
		return errWrongDataLen
	}
	return nil
}

func DecInt8R(p []byte, v **int8) error {
	if p != nil {
		*v = new(int8)
		return DecInt8(p, *v)
	}
	*v = nil
	return nil
}

func DecInt16(p []byte, v *int16) error {
	switch len(p) {
	case 0:
		*v = 0
	case 1:
		*v = int16(int8(p[0]))
	default:
		return errWrongDataLen
	}
	return nil
}

func DecInt16R(p []byte, v **int16) error {
	if p != nil {
		*v = new(int16)
		return DecInt16(p, *v)
	}
	*v = nil
	return nil
}

func DecInt32(p []byte, v *int32) error {
	switch len(p) {
	case 0:
		*v = 0
	case 1:
		*v = int32(int8(p[0]))
	default:
		return errWrongDataLen
	}
	return nil
}

func DecInt32R(p []byte, v **int32) error {
	if p != nil {
		*v = new(int32)
		return DecInt32(p, *v)
	}
	*v = nil
	return nil
}

func DecInt64(p []byte, v *int64) error {
	switch len(p) {
	case 0:
		*v = 0
	case 1:
		*v = int64(int8(p[0]))
	default:
		return errWrongDataLen
	}
	return nil
}

func DecInt64R(p []byte, v **int64) error {
	if p != nil {
		*v = new(int64)
		return DecInt64(p, *v)
	}
	*v = nil
	return nil
}

func DecInt(p []byte, v *int) error {
	switch len(p) {
	case 0:
		*v = 0
	case 1:
		*v = int(int8(p[0]))
	default:
		return errWrongDataLen
	}
	return nil
}

func DecIntR(p []byte, v **int) error {
	if p != nil {
		*v = new(int)
		return DecInt(p, *v)
	}
	*v = nil
	return nil
}

func DecUint8(p []byte, v *uint8) error {
	switch len(p) {
	case 0:
		*v = 0
	case 1:
		*v = p[0]
	default:
		return errWrongDataLen
	}
	return nil
}

func DecUint8R(p []byte, v **uint8) error {
	if p != nil {
		*v = new(uint8)
		return DecUint8(p, *v)
	}
	*v = nil
	return nil
}

func DecUint16(p []byte, v *uint16) error {
	switch len(p) {
	case 0:
		*v = 0
	case 1:
		*v = uint16(p[0])
	default:
		return errWrongDataLen
	}
	return nil
}

func DecUint16R(p []byte, v **uint16) error {
	if p != nil {
		*v = new(uint16)
		return DecUint16(p, *v)
	}
	*v = nil
	return nil
}

func DecUint32(p []byte, v *uint32) error {
	switch len(p) {
	case 0:
		*v = 0
	case 1:
		*v = uint32(p[0])
	default:
		return errWrongDataLen
	}
	return nil
}

func DecUint32R(p []byte, v **uint32) error {
	if p != nil {
		*v = new(uint32)
		return DecUint32(p, *v)
	}
	*v = nil
	return nil
}

func DecUint64(p []byte, v *uint64) error {
	switch len(p) {
	case 0:
		*v = 0
	case 1:
		*v = uint64(p[0])
	default:
		return errWrongDataLen
	}
	return nil
}

func DecUint64R(p []byte, v **uint64) error {
	if p != nil {
		*v = new(uint64)
		return DecUint64(p, *v)
	}
	*v = nil
	return nil
}

func DecUint(p []byte, v *uint) error {
	switch len(p) {
	case 0:
		*v = 0
	case 1:
		*v = uint(p[0])
	default:
		return errWrongDataLen
	}
	return nil
}

func DecUintR(p []byte, v **uint) error {
	if p != nil {
		*v = new(uint)
		return DecUint(p, *v)
	}
	*v = nil
	return nil
}

func DecString(p []byte, v *string) error {
	switch len(p) {
	case 0:
		if p != nil {
			*v = "0"
		} else {
			*v = ""
		}
	case 1:
		*v = strconv.FormatInt(int64(int8(p[0])), 10)
	default:
		return errWrongDataLen
	}
	return nil
}

func DecStringR(p []byte, v **string) error {
	if p != nil {
		*v = new(string)
		return DecString(p, *v)
	}
	*v = nil
	return nil
}

func DecBigInt(p []byte, v *big.Int) error {
	switch len(p) {
	case 0:
		v.SetInt64(0)
	case 1:
		v.SetInt64(int64(int8(p[0])))
	default:
		return errWrongDataLen
	}
	return nil
}

func DecBigIntR(p []byte, v **big.Int) error {
	if p != nil {
		*v = big.NewInt(0)
		return DecBigInt(p, *v)
	}
	*v = nil
	return nil
}

func DecReflect(p []byte, v reflect.Value) error {
	if v.IsNil() {
		return fmt.Errorf("failed to unmarshal tinyint: can not unmarshal into nil reference (%T)(%#[1]v)", v.Interface())
	}

	switch v = v.Elem(); v.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return decReflectInts(p, v)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return decReflectUints(p, v)
	case reflect.String:
		return decReflectString(p, v)
	default:
		return fmt.Errorf("failed to unmarshal tinyint: unsupported value type (%T)(%#[1]v)", v.Interface())
	}
}

func DecReflectR(p []byte, v reflect.Value) error {
	if p != nil {
		zeroValue := reflect.New(v.Type().Elem().Elem())
		v.Elem().Set(zeroValue)
		return DecReflect(p, v.Elem())
	}
	nilValue := reflect.Zero(v.Elem().Type())
	v.Elem().Set(nilValue)
	return nil
}

func decReflectInts(p []byte, v reflect.Value) error {
	switch len(p) {
	case 0:
		v.SetInt(0)
	case 1:
		v.SetInt(int64(int8(p[0])))
	default:
		return errWrongDataLen
	}
	return nil
}

func decReflectUints(p []byte, v reflect.Value) error {
	switch len(p) {
	case 0:
		v.SetUint(0)
	case 1:
		v.SetUint(uint64(p[0]))
	default:
		return errWrongDataLen
	}
	return nil
}

func decReflectString(p []byte, v reflect.Value) error {
	switch len(p) {
	case 0:
		if p != nil {
			v.SetString("0")
		} else {
			v.SetString("")
		}
	case 1:
		v.SetString(strconv.FormatInt(int64(int8(p[0])), 10))
	default:
		return errWrongDataLen
	}
	return nil
}
