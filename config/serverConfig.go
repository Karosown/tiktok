package config

var Server *server

func getServer() {
	c := Config{}
	c = *InitConfig()
	Server = &c.Svc
}
