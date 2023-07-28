package conf

type System struct {
	Add  string `yaml:"addr"`
	Port string `yaml:"port"`
	Web  string `yaml:"web"`
}

func (s System) Address() string {
	return s.Add + ":" + s.Port
}
