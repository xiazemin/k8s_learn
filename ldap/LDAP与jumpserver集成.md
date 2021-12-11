1.首先我们以管理用户登陆到jumpserver中，点击系统设置，知道LDAP设置



LDAP地址：填LDAP的地址

绑定DN：输入管理员条目的DN

密码：输入管理员的密码

用户OU：输入你的创建的分组OU，他就只会去匹配你分组底下的用户，多个OU用I分割

用户过滤器：就是匹配分组的账户下的某个字段可以有uid sn cn都可以

LDAP属性映射：将ldap分组用户下的某个属性映射jumpserver的某个属性

填好上述内容后，点击test测试一下看可以匹配到用户，如果测试匹配到用户，就可以点击提交，然后用ldap的用户去登陆jumpserver,登陆成功后，管理员用户可以看到ldap的用户加入到用户设置中了。

五.LDAP与jenkins集成

1.首先我们以管理员用户登陆上jenkins，找到设置ldap的地方。



2.启用ldap



3.编辑配置，并且点击测试，测试成功即可保存，登出jenkins然后用授权的账号进行登陆。

https://www.kancloud.cn/linux_jia/haha/1393994