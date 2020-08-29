package config

import (
	"sort"
	"strings"
	"sync"
)

var (
	DefaultStartPath = "touch"
	DefaultSeparator = "."
	mux              sync.Mutex
	configs          = &allConfigs{
		valuesMap: map[string]Options{},
	}
)

type allConfigs struct {
	path      string
	mux       sync.Mutex
	valuesMap map[string]Options
	options   Options
	values    []Options
}

func (a *allConfigs) init() {
	if len(a.valuesMap) > 0 {
		a.sort()
		a.parse()
	}
}

func (a *allConfigs) sort() {
	mux.Lock()
	defer mux.Unlock()

	keys := make([]string, 0, len(a.valuesMap))
	for k, _ := range a.valuesMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, key := range keys {
		a.values = append(a.values, a.valuesMap[key])
	}
}

func (a *allConfigs) parse() {
	// the first one is root path
	a.options = a.values[0]

	// the rest of options
	if len(a.values) > 1 {
		restOptions := a.values[1:]
		iterate(a.options, restOptions)
	}
}

func iterate(thisOptions Options, restOptions []Options) {
	var lastPath = thisOptions.Path() + DefaultSeparator
	tempRest := make([]Options, 0)
	for _, options := range restOptions {
		tempStrs := strings.Split(options.Path(), lastPath)
		if len(tempStrs) > 1 && !strings.Contains(tempStrs[1], DefaultSeparator) {
			thisOptions.AddSub(options)
		} else {
			tempRest = append(tempRest, options)
		}
	}

	if len(tempRest) > 0 {
		for _, options := range thisOptions.subs() {
			iterate(options, tempRest)
		}
	}
}

func remove(slice []Options, s int) []Options {
	if s+1 > len(slice) {
		return nil
	}
	return append(slice[:s], slice[s+1:]...)
}

type Options interface {
	Path() string
	AddSub(options Options)
	subs() []Options
}

type SubOptions map[string]Options

type AbstractOptions struct {
	Options
	mux sync.Mutex
	SubOptions
}

func (a *AbstractOptions) AddSub(options Options) {
	mux.Lock()
	defer mux.Unlock()

	if a.SubOptions == nil {
		a.SubOptions = SubOptions{}
	}

	a.SubOptions[options.Path()] = options
}

func (a *AbstractOptions) Subs() []Options {
	v := make([]Options, 0, len(a.SubOptions))

	for _, value := range a.SubOptions {
		v = append(v, value)
	}

	return v
}

func Register(options ...Options) {
	mux.Lock()
	defer mux.Unlock()

	for _, option := range options {
		configs.valuesMap[option.Path()] = option
	}
}

func parse() {

}
