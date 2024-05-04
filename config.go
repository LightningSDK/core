package core

import (
	"dario.cat/mergo"
	"gopkg.in/yaml.v3"
	"os"
)

type App struct {
	Include []string `yaml:"include"`
	Modules Modules  `yaml:"modules"`
}

// BootstrapConfig this should be the very first function called.
// It loads the auto-generated module list and passes the primary config yaml file.
func BootstrapConfig(f string, initter func(app *App) map[string]Module) (*App, error) {
	// load the yaml file
	y, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	ys := [][]byte{y}
	_ = map[string]bool{
		f: true,
	}

	// unmarshal main, this will get a list of includes
	// TODO: this should run recursively?
	app := &App{
		Include: []string{},
	}

	app.Modules = initter(app)

	err = yaml.Unmarshal(y, app)
	if err != nil {
		return nil, err
	}

	// iterate through each of the modules. if it has its own config,
	// then the yaml should be unmarshalled to it
	for _, m := range app.Modules {
		if c := m.GetEmptyConfig(); c != nil {
			for _, y := range ys {
				temp := m.GetEmptyConfig()
				err = yaml.Unmarshal(y, temp)
				if err != nil {
					return nil, err
				}
				err = mergo.Merge(c, temp)
				if err != nil {
					return nil, err
				}
				m.SetConfig(c)
			}
		}
	}

	return app, nil
}
