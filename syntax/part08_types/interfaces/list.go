package interfaces

// List List接口
type List interface {
	Add(index int, value any) error
	Append(value any)
	Delete(index int) (any, error)
}
