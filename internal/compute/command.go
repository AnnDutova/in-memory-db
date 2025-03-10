package compute

const (
	GetCommand = "GET"
	SetCommand = "SET"
	DelCommand = "DEL"
)

var cmdToArgsCount = map[Command]int{
	GetCommand: 1,
	SetCommand: 2,
	DelCommand: 1,
}

func (c *Command) validate(in string) error {
	switch in {
	case GetCommand, SetCommand, DelCommand:
		return nil
	default:
		return ErrInvalidCommand
	}
}

func (c *Arguments) validate(cmd Command, args []string) error {
	count, ok := cmdToArgsCount[cmd]
	if !ok {
		return ErrInvalidCommand
	}

	if count != len(args) {
		return ErrInvalidArgumentsCount
	}

	return nil
}
