package path

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

const (
	IS_RELATIVE = "IsRelative"
	IS_DIR      = "IsDir"
	VALUE       = "Value"
	EXISTS      = "Exists"
)

func Marshal(v interface{}) ([]byte, error) {
	switch t := v.(type) {
	case string:
		return []byte(t), nil

	case []string:
		s := consts.PathSlash + strings.Join(t, consts.PathSlash)
		return []byte(s), nil

	case map[string]interface{}:
		name, err := types.ConvertGoType(t[VALUE], types.String)
		if err != nil {
			return nil, fmt.Errorf("unable to get '%s' from %v", VALUE, t)
		}
		return []byte(name.(string)), nil

	case []interface{}:
		if len(t) == 0 {
			return nil, nil
		}
		return marshalPathInterface(t)

	default:
		return nil, fmt.Errorf("%s can only marshal arrays. Instead got %T", types.Path, t)
	}
}

func marshalPathInterface(v []interface{}) ([]byte, error) {
	a := make([]string, len(v))

	for i := range v {
		switch v[i].(type) {
		case map[string]interface{}:
			name, err := types.ConvertGoType(v[i].(map[string]interface{})[VALUE], types.String)
			if err != nil {
				return nil, fmt.Errorf("unable to get '%s' from %v", VALUE, v[i])
			}
			a[i] = name.(string)

		default:
			name, err := types.ConvertGoType(v[i], types.String)
			if err != nil {
				return nil, err
			}
			a[i] = name.(string)
		}
	}

	s := strings.Join(a, consts.PathSlash)
	if f, _ := os.Stat(s); f != nil && f.IsDir() {
		s += consts.PathSlash
	}

	s = path.Clean(s)

	return []byte(s), nil
}

var pathSlashByte = consts.PathSlash[0]

func Unmarshal(b []byte) (interface{}, error) {
	if len(b) == 0 {
		b = []byte{'.'}
	}

	relative := b[0] != pathSlashByte
	path := string(b)

	f, err := os.Stat(path)
	dir := err == nil && f.IsDir()

	split := Split(path)

	notExists := make([]bool, len(split))
	for i := len(split) - 1; i > -1; i-- {
		notExists[i] = os.IsNotExist(err)
		if !notExists[i] {
			break
		}
		_, err = os.Stat(strings.Join(split[:i], consts.PathSlash))
	}

	v := make([]interface{}, len(split))

	for i := range split {
		v[i] = map[string]interface{}{
			IS_RELATIVE: relative && i == 0,
			IS_DIR:      (dir && i == len(split)-1) || i < len(split)-1,
			VALUE:       split[i],
			EXISTS:      !notExists[i],
		}
	}

	return v, nil
}
