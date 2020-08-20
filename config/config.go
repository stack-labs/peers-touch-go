package config

var (
	DefaultStartPath = ""
)

type allConfigs struct {
	values 
}


type Options interface {
	Path() string
	Parse(options Options)
}

func Parse(options Options) {
	configs := allConfigs.
}
