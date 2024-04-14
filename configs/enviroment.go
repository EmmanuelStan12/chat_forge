package config

type Environment struct{}

const (
	DEV        string = "dev"
	STAGING    string = "staging"
	PRODUCTION string = "prod"
)

func (e Environment) isDev() {

}

func (e Environment) isProd() {

}

func (e Environment) isStaging() {

}
