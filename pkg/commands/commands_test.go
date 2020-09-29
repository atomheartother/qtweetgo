package commands

import "testing"

type ParseTest struct {
	input  string
	output Command
	valid  bool
}

func TestParse(t *testing.T) {
	var tests = []ParseTest{
		{
			input: "command",
			output: Command{
				cmd:     "command",
				args:    []string{},
				flags:   []string{},
				options: [][2]string{},
			},
			valid: true,
		},
		{
			input: "cmd --flag arg1 arg2",
			output: Command{
				cmd:     "cmd",
				args:    []string{"arg1", "arg2"},
				flags:   []string{"flag"},
				options: [][2]string{},
			},
			valid: true,
		},
		{
			input: "mycommand --option=value \"multiple word argument\"",
			output: Command{
				cmd:     "mycommand",
				args:    []string{"multiple word argument"},
				flags:   []string{},
				options: [][2]string{{"option", "value"}},
			},
			valid: true,
		},
		{
			input: "",
			output: Command{
				cmd:     "",
				args:    []string{},
				flags:   []string{},
				options: [][2]string{},
			},
			valid: false,
		},
		{
			input: "z \"a good parser\" --can=\"parse anything??\"      --maybe   --maaaaybe",
			output: Command{
				cmd:     "z",
				args:    []string{"a good parser"},
				flags:   []string{"maybe", "maaaaybe"},
				options: [][2]string{{"can", "parse anything??"}},
			},
			valid: true,
		},
	}
	for _, test := range tests {
		res := Parse(test.input)
		// Check that it's valid
		if res == nil && !test.valid {
			// The test was supposed to be invalid and the function returned nil, all good
			continue
		}
		if res == nil && test.valid {
			t.Error("Expected valid output, got nil")
		}
		if res != nil && !test.valid {
			t.Error("Expected nil output, got", res)
		}
		// Check resulting command
		if res.cmd != test.output.cmd {
			t.Error("Expected command:", test.output.cmd, "got:", res.cmd)
		}
		// Check args, flags and options lengths
		if len(res.args) != len(test.output.args) {
			t.Error("Expected args:", test.output.args, "got:", res.args)
		}
		if len(res.flags) != len(test.output.flags) {
			t.Error("Expected flags:", test.output.flags, "got:", res.flags)
		}
		if len(res.options) != len(test.output.options) {
			t.Error("Expected options:", test.output.options, "got:", res.options)
		}

		// Check contents of arrays
		for idx, s := range res.args {
			if s != test.output.args[idx] {
				t.Error("Expected arg:", test.output.args[idx], "got:", s)
			}
		}
		for idx, s := range res.flags {
			if s != test.output.flags[idx] {
				t.Error("Expected flag:", test.output.flags[idx], "got:", s)
			}
		}
		for idx, opt := range res.options {
			if opt[0] != test.output.options[idx][0] || opt[1] != test.output.options[idx][1] {
				t.Error("Expected option:", test.output.options[idx], "got:", opt)
			}
		}
	}
}
