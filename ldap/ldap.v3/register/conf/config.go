package conf

var Conf Configure = Configure{
	Ldap: LDAP{
		Host: "localhost",
		Port: "389",                        //"636","389"
		User: "cn=admin,dc=example,dc=org", //"admin","cn=admin,dc=example,dc=org"
		Pswd: "admin",
		Base: "org",
		Dc:   "example",
	},
}

type LDAP struct {
	Host string
	Port string
	User string
	Pswd string
	Base string
	Dc   string
}

type Configure struct {
	Ldap LDAP
}
