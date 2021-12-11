go get gopkg.in/ldap.v3

https://github.com/go-ldap/ldap/tree/v3.4.1

https://blog.csdn.net/weixin_35396246/article/details/112630502

https://blog.csdn.net/wyansai/article/details/99416987


在项目中需要获取的ldap数据字段包括：uid(用户id) 、deportment(部门)，displayName(姓名)、mail(邮箱)

LdapUsernameKey   = "uid"         //对应config.Ldap().Attributes.UNameKey
LdapNameKey       = "displayName" //对应config.Ldap().Attributes.NameKey
LdapEmailKey      = "mail"        //对应config.Ldap().Attributes.EmailKey
LdapDepartmentKey = "department"  //对应config.Ldap().Attributes.DepartmentKey

https://zhuanlan.zhihu.com/p/144315778

git clone https://github.com/go-asn1-ber/asn1-ber.git
git clone https://github.com/go-ldap/ldap.git

https://studygolang.com/articles/21479?fr=sidebar

ldap server 使用了openldap，基于docker 运行

https://www.cnblogs.com/rongfengliang/p/13659051.html


github.com/jtblin/go-ldap-client

http://www.imooc.com/wenda/detail/657174