package component

type Component[T ComponentDeclaration] struct {
}

type ComponentDeclaration struct {
}

type ComponentBuilder struct {
	Make func(name string) BuilderContext
}

type BuilderContext struct {
  properties: map[string]Property{}
}

type PropertyType string;

type PropertyName string

validate.Property("allowed-property-names",
  func (propertyName PropertyName) {
  if (propertyName == "") {
    panic("Property name cannot be empty")
  }
}

type Property func(name string, PropertyAllowedValues)

type PropertyAllowedValues struct {

}

func (bc *BuilderContext) Property(name: PropertyName)
