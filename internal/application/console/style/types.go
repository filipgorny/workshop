package style

import "go/types"

type StyleValue struct {
	bytes  []*byte
	length int
}

type StyleType struct {
	value    StyleTypeValue
	typeName types.Type
}

type StyleTypeFactory func(maxLength int, typeName types.Type) StyleType
