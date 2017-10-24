package system

func Open(name string) error {
	cmd := open(name)
	return cmd.Run()
}
