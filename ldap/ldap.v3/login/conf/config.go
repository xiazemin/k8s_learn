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

var Conf1 Configure = Configure{
	Ldap: LDAP{
		Host: "localhost",
		Port: "389", //"636","389"
		Pswd: "xiazemin",
		Base: "org",
		Dc:   "example",
		Ou:   "People",
		Cn:   "xia zemin",
		User: "xiazemin",
	},
}

type LDAP struct {
	Host string
	Port string
	User string
	Pswd string
	Base string
	Dc   string
	//cn=xia zemin,ou=People,dc=example,dc=org
	Cn string
	Ou string
}

type Configure struct {
	Ldap LDAP
}
