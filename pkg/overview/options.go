package overview

type Options struct {
	Namespace    string
	AllNamespace bool
}

func NewDefaultOptions() *Options {
	return &Options{
		AllNamespace: false,
		Namespace:    "default",
	}
}
