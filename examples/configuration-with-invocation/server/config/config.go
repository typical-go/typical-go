package config

// Config of app
type Config struct {
	Address string `default:":8080" required:"true"`
}
