package argparse

type subCommandError struct {
	error
	cmd *Command
}

func (e subCommandError) Error() string {
	return "[sub]Command required\n"
}

func newSubCommandError(cmd *Command) error {
	return subCommandError{cmd: cmd}
}
