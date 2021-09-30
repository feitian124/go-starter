package schema

type Column struct {
	ColumnName string
	DefaultValue interface{}
	IsNullable bool
	DataType string
	NumericPrecision int
	NumericScale int
	MaxLength int
	Extra string
	Comment string
	Ordinal int
}