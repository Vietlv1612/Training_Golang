package plugin1

type Plugin1 struct{}

func (p Plugin1) Run(data string) string {
	return "Plugin1 processed: " + data
}
