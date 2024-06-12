package model

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var CliCommandList map[string]CliCommand
