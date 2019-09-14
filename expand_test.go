package expand

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	qpc, err := Parse("expand.yaml")
	if err != nil {
		t.Error(err.Error())
	}

	expectedQpc := QueryParametersConfig{
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

	if diff := cmp.Diff(qpc, expectedQpc); diff != "" {
		t.Errorf("Tasks mismatch (-tasks +expectedTasks):\n%s", diff)
	}
}
