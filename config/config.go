package config

type Mysql struct {
	Host     string
	Port     int64
	User     string
	Password string
	Database string
}

type Redis struct {
	Host     string
	Port     int64
	Password string
}

type Elasticsearch struct {
	Host string
	Port int64
}

type MongoDB struct {
	Host     string
	Port     int64
	User     string
	Password string
}

type NaCos struct {
	NamespaceId string
	Host        string
	Port        int64
	DataId      string
	Group       string
}

type AppConfig struct {
	Mysql         `json:"mysql"`
	Redis         `json:"redis"`
	Elasticsearch `json:"elasticsearch"`
	MongoDB       `json:"mongoDB"`
	NaCos         `json:"naCos"`
}
