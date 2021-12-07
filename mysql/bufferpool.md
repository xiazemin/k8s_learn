mysql> set global innodb_buffer_pool_size=0;
Query OK, 0 rows affected, 2 warnings (0.00 sec)

 show variables like "%innodb_buffer_pool_instances%";
+------------------------------+-------+
| Variable_name                | Value |
+------------------------------+-------+
| innodb_buffer_pool_instances | 1     |
+------------------------------+-------+
1 row in set (0.00 sec)


> show variables like '%innodb_buffer_pool_size%';
+-------------------------+-----------+
| Variable_name           | Value     |
+-------------------------+-----------+
| innodb_buffer_pool_size | 134217728 |
+-------------------------+-----------+
1 row in set (0.00 sec)



 show variables like "%pool%";
+-------------------------------------+----------------+
| Variable_name                       | Value          |
+-------------------------------------+----------------+
| innodb_buffer_pool_chunk_size       | 134217728      |
| innodb_buffer_pool_dump_at_shutdown | ON             |
| innodb_buffer_pool_dump_now         | OFF            |
| innodb_buffer_pool_dump_pct         | 25             |
| innodb_buffer_pool_filename         | ib_buffer_pool |
| innodb_buffer_pool_in_core_file     | ON             |
| innodb_buffer_pool_instances        | 1              |
| innodb_buffer_pool_load_abort       | OFF            |
| innodb_buffer_pool_load_at_startup  | ON             |
| innodb_buffer_pool_load_now         | OFF            |
| innodb_buffer_pool_size             | 134217728      |
+-------------------------------------+----------------+
11 rows in set (0.03 sec)


 set global innodb_buffer_pool_chunk_size=0;
ERROR 1238 (HY000): Variable 'innodb_buffer_pool_chunk_size' is a read only variable

https://blog.csdn.net/nanyanglu/article/details/79109838

https://www.cnblogs.com/yybrhr/p/11445748.html

set global innodb_buffer_pool_instances=0;
ERROR 1238 (HY000): Variable 'innodb_buffer_pool_instances' is a read only variable

