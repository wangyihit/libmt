package processor

type Variant interface{}

type Processor interface {
	Run(data Variant) (Variant, error)
	Name() string
}
