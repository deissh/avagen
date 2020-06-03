package plugins

import "github.com/pkg/errors"

// ValidateArgs and setup default values if not present in args
func ValidateArgs(plugin Plugin, args ParsedArg) (ParsedArg, error) {
	for _, v := range plugin.Args {
		_, ok := args[v.Key]

		if !ok && v.Required {
			return nil, errors.Errorf("Key %s is required but not founded in request args", v.Key,)
		}
		if !ok {
			args[v.Key] = v.Default
		}
	}

	return args, nil
}