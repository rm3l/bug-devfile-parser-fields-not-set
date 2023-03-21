package devfile

import (
	"github.com/devfile/library/v2/pkg/devfile/parser"
)

func Parse(path string, flatten *bool) (parser.DevfileObj, error) {
	d, err := parser.ParseDevfile(parser.ParserArgs{Path: path, FlattenedDevfile: flatten})
	if err != nil {
		return parser.DevfileObj{}, err
	}
	return d, nil
}
