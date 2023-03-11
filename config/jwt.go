package config

type jwt struct {
	Secret string
	Issuer string
	Ttl    int
}

var Jwt = jwt{Secret: "12321114rdst43rqsa", Issuer: "test", Ttl: 1100000}
