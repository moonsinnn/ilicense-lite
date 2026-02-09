package config

type MysqlDataSource struct {
	DBDriver        string                 `yaml:"db_driver"`
	DBName          string                 `yaml:"db_name"`
	UserName        string                 `yaml:"user_name"`
	Password        string                 `yaml:"password"`
	Address         string                 `yaml:"address"`
	ConnTimeout     int64                  `yaml:"conn_timeout"`
	ReadTimeout     int64                  `yaml:"read_timeout"`
	WriteTimeout    int64                  `yaml:"write_timeout"`
	MaxOpenConn     int                    `yaml:"max_open_conn"`
	MaxIdleConn     int                    `yaml:"max_idle_conn"`
	ConnMaxLifeTime int64                  `yaml:"conn_max_life_time"`
	Params          []MysqlDataSourceParam `yaml:"params"`
}

type MysqlDataSourceParam struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type MysqlConfig struct {
	DataSource       MysqlDataSource `yaml:"data_source"`
	LogFile          string          `json:"log_file"`
	LogSlowThreshold int64           `json:"log_slow_threshold"`
	Colorful         bool            `json:"colorful"`
}
