package main

import (
	"fmt"
	"reg/conf"
	"strings"

	ldap "github.com/go-ldap/ldap/v3"
	"golang.org/x/text/encoding/unicode"
)

func main() {
	params := struct {
		Name       string `json:"name"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
		EmployeeID string `json:"employeeId"`
		Password   string `json:"password"`
	}{
		Name:       "xiazemin",
		Phone:      "12345",
		Email:      "xiazemin@exapmle.org",
		EmployeeID: "xiazemin_id",
		Password:   "xiazemin",
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

	err = conn.Bind(conf.Conf.Ldap.User, conf.Conf.Ldap.Pswd)
	if err != nil {
		fmt.Println("Ldap server unbind.", err)
		return
	}

	// 获取用户名，为邮箱前缀
	var username = strings.Split(params.Email, "@")[0]

	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, _ := utf16.NewEncoder().String("\"" + params.Password + "\"")

	var userDn = fmt.Sprintf("cn=%s,dc=%s,dc=%s", username, conf.Conf.Ldap.Dc, conf.Conf.Ldap.Base)
	fmt.Println("user DN:", userDn)
	// 添加账号
	sqlInsert := ldap.NewAddRequest(userDn, nil)
	sqlInsert.Attribute("cn", []string{username})
	sqlInsert.Attribute("sAMAccountName", []string{username}) //UserPrincipalName sAMAccountName
	sqlInsert.Attribute("userAccountControl", []string{"512"})
	sqlInsert.Attribute("unicodePwd", []string{pwdEncoded})
	sqlInsert.Attribute("displayName", []string{username})
	sqlInsert.Attribute("mobile", []string{params.Phone})
	sqlInsert.Attribute("employeeID", []string{params.EmployeeID})
	sqlInsert.Attribute("mail", []string{params.Email})
	sqlInsert.Attribute("givenName", []string{params.Name})
	sqlInsert.Attribute("userPrincipalName", []string{params.Email})
	sqlInsert.Attribute("objectClass", []string{"top", "person", "organizationalPerson", "user"})

	if err = conn.Add(sqlInsert); err != nil {
		if ldap.IsErrorWithCode(err, 68) {
			fmt.Println("User already exist.", err)
		} else {
			fmt.Println("User insert error.", err)
		}
		return
	}
	fmt.Println("sucess")
}
