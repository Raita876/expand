package expand

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var expectedQpc = QueryParametersConfig{
	Host: "localhost:8888",
	Path: "/sample",
	QueryParameters: []QueryParameter{
		QueryParameter{
			Name:     "name1",
			Type:     "list",
			NullAble: true,
			Values: []string{
				"AAAAA",
				"BBBBB",
				"CCCCC",
			},
		},
		QueryParameter{
			Name:     "name2",
			Type:     "list",
			NullAble: false,
			Values: []string{
				"DDDDD",
				"EEEEE",
			},
		},
		QueryParameter{
			Name:     "name3",
			Type:     "range",
			NullAble: false,
			RangeConfig: RangeConfig{
				Start: 0,
				Stop:  30,
				Step:  10,
			},
		},
	},
}

var expectedURLArray = []string{
	"http://localhost:8888/sample?name1=AAAAA&name2=DDDDD&name3=0",
	"http://localhost:8888/sample?name1=AAAAA&name2=DDDDD&name3=10",
	"http://localhost:8888/sample?name1=AAAAA&name2=DDDDD&name3=20",
	"http://localhost:8888/sample?name1=AAAAA&name2=DDDDD&name3=30",
	"http://localhost:8888/sample?name1=AAAAA&name2=EEEEE&name3=0",
	"http://localhost:8888/sample?name1=AAAAA&name2=EEEEE&name3=10",
	"http://localhost:8888/sample?name1=AAAAA&name2=EEEEE&name3=20",
	"http://localhost:8888/sample?name1=AAAAA&name2=EEEEE&name3=30",
	"http://localhost:8888/sample?name1=BBBBB&name2=DDDDD&name3=0",
	"http://localhost:8888/sample?name1=BBBBB&name2=DDDDD&name3=10",
	"http://localhost:8888/sample?name1=BBBBB&name2=DDDDD&name3=20",
	"http://localhost:8888/sample?name1=BBBBB&name2=DDDDD&name3=30",
	"http://localhost:8888/sample?name1=BBBBB&name2=EEEEE&name3=0",
	"http://localhost:8888/sample?name1=BBBBB&name2=EEEEE&name3=10",
	"http://localhost:8888/sample?name1=BBBBB&name2=EEEEE&name3=20",
	"http://localhost:8888/sample?name1=BBBBB&name2=EEEEE&name3=30",
	"http://localhost:8888/sample?name1=CCCCC&name2=DDDDD&name3=0",
	"http://localhost:8888/sample?name1=CCCCC&name2=DDDDD&name3=10",
	"http://localhost:8888/sample?name1=CCCCC&name2=DDDDD&name3=20",
	"http://localhost:8888/sample?name1=CCCCC&name2=DDDDD&name3=30",
	"http://localhost:8888/sample?name1=CCCCC&name2=EEEEE&name3=0",
	"http://localhost:8888/sample?name1=CCCCC&name2=EEEEE&name3=10",
	"http://localhost:8888/sample?name1=CCCCC&name2=EEEEE&name3=20",
	"http://localhost:8888/sample?name1=CCCCC&name2=EEEEE&name3=30",
}

func TestParse(t *testing.T) {
	qpc, err := Parse("expand.yaml")
	if err != nil {
		t.Error("parse of yaml file failed")
	}

	if diff := cmp.Diff(qpc, expectedQpc); diff != "" {
		t.Errorf("QueryParametersConfig mismatch (-qpc +expectedQpc):\n%s", diff)
	}
}

func TestParseError(t *testing.T) {
	_, err := Parse("xxxxxx.yaml") //non-existent
	if err == nil {
		t.Error("no error returned with non-existent yaml file")
	} else {
		t.Log(err.Error())
	}
}

func TestCreate(t *testing.T) {
	urlArray, err := expectedQpc.Create()
	if err != nil {
		t.Error("Failed to create urlArray")
	}

	if diff := cmp.Diff(urlArray, expectedURLArray); diff != "" {
		t.Errorf("QueryParametersConfig mismatch (-urlArray +expectedURLArray):\n%s", diff)
	}
}

func TestCreateError(t *testing.T) {
	incorrectQpc := QueryParametersConfig{
		Host: "localhost:8888",
		Path: "/sample",
		QueryParameters: []QueryParameter{
			QueryParameter{
				Name:     "name1",
				Type:     "List", //Wrong Value
				NullAble: true,
				Values: []string{
					"AAAAA",
					"BBBBB",
					"CCCCC",
				},
			},
			QueryParameter{
				Name:     "name2",
				Type:     "list",
				NullAble: false,
				Values: []string{
					"DDDDD",
					"EEEEE",
				},
			},
		},
	}

	_, err := incorrectQpc.Create()
	if err == nil {
		t.Error("urlArray is generated with an invalid value")
	} else {
		t.Log(err.Error())
	}
}
