package commands

type Runner interface {
	Run() (string, error)
}

type Command struct {
	Runner
}

func (c Command) Exec() (string, error) {
	return c.Run()
}
