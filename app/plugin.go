package app

import (
	"errors"
	"log"
)

type ParsedArg map[string]string

type Arg struct {
	Key      string
	Required bool
	Default  string
}

type Plugin struct {
	Name        string
	Version     uint
	Description string

	Args []Arg

	Preload  func() error
	Generate func(args ParsedArg) ([]byte, error)
}

var (
	// contains registered plugins
	plugins = map[string]Plugin{}
)

// RegisterPlugin plugin and call Ready() when preload func done
func RegisterPlugin(corpus Plugin) {
	if err := corpus.Preload(); err != nil {
		log.Panicln(" > error preload:", err)
		return
	}

	plugins[corpus.Name] = corpus
}

// GetPlugins return all loaded plugins
func GetPlugins() []Plugin {
	res := make([]Plugin, 0)
	for _, corpus := range plugins {
		res = append(res, corpus)
	}

	return res
}

// GetPlugin by name
func GetPlugin(name string) (Plugin, error) {
	if val, ok := plugins[name]; ok {
		return val, nil
	}

	return Plugin{}, errors.New("plugin " + name + "not found")
}
