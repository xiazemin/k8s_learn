http://www.ldap.org.cn/

账号的统一管理。
1：OpenLDAP服务端的搭建
2：PhpLDAPAdmin的搭建
3：OpenLDAP的打开日志信息
4：OpenLDAP与migrationtools实现导入系统账号的相关信息
5：OpenLDAP客户端的配置
6：OpenLDAP与SSH
7：OpenLDAP限制用户登录系统
8：OpenLDAP强制用户一登录系统更改密码
9：OpenLDAP与系统账号结合Samba
10：OpenLDAP的主从
11：OpenLDAP的双主

1）安装OpenLDAP的相关
         yum -y install openldap openldap-servers openldap-clients openldap-devel compat-openldap   其中compat-openldap这个包与主从有很大的关系

2）OpenLDAP的相关配置文件信息
      /etc/openldap/slapd.conf：OpenLDAP的主配置文件，记录根域信息，管理员名称，密码，日志，权限等
      /etc/openldap/slapd.d/*：这下面是/etc/openldap/slapd.conf配置信息生成的文件，每修改一次配置信息，这里的东西就要重新生成
      /etc/openldap/schema/*：OpenLDAP的schema存放的地方
      /var/lib/ldap/*：OpenLDAP的数据文件
      /usr/share/openldap-servers/slapd.conf.obsolete 模板配置文件
      /usr/share/openldap-servers/DB_CONFIG.example 模板数据库配置文件

一步搭建OpenLDAP服务端的时候，并没有把管理员的账号信息导入，编辑root.ldif，然后导入
 dn: dc=lemon,dc=com
 objectclass: dcObject
 objectclass: organization
 o: Yunzhi,Inc.
 dc: lemon
 
 dn: cn=Captain,dc=lemon,dc=com
 objectclass: organizationalRole
 cn: Captain
 这里得注意每一个属性： 后必须有空格，但是值的后面不能有任何空格
 然后导入：ldapadd -x -D "cn=Captain,dc=lemon,dc=com" -W -f root.ldif


 https://www.cnblogs.com/lemon-le/p/6266921.html
 https://www.cnblogs.com/ccbloom/p/14100244.html
 https://blog.csdn.net/weixin_41004350/article/details/89521170


https://wiki.archlinux.org/title/OpenLDAP_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87)

 Docker-OpenLDAP
GitHub: osixia/docker-openldap
GitHub:osixia/docker-phpLDAPadmin

https://github.com/osixia/docker-openldap

https://github.com/osixia/docker-phpLDAPadmin

https://www.jianshu.com/p/c079e1508184

https://hub.docker.com/r/osixia/openldap/

By default the admin has the password admin


docker run \
	--env LDAP_ORGANISATION="My Company" \
	--env LDAP_DOMAIN="my-company.com" \
	--env LDAP_ADMIN_PASSWORD="JonSn0w" \
	--detach osixia/openldap:1.5.0

The directories /var/lib/ldap (LDAP database files) and /etc/ldap/slapd.d (LDAP config files) are used to persist the schema and data information, 

This image can load ldif files at startup with either ldapadd or ldapmodify. Mount .ldif in /container/service/slapd/assets/config/bootstrap/ldif directory if you want to overwrite image default bootstrap ldif files or in /container/service/slapd/assets/config/bootstrap/ldif/custom (recommended) to extend image config.



https://hub.docker.com/r/osixia/phpldapadmin/

docker pull osixia/phpldapadmin:latest


docker run -p 6443:443 \
        --env PHPLDAPADMIN_LDAP_HOSTS=ldap.example.com \
        --detach osixia/phpldapadmin:0.9.0




#!/bin/bash -e
docker run --name ldap-service --hostname ldap-service --detach osixia/openldap:1.1.8
docker run --name phpldapadmin-service --hostname phpldapadmin-service --link ldap-service:ldap-host --env PHPLDAPADMIN_LDAP_HOSTS=ldap-host --detach osixia/phpldapadmin:0.9.0

PHPLDAP_IP=$(docker inspect -f "{{ .NetworkSettings.IPAddress }}" phpldapadmin-service)

echo "Go to: https://$PHPLDAP_IP"
echo "Login DN: cn=admin,dc=example,dc=org"
echo "Password: admin"



PHPLDAPADMIN_LDAP_HOSTS: Set phpLDAPadmin server config. Defaults to :

- ldap.example.org:
  - server:
    - tls: true
  - login:
    - bind_id: cn=admin,dc=example,dc=org
- ldap2.example.org
- ldap3.example.org




[root@ldaptest openldap]# ldapadd -x -D "cn=admin,dc=ultrapower,dc=com" -W -f /tmp/base.ldif
Enter LDAP Password:
ldap_sasl_bind(SIMPLE): Can't contact LDAP server (-1)

https://www.cnblogs.com/rusking/p/8035212.html


docker exec $LDAP_CID ldapadd -x -D "cn=admin,dc=example,dc=org" -w admin -f /container/service/slapd/assets/test/new-user.ldif -H ldap://ldap.example.org -ZZ


https://blog.51cto.com/u_9099998/2426546

https://blog.csdn.net/weixin_40352715/article/details/106826239
