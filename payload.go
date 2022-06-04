package decimal128

import "fmt"

const (
	payloadOpNaN Payload = iota + 1
	payloadOpParse
	payloadOpScan
	payloadOpUnmarshalText

	payloadOpAdd
	payloadOpMul
	payloadOpQuo
	payloadOpSub
)

const (
	payloadValPosZero Payload = iota + 1
	payloadValNegZero
	payloadValPosFinite
	payloadValNegFinite
	payloadValPosInfinite
	payloadValNegInfinite
)

func (d Decimal) Payload() Payload {
	if !d.IsNaN() {
		panic("Decimal(!NaN).Payload()")
	}

	return Payload(d.lo)
}

// Payload represents the payload value of a NaN decimal. This value can
// contain additional information about the operation that caused the value to
// be set to NaN.
type Payload uint64

// String returns a string representation of the payload.
func (p Payload) String() string {
	if p == 0 {
		return "Payload(0)"
	}

	if p > 0x00ff_ffff {
		return fmt.Sprintf("Payload(%d)", uint64(p))
	}

	switch p & 0xff {
	case payloadOpNaN:
		return "NaN()"
	case payloadOpParse:
		return "Parse()"
	case payloadOpScan:
		return "Scan()"
	case payloadOpUnmarshalText:
		return "UnmarshalText()"
	case payloadOpAdd:
		return "Add(" + p.argString(8) + ", " + p.argString(16) + ")"
	case payloadOpMul:
		return "Mul(" + p.argString(8) + ", " + p.argString(16) + ")"
	case payloadOpQuo:
		return "Quo(" + p.argString(8) + ", " + p.argString(16) + ")"
	case payloadOpSub:
		return "Sub(" + p.argString(8) + ", " + p.argString(16) + ")"
	default:
		return fmt.Sprintf("Payload(%d)", uint64(p))
	}
}

func (p Payload) argString(offset int) string {
	switch p >> offset & 0xff {
	case payloadValPosZero:
		return "Zero"
	case payloadValNegZero:
		return "-Zero"
	case payloadValPosFinite:
		return "Finite"
	case payloadValNegFinite:
		return "-Finite"
	case payloadValPosInfinite:
		return "Infinite"
	case payloadValNegInfinite:
		return "-Infinite"
	default:
		return "Unknown"
	}
}
