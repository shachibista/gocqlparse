package ast

type GenericType interface{}

type NativeType struct {
	Name string
}

type CustomType struct {
	Name string
}

type ListType struct {
	Elem GenericType
}

type SetType struct {
	Elem GenericType
}

type MapType struct {
	Key   GenericType
	Value GenericType
}

type TupleType struct {
	Elements []GenericType
}

type VectorType struct {
	Elem       GenericType
	Dimensions int
}

type ObjectRef struct {
	Keyspace Identifier
	Name     Identifier
}

type FrozenType struct {
	Elem GenericType
}

type Marker struct {
	Name Identifier
}

func (o *ObjectRef) Equals(r *ObjectRef) bool {
	if o.Keyspace == nil && r.Keyspace != nil {
		return false
	}

	if o.Keyspace != nil && r.Keyspace == nil {
		return false
	}

	if o.Keyspace == nil && r.Keyspace == nil {
		return o.Name.String() == r.Name.String()
	}

	return o.Keyspace.String() == r.Keyspace.String() && o.Name.String() == r.Name.String()
}
