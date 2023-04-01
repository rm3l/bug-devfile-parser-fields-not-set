package devfile

import (
	"github.com/devfile/library/v2/pkg/devfile/parser"
	"k8s.io/utils/pointer"
)

func Parse(path string, flatten *bool) (parser.DevfileObj, error) {
	d, err := parser.ParseDevfile(parser.ParserArgs{
		Path:               path,
		FlattenedDevfile:   flatten,
		SetBooleanDefaults: pointer.Bool(false),
	})
	if err != nil {
		return parser.DevfileObj{}, err
	}
	return d, nil
}
