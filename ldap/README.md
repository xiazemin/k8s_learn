问题：

1. 我们日常的办公系统是不是有多个？

2. 每个系统之间是不是都有独立的账号密码？

3. 密码多了，有时候半天想不起来哪个密码对应哪个系统？

4. 每次新项目的开发，都需要重新开发和维护一套用户密码？

5. 维护多套系统的用户是不是非常头疼？

LDAP统一认证服务”已经帮助大家解决这些问题了


LDAP是开放的Internet标准，支持跨平台的Internet协议，在业界中得到广泛认可的，并且市场上或者开源社区上的大多产品都加入了对LDAP的支持，因此对于这类系统，不需单独定制，只需要通过LDAP做简单的配置就可以与服务器做认证交互。“简单粗暴”，可以大大降低重复开发和对接的成本。

我们拿开源系统（YAPI）做案例，只需做一下简单的几步配置就可以达到LDAP的单点登录认证了：

{
"ldapLogin": {
      "enable": true,
      "server": "ldap://l-ldapt1.ops.dev.cn0.qunar.com",
      "baseDn": "CN=Admin,CN=Users,DC=test,DC=com",
      "bindPassword": "password123",
      "searchDn": "OU=UserContainer,DC=test,DC=com",
      "searchStandard": "mail"
   }
}

LDAP仅仅是一个访问协议，那么我们的数据究竟存储在哪里呢？
OpenLDAP 开源的项目，速度很快，但是非主 流应用。

一）目录树概念

1. 目录树：在一个目录服务系统中，整个目录信息集可以表示为一个目录信息树，树中的每个节点是一个条目。

2. 条目：每个条目就是一条记录，每个条目有自己的唯一可区别的名称（DN）。

3. 对象类：与某个实体类型对应的一组属性，对象类是可以继承的，这样父类的必须属性也会被继承下来。

4. 属性：描述条目的某个方面的信息，一个属性由一个属性类型和一个或多个属性值组成，属性有必须属性和非必须属性。

（二）DC、UID、OU、CN、SN、DN、RDN

关键字

英文全称

含义

dc

Domain Component

域名的部分，其格式是将完整的域名分成几部分，如域名为example.com变成dc=example,dc=com（一条记录的所属位置）

uid

User Id

用户ID songtao.xu（一条记录的ID）

ou

Organization Unit

组织单位，组织单位可以包含其他各种对象（包括其他组织单元），如“oa组”（一条记录的所属组织）

cn

Common Name

公共名称，如“Thomas Johansson”（一条记录的名称）

sn

Surname

姓，如“许”

dn

Distinguished Name

“uid=songtao.xu,ou=oa组,dc=example,dc=com”，一条记录的位置（唯一）

rdn

Relative dn

相对辨别名，类似于文件系统中的相对路径，它是与目录树结构无关的部分，如“uid=tom”或“cn= Thomas Johansson”

统一身份认证主要是改变原有的认证策略，使需要认证的软件都通过LDAP进行认证，在统一身份认证之后，用户的所有信息都存储在AD Server中。终端用户在需要使用公司内部服务的时候，都需要通过AD服务器的认证。


 我们以PHP脚本作为例子：

$ldapconn = ldap_connect(“10.1.8.78")
$ldapbind = ldap_bind($ldapconn, 'username', $ldappass);
$searchRows= ldap_search($ldapconn, $basedn, "(cn=*)");
$searchResult = ldap_get_entries($ldapconn, $searchRows);
ldap_close($ldapconn);
1. 连接到LDAP服务器；

2. 绑定到LDAP服务器；

3. 在LDAP服务器上执行所需的任何操作；

4. 释放LDAP服务器的连接；

https://www.cnblogs.com/wilburxu/p/9174353.html

dn:cn=honglv,ou=bei,ou=xi,ou=dong,dc=ljheee
其中dn标识一条记录，描述了一条数据的详细路径。

为什么ou会有多个值？你想想，从树根到达苹果的位置，可能要经过好几个树杈，所有ou可能有多个值。关于dn后面一长串，分别是cn，ou,dc；中间用逗号隔开。


JNDI是 Java 命名与目录接口（Java Naming and Directory Interface），在J2EE规范中是重要的规范之一

https://blog.csdn.net/wn084/article/details/80729230

JNDI连接LDAP服务器
https://www.jianshu.com/p/7e4d99f6baaf


