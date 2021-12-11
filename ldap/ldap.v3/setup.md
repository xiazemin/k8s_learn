 % cd register 
 % go mod init reg
 % go mod tidy

 go: reg imports
        gopkg.in/ldap.v3: gopkg.in/ldap.v3@v3.1.3: parsing go.mod:
        module declares its path as: github.com/go-ldap/ldap/v3
                but was required as: gopkg.in/ldap.v3


"gopkg.in/ldap.v3" 
	github.com/go-ldap/ldap/v3


% go run main.go
Ldap server disconnect. LDAP Result Code 200 "Network Error": EOF
% go run main.go
Ldap server unbind. LDAP Result Code 34 "Invalid DN Syntax": invalid DN
% go run main.go
User insert error. LDAP Result Code 34 "Invalid DN Syntax": invalid DN
 % go run main.go
user DN: cn=xiazemin,dc=example,dc=org
User insert error. LDAP Result Code 17 "Undefined Attribute Type": sAMAccountName: attribute type undefined

UserPrincipalName：

用户登录名格式：xiaowen@azureyun.com
是基于Internet标准RFC 822的用户Internet样式登录名;
在目录林中的所有安全主体对象中应该是唯一的;
UPN是可选的，在创建用户帐户时可指定也可不单独指定；
SamAccountName：

与早期版本的Windows（pre-windows 2000）一起使用;
用户登录名格式：azureyun\xiaowen
不能超过20个字符；
在域中的所有安全主体对象中是唯一的；


SamAccountName：只有windows支持需要换成 UserPrincipalName：
https://www.cnblogs.com/wenzhongxiang/p/10416122.html
% go run main.go
user DN: cn=xiazemin,dc=example,dc=org
User insert error. LDAP Result Code 17 "Undefined Attribute Type": UserPrincipalName: attribute type undefined


% go mod init login
% go mod tidy
 % go run main.go 
Ldap server search failed. LDAP Result Code 34 "Invalid DN Syntax": invalid DN

% go mod init reset
% go mod tidy
% go run main.go 
Not found user. <nil>

能通过用户名和密码找到记录说明认证成功