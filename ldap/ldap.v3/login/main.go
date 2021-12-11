package main

import (
	"fmt"
	"login/conf"

	ldap "github.com/go-ldap/ldap/v3"
)

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	params := Params{
		Username: "admin",
		Password: "admin",
	}

	// conn, err := ldap.DialTLS("tcp", conf.Conf.Ldap.Host+":"+conf.Conf.Ldap.Port, &tls.Config{
	// 	InsecureSkipVerify: true,
	// })
	conn, err := ldap.Dial("tcp", conf.Conf.Ldap.Host+":"+conf.Conf.Ldap.Port)
	if err != nil {
		fmt.Println("Ldap server disconnect.", err)
		return
	}
	defer conn.Close()

	err = conn.Bind(fmt.Sprintf("cn=%s,dc=%s,dc=%s", params.Username, conf.Conf.Ldap.Dc, conf.Conf.Ldap.Base), params.Password)
	if err != nil {
		fmt.Println("Password error.", err)
		return
	}
	getUserInfo(conn, params, conf.Conf)
	getUserInfo(conn, Params{
		Username: "xia zemin",
		Password: "xiazemin",
	}, conf.Conf1)
}

func getUserInfo(conn *ldap.Conn, params Params, config conf.Configure) {
	//cn=xia zemin,ou=People,dc=example,dc=org
	cn := fmt.Sprintf("dc=%s,dc=%s", config.Ldap.Dc, config.Ldap.Base)
	if config.Ldap.Ou != "" {
		cn = "ou=" + config.Ldap.Ou + "," + cn
	}
	fmt.Println(cn)
	sql := ldap.NewSearchRequest(cn, //conf.Conf.Ldap.Base
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		0,
		false,
		fmt.Sprintf("(cn=%s)", params.Username),
		[]string{"cn", "displayName", "mail", "mobile", "employeeID", "givenName"},
		nil)

	var cur *ldap.SearchResult

	cur, err := conn.Search(sql)
	if err != nil {
		fmt.Println("Ldap server search failed.", err)
		return
	}

	if len(cur.Entries) == 0 {
		fmt.Println("Not found user.", nil)
		return
	}
	fmt.Printf("%#v,arrt:%#v", cur.Entries[0], cur.Entries[0].Attributes)

	for _, at := range cur.Entries[0].Attributes {
		fmt.Printf("\t%#v\n", at)
	}

	var result = struct {
		Name       string `json:"name"`
		Account    string `json:"account"`
		Email      string `json:"email"`
		Phone      string `json:"phone"`
		EmployeeId string `json:"employeeId"`
	}{
		Name:       cur.Entries[0].GetAttributeValue("givenName"),
		Account:    cur.Entries[0].GetAttributeValue("sAMAccountName"),
		Email:      cur.Entries[0].GetAttributeValue("mail"),
		Phone:      cur.Entries[0].GetAttributeValue("mobile"),
		EmployeeId: cur.Entries[0].GetAttributeValue("employeeID"),
	}
	fmt.Println(result)

}
