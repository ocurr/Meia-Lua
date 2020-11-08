package types

func IsNumber(t Type) bool {
	switch t.(type) {
	case Int, Float:
		return true
	default:
		return false
	}
}

func IsBool(t Type) bool {
	switch t.(type) {
	case Bool:
		return true
	default:
		return false
	}
}

func IsFloat(t Type) bool {
	switch t.(type) {
	case Float:
		return true
	}
	return false
}
