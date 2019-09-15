package expand

import (
	"fmt"
	"io/ioutil"
	"strconv"

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

// ToStrArray method
func (qp *QueryParameter) ToStrArray() ([]string, error) {
	qpStrArray := []string{}

	switch qp.Type {
	case "list":
		for _, v := range qp.Values {
			qpStr := qp.Name + "=" + v
			qpStrArray = append(qpStrArray, qpStr)
		}
	case "range":
		for i := qp.RangeConfig.Start; i <= qp.RangeConfig.Stop; i += qp.RangeConfig.Step {
			qpStr := qp.Name + "=" + strconv.FormatInt(i, 10)
			qpStrArray = append(qpStrArray, qpStr)
		}
	default:
		err := fmt.Errorf("%s: type of QueryParameter is invalid", qp.Name)
		return qpStrArray, err
	}

	if qp.NullAble {
		qpStrArray = append(qpStrArray, "")
	}

	return qpStrArray, nil
}

// Create method
func (qpc *QueryParametersConfig) Create() ([]string, error) {
	var urlArray []string

	var qpStr2DArray [][]string
	for _, qp := range qpc.QueryParameters {
		qpStrArray, err := qp.ToStrArray()
		if err != nil {
			return urlArray, err
		}

		qpStr2DArray = append(qpStr2DArray, qpStrArray)
	}

	urlArray = qpc.AddRequestPath(Cartesian(qpStr2DArray))

	return urlArray, nil
}

// AddRequestPath method
func (qpc *QueryParametersConfig) AddRequestPath(strArray []string) []string {
	var result []string

	for _, s := range strArray {
		url := "http://" + qpc.Host + qpc.Path + "?" + s
		result = append(result, url)
	}

	return result
}
