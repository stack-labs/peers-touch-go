package config

import "sync"

var (
	DefaultStartPath = ""
	DefaultSeparator = "."
	mux              sync.Mutex
	configs          = &allConfigs{
		values: map[string]Options{},
	}
)

type allConfigs struct {
	values map[string]Options
}

type Options interface {
	Path() string
}

type AbstractOptions struct {
	Subs map[string]Options
}

func (a *AbstractOptions)

func Register(options Options) {
	mux.Lock()
	defer mux.Unlock()

	configs.values[options.Path()] = options
}

func parse() {

}
