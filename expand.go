package expand

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

// QueryParametersConfig struct
type QueryParametersConfig struct {
	Host            string           `yaml:"host"`
	Path            string           `yaml:"path"`
	QueryParameters []QueryParameter `yaml:"query-parameters"`
}

// QueryParameter struct
type QueryParameter struct {
	Name        string      `yaml:"name"`
	Type        string      `yaml:"type"`
	NullAble    bool        `yaml:"null-able"`
	Values      []string    `yaml:"values"`
	RangeConfig RangeConfig `yaml:"range-config"`
}

// RangeConfig struct
type RangeConfig struct {
	Start int64 `yaml:"start"`
	Stop  int64 `yaml:"stop"`
	Step  int64 `yaml:"step"`
}

// Parse method
func Parse(filePath string) (QueryParametersConfig, error) {
	var qpc QueryParametersConfig

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return qpc, err
	}

	err = yaml.Unmarshal(buf, &qpc)
	if err != nil {
		return qpc, err
	}

	return qpc, nil
}
