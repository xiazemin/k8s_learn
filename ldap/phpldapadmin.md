https://www.cnblogs.com/xiaoqshuo/p/10132588.html

添加用户和组
4.1 创建Groups和People OU
People OU

4.2 创建组和用户
4.2.1 组 dev devops test
4.2.2 用户
填写基本信息，选择组和Login Shell
注意修改Common Name

cn=xia zemin,ou=People,dc=example,dc=org
First name	xia	
Last name	zemin	
Common Name	xia zemin	
User ID	xzemin	
Password	****************	 //xiazemin
UID Number	1000	
GID Number	500	
Home directory	/home/users/xzemin	
Login shell	/bin/bash	
objectClass	inetOrgPerson
posixAccount

4.3 为每个用户添加Email，没有Email无法登陆gitlab
