1,数据叶子结点每条记录占用数据大小：
A，每个字段占用大小之和（bit占用一个字节，不会对齐）
B，隐藏的字段
6字节的row_id（如果需要）
6字节的transaction_id
7字节的roll_pointer
C，5字节的记录头

2,聚合索引每条记录大小：
A，联合索引每个字段的长度之和
B，4字节的（page number）
C，5字节的记录头

3，二级索引叶子结点每条记录大小：
A，索引的每个字段长度之和
B，聚集索引的每个字段长度之和（字段名字重复的算两次）
C，5字节的记录头

4，二级索引每条记录大小：
A，索引的每个字段长度之和
B，聚集索引的每个字段长度之和（字段名字重复的算两次）
C，4字节的（page number）
4，5字节的记录头

如何验证大小计算是否正确，采用工具innodb_space
1,innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge1 space-page-type-regions 查看索引占用的页号
2，innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge1 -p 3  page-illustrate 查看具体每一页的占用
  █ Record Header                        10    0.06%
  █ Record Data                          44    0.27%
其中Record Header 的字节数/5就是记录条数
Record Data/条数就是每条记录不包括头的大小+5就是每条记录的大小
