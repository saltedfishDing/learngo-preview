package conf

type Database struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Dbname       string `yaml:"db-name"`
	Config       string `yaml:"config"`         // 高级配置
	MaxIdleConns int    `yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `yaml:"log-mode"`       // 是否开启Gorm全局日志
}

func (m *Database) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
