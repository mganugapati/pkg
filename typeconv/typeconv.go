package typeconv

import (
	"strconv"
)

const Unused = 0

func String2Int(v *int, s string) error {
	i, err := strconv.Atoi(s)
	if err == nil {
		*v = i
	}
	return err
}

func String2Int8(v *int8, s string) error {
	i, err := strconv.ParseInt(s, 10, 8)
	if err == nil {
		*v = int8(i)
	}
	return err
}

func String2Int16(v *int16, s string) error {
	i, err := strconv.ParseInt(s, 10, 16)
	if err == nil {
		*v = int16(i)
	}
	return err
}

func String2Int32(v *int32, s string) error {
	i, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		*v = int32(i)
	}
	return err
}

func String2Int64(v *int64, s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = int64(i)
	}
	return err
}

func String2Uint(v *uint, s string) error {
	i, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		*v = uint(i)
	}
	return err
}

func String2Uint8(v *uint8, s string) error {
	i, err := strconv.ParseUint(s, 10, 8)
	if err == nil {
		*v = uint8(i)
	}
	return err
}

func String2Uint16(v *uint16, s string) error {
	i, err := strconv.ParseUint(s, 10, 16)
	if err == nil {
		*v = uint16(i)
	}
	return err
}

func String2Uint32(v *uint32, s string) error {
	i, err := strconv.ParseUint(s, 10, 32)
	if err == nil {
		*v = uint32(i)
	}
	return err
}

func String2Uint64(v *uint64, s string) error {
	i, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		*v = i
	}
	return err
}

func String2Bool(v *bool, s string) error {
	b, err := strconv.ParseBool(s)
	if err == nil {
		*v = b
	}
	return err
}

func String2Float32(v *float32, s string) error {
	f, err := strconv.ParseFloat(s, 32)
	if err == nil {
		*v = float32(f)
	}
	return err
}

func String2Float64(v *float64, s string) error {
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		*v = f
	}
	return err
}
