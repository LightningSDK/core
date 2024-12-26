package service

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
)

type Plugins map[string][]byte

func (e *Plugins) UnmarshalYAML(n *yaml.Node) error {
	for i, nn := range n.Content {
		if i%2 == 1 {
			continue
		}
		if nn.Value != "plugins" {
			continue
		}
		plugins := n.Content[i+1]
		if plugins.Tag != "!!map" {
			panic("expecting plugins section of configuration to be a map")
		}

		(*e) = map[string][]byte{}
		for j, c := range plugins.Content {
			if j%2 == 1 {
				continue
			}
			if c.Tag != "!!str" {
				panic("expecting a key for the plugin configuration")
			}
			(*e)[c.Value] = getJsonPluginConfig(plugins.Content[j+1])
		}

	}

	return nil
}

// todo: does this need other types like int/string/null?
func getJsonPluginConfig(n *yaml.Node) []byte {
	switch n.Tag {
	case "!!map":
		cfg := map[string]interface{}{}
		return remarshal(cfg, n)
	case "!!list":
		cfg := []interface{}{}
		return remarshal(cfg, n)
	}
	return nil
}

func remarshal(ti interface{}, n *yaml.Node) []byte {
	err := n.Decode(ti)
	if err != nil {
		panic(err)
	}
	enc, err := json.Marshal(ti)
	if err != nil {
		panic(err)
	}
	return enc
}

type Config struct {
	Plugins `yaml:"plugins"`
	Include []string `yaml:"include"`
}

// loads all the raw config data
func loadConfigurations(configPath string) *Config {
	cfg := &Config{}

	// read the source file
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		panic(err) // or handle the error appropriately
	}

	// unmarshal into the struct
	err = yaml.Unmarshal(yamlFile, &cfg)

	return cfg
}

// loads the config for a specific plugin
func (cfgs *Config) getConfig(plugin string) []byte {
	if cfg, ok := cfgs.Plugins[plugin]; ok {
		return cfg
	}
	return []byte("{}")
}

//// BootstrapConfig this should be the very first function called.
//// It loads the auto-generated plugin list and passes the primary config yaml file.
//func BootstrapConfig(f string, initter func(app *App) map[string]model.Plugin) (*App, error) {
//	// load the yaml file
//	y, err := os.ReadFile(f)
//	if err != nil {
//		return nil, err
//	}
//	ys := [][]byte{y}
//	_ = map[string]bool{
//		f: true,
//	}
//
//	// unmarshal main, this will get a list of includes
//	// TODO: this should run recursively?
//	app := &App{
//		Include: []string{},
//	}
//
//	app.Plugins = initter(app)
//
//	err = yaml.Unmarshal(y, app)
//	if err != nil {
//		return nil, err
//	}
//
//	// iterate through each of the plugins. if it has its own config,
//	// then the yaml should be unmarshalled to it
//	for _, m := range app.Plugins {
//		if c := m.GetEmptyConfig(); c != nil {
//			for _, y := range ys {
//				temp := m.GetEmptyConfig()
//				err = defaults.Set(temp)
//				if err != nil {
//					return nil, err
//				}
//				err = yaml.Unmarshal(y, temp)
//				if err != nil {
//					return nil, err
//				}
//				err = mergo.Merge(c, temp)
//				if err != nil {
//					return nil, err
//				}
//				m.SetConfig(c)
//			}
//		}
//	}
//
//	return app, nil
//}
