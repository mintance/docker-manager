package main

var config *Config

//TODO: Write metrics daemon.

func main()  {

	// Get all main config data.
	config = getConfig()

	// Let's see if we have some configurations.

	// Let's sync active containers with our plans.

	// Start monitoring containers.

	// We need to listen client sometimes.
	startClientServer(config.Client.Host, config.Client.Port)

}