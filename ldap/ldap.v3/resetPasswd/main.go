package main

import (
	"fmt"
	"reset/conf"

	ldap "github.com/go-ldap/ldap/v3"
	"golang.org/x/text/encoding/unicode"
)

func main() {
	params := struct {
		Username    string `json:"username"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}{
		Username:    "admin",
		OldPassword: "admin",
		NewPassword: "admin",
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

	err = conn.Bind(fmt.Sprintf("cn=%s,dc=%s,dc=%s", params.Username, conf.Conf.Ldap.Dc, conf.Conf.Ldap.Base), params.OldPassword)
	if err != nil {
		fmt.Println("Password error.", err)
		return
	}
	sql := ldap.NewSearchRequest(fmt.Sprintf("dc=%s,dc=%s", conf.Conf.Ldap.Dc, conf.Conf.Ldap.Base), //conf.Conf.Ldap.Base
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		0,
		false,
		fmt.Sprintf("(sAMAccountName=%s)", params.Username),
		[]string{"sAMAccountName", "displayName", "mail", "mobile", "employeeID"},
		nil)

	var cur *ldap.SearchResult

	if cur, err = conn.Search(sql); err != nil {
		fmt.Println("Ldap server search failed.", err)
		return
	}

	if len(cur.Entries) == 0 {
		fmt.Println("Not found user.", nil)
		return
	}

	err = conn.Bind(conf.Conf.Ldap.User, conf.Conf.Ldap.Pswd)
	if err != nil {
		fmt.Println("Ldap server unbind.", err)
		return
	}

	var userDn = fmt.Sprintf("CN=%s,dc=%s,dc=%s", params.Username, conf.Conf.Ldap.Dc, conf.Conf.Ldap.Base)

	sql2 := ldap.NewModifyRequest(userDn, nil)

	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, _ := utf16.NewEncoder().String("\"" + params.NewPassword + "\"")

	sql2.Replace("unicodePwd", []string{pwdEncoded})
	sql2.Replace("userAccountControl", []string{"512"})

	if err := conn.Modify(sql2); err != nil {
		fmt.Println("Ldap password modify failed.", err)
		return
	}
	fmt.Println("sucess")
}
