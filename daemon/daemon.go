package main

var config *Config

//TODO: Write metrics daemon.

func main()  {

	config = getConfig()

	startServer(config.System.Host, config.System.Port)

}