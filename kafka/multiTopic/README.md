1，针对同一个连接 只能调用一次Subscribe，第二次不会生效
2，如果希望在一个项目中分开消费多个topic ，需要初始化多个consumer实例，即多个连接

