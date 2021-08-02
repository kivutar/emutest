// Package options deals with configuration at the libretro core level. Each
// core exports a list of variables that can take different values. This package
// can list them, load them, and save them.
package options

import (
	"io/ioutil"
	"strings"
	"sync"

	"github.com/kivutar/emutest/state"
	"github.com/kivutar/emutest/utils"

	"github.com/pelletier/go-toml"
)

// Variable represents one core option. A variable can take a limited number of
// values. The possibilities are stored in v.Choices. The current value
// can be accessed with v.Choices[v.Choice]
type Variable struct {
	Key     string   // unique id of the variable
	Desc    string   // human readable name of the variable
	Choices []string // available values
	Choice  int      // index of the current value
	Default string
}

// Options is a container type for core options internals
type Options struct {
	Vars    []*Variable // the variables exposed by the core
	Updated bool        // notify the core that values have been updated

	sync.Mutex
}

// VariableInterface is used as a compatibility layer for old v0 options and v1 options
type VariableInterface interface {
	Key() string
	Desc() string
	Choices() []string
	DefaultValue() string
}

// New instantiate a core options manager
func New(vars []VariableInterface) (*Options, error) {
	o := &Options{}

	// Cache core options
	for _, v := range vars {
		v := v
		o.Vars = append(o.Vars, &Variable{
			Key:     v.Key(),
			Desc:    v.Desc(),
			Choices: v.Choices(),
			Default: v.DefaultValue(),
			Choice:  utils.IndexOfString(v.DefaultValue(), v.Choices()),
		})
	}
	o.Updated = true
	err := o.Load()
	return o, err
}

// Load core options from a file
func (o *Options) Load() error {
	o.Lock()
	defer o.Unlock()

	if state.OptionsPath == "" && state.OptionsToml == "" {
		return nil
	}

	var opts map[string]string

	if state.OptionsPath != "" {
		b, err := ioutil.ReadFile(state.OptionsPath)
		if err != nil {
			return err
		}

		if err := toml.Unmarshal(b, &opts); err != nil {
			return err
		}
	}

	if state.OptionsToml != "" {
		if err := toml.Unmarshal([]byte(state.OptionsToml), &opts); err != nil {
			return err
		}
	}

	for optk, optv := range opts {
		for _, variable := range o.Vars {
			if variable.Key == strings.Replace(optk, "___", ".", 1) {
				for j, c := range variable.Choices {
					if c == optv {
						variable.Choice = j
					}
				}
			}
		}
	}

	return nil
}
