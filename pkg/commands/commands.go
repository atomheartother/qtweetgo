package commands

import (
	"github.com/atomheartother/qtweet/pkg/config"
)

// Command represents a parsed command, its command, arguments and options haven't been
// interpreted yet, just parsed to regognize multiple words surrounded by "" for example
// and to differentiate args from flags from options
type Command struct {
	cmd     string
	args    []string
	flags   []string
	options [][2]string
}

// Parse will parse a string and determine whether or not it is a recognizable command
// A return value of nil indicates it isn't a valid command
func Parse(s string) *Command {
	res := config.CommandRegexp.FindAllStringSubmatch(s, -1)
	if len(res) != 1 {
		return nil
	}
	parsedCommand := &Command{res[0][1], []string{}, []string{}, [][2]string{}}
	arguments := config.CommandArgsRegexp.FindAllStringSubmatch(res[0][2], -1)
	for _, match := range arguments {
		if match[6] != "" {
			// A simple argument
			parsedCommand.args = append(parsedCommand.args, match[6])
		} else if match[5] != "" {
			// Multi-word argument
			parsedCommand.args = append(parsedCommand.args, match[5])
		} else if match[1] != "" && match[3] == "" && match[4] == "" {
			// Flag
			parsedCommand.flags = append(parsedCommand.flags, match[1])
		} else if match[1] != "" && match[3] != "" {
			parsedCommand.options = append(parsedCommand.options, [2]string{match[1], match[3]})
		} else if match[1] != "" && match[4] != "" {
			parsedCommand.options = append(parsedCommand.options, [2]string{match[1], match[4]})
		}
	}
	return parsedCommand
}
