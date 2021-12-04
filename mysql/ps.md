%  csrutil status
System Integrity Protection status: enabled.

 % csrutil disable
csrutil: This tool needs to be executed from Recovery OS.


  % sudo dtruss -ap 709
: probe description syscall:::entry does not match any probes. System Integrity Protection is on


mysql -uroot -h127.0.0.1 -e"show variables like '%innodb_purge_threads%';\G"
mysql -uroot -h127.0.0.1 -e"show variables like '%innodb_page_cleaners%';\G"
mysql -uroot -h127.0.0.1 -e"show variables like '%innodb_undo%';\G"
ps -AvM 709
ps aux |grep mysql
lsof -i tcp:3306
cat /usr/local/var/mysql/bogon.pid
ls -F /proc/30580/task
ps -L aux | grep -e PID -e `pidof mysqld` | more
ps -L aux |grep -e PID -e `pidof mysqld` |more
ps -L -p `pidof mysqld` | more
lsof -i tcp:3306
ps -L -p 709 | more
ps -AvM 709 |grep 709 |wc -l
47 个线程比mysql 里看到的多

ps -T -p 709
dtruss -ap 709


 % ps -AvM 1733 |grep 1733 |wc -l
      38
  
  活动监视器看到的也是38个


启动一个链接后变成了39
 % ps -AvM 1733 |grep 1733 |wc -l
      39

退出后并没有减少，因为被缓存了



 select thread_id, thread_os_id, name from performance_schema.threads ;
+-----------+--------------+---------------------------------------------+
| thread_id | thread_os_id | name                                        |
+-----------+--------------+---------------------------------------------+
|         1 |        14370 | thread/sql/main                             |
|         2 |        15698 | thread/mysys/thread_timer_notifier          |
|         4 |        15820 | thread/innodb/io_ibuf_thread                |
|         5 |        15821 | thread/innodb/io_log_thread                 |
|         6 |        15822 | thread/innodb/io_read_thread                |
|         7 |        15823 | thread/innodb/io_read_thread                |
|         8 |        15824 | thread/innodb/io_read_thread                |
|         9 |        15825 | thread/innodb/io_read_thread                |
|        10 |        15826 | thread/innodb/io_write_thread               |
|        11 |        15827 | thread/innodb/io_write_thread               |
|        12 |        15828 | thread/innodb/io_write_thread               |
|        13 |        15829 | thread/innodb/io_write_thread               |
|        14 |        15830 | thread/innodb/page_flush_coordinator_thread |
|        16 |        15832 | thread/innodb/log_checkpointer_thread       |
|        17 |        15833 | thread/innodb/log_flush_notifier_thread     |
|        18 |        15834 | thread/innodb/log_flusher_thread            |
|        19 |        15835 | thread/innodb/log_write_notifier_thread     |
|        20 |        15836 | thread/innodb/log_writer_thread             |
|        25 |        16102 | thread/innodb/srv_lock_timeout_thread       |
|        26 |        16103 | thread/innodb/srv_error_monitor_thread      |
|        27 |        16104 | thread/innodb/srv_monitor_thread            |
|        28 |        16236 | thread/innodb/buf_resize_thread             |
|        29 |        16237 | thread/innodb/srv_master_thread             |
|        30 |        16238 | thread/innodb/dict_stats_thread             |
|        31 |        16239 | thread/innodb/fts_optimize_thread           |
|        32 |        16241 | thread/mysqlx/worker                        |
|        33 |        16242 | thread/mysqlx/worker                        |
|        34 |        16243 | thread/mysqlx/acceptor_network              |
|        38 |        16247 | thread/innodb/buf_dump_thread               |
|        39 |        16248 | thread/innodb/clone_gtid_thread             |
|        40 |        16249 | thread/innodb/srv_purge_thread              |
|        41 |        16250 | thread/innodb/srv_worker_thread             |
|        42 |        16249 | thread/innodb/srv_purge_thread              |
|        43 |        16250 | thread/innodb/srv_worker_thread             |
|        44 |        16251 | thread/innodb/srv_worker_thread             |
|        45 |        16252 | thread/innodb/srv_worker_thread             |
|        46 |        16251 | thread/innodb/srv_worker_thread             |
|        47 |        16252 | thread/innodb/srv_worker_thread             |
|        48 |        16253 | thread/sql/event_scheduler                  |
|        49 |        16254 | thread/sql/signal_handler                   |
|        50 |        16255 | thread/mysqlx/acceptor_network              |
|        52 |        16257 | thread/sql/compress_gtid_table              |
|        54 |        48638 | thread/sql/one_connection                   |
+-----------+--------------+---------------------------------------------+
43 rows in set (0.03 sec)


select distinct PROCESSLIST_STATE  from performance_schema.threads\G
*************************** 1. row ***************************
PROCESSLIST_STATE: NULL
*************************** 2. row ***************************
PROCESSLIST_STATE: waiting for handler commit
*************************** 3. row ***************************
PROCESSLIST_STATE: Waiting on empty queue
*************************** 4. row ***************************
PROCESSLIST_STATE: Suspending
*************************** 5. row ***************************
PROCESSLIST_STATE: executing
5 rows in set (0.01 sec)



 select PROCESSLIST_STATE,count(*) from performance_schema.threads group by(PROCESSLIST_STATE) \G
*************************** 1. row ***************************
PROCESSLIST_STATE: NULL
         count(*): 39
*************************** 2. row ***************************
PROCESSLIST_STATE: waiting for handler commit
         count(*): 1
*************************** 3. row ***************************
PROCESSLIST_STATE: Waiting on empty queue
         count(*): 1
*************************** 4. row ***************************
PROCESSLIST_STATE: Suspending
         count(*): 1
*************************** 5. row ***************************
PROCESSLIST_STATE: executing
         count(*): 1
5 rows in set (0.00 sec)



 select * from  performance_schema.threads where PROCESSLIST_STATE in ("executing","Suspending","Waiting on empty queue","waiting for handler commit")\G
*************************** 1. row ***************************
          THREAD_ID: 39
               NAME: thread/innodb/clone_gtid_thread
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: 3358
  PROCESSLIST_STATE: waiting for handler commit
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16248
     RESOURCE_GROUP: NULL
*************************** 2. row ***************************
          THREAD_ID: 48
               NAME: thread/sql/event_scheduler
               TYPE: FOREGROUND
     PROCESSLIST_ID: 5
   PROCESSLIST_USER: event_scheduler
   PROCESSLIST_HOST: localhost
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Daemon
   PROCESSLIST_TIME: 3358
  PROCESSLIST_STATE: Waiting on empty queue
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16253
     RESOURCE_GROUP: NULL
*************************** 3. row ***************************
          THREAD_ID: 52
               NAME: thread/sql/compress_gtid_table
               TYPE: FOREGROUND
     PROCESSLIST_ID: 7
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Daemon
   PROCESSLIST_TIME: 3358
  PROCESSLIST_STATE: Suspending
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16257
     RESOURCE_GROUP: NULL
*************************** 4. row ***************************
          THREAD_ID: 54
               NAME: thread/sql/one_connection
               TYPE: FOREGROUND
     PROCESSLIST_ID: 9
   PROCESSLIST_USER: root
   PROCESSLIST_HOST: localhost
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Query
   PROCESSLIST_TIME: 0
  PROCESSLIST_STATE: executing
   PROCESSLIST_INFO: select * from  performance_schema.threads where PROCESSLIST_STATE in ("executing","Suspending","Waiting on empty queue","waiting for handler commit")
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: Socket
       THREAD_OS_ID: 48638
     RESOURCE_GROUP: NULL
4 rows in set (0.01 sec)



  select count(*) from performance_schema.threads group by(PROCESSLIST_STATE) \G
*************************** 1. row ***************************
count(*): 40
*************************** 2. row ***************************
count(*): 1
*************************** 3. row ***************************
count(*): 1
*************************** 4. row ***************************
count(*): 1
*************************** 5. row ***************************
count(*): 1
5 rows in set (0.00 sec)




% ps -AvM 1733 |grep 1733 |wc -l
      40




select * from  performance_schema.threads where name ="thread/sql/one_connection" \G
*************************** 1. row ***************************
          THREAD_ID: 54
               NAME: thread/sql/one_connection
               TYPE: FOREGROUND
     PROCESSLIST_ID: 9
   PROCESSLIST_USER: root
   PROCESSLIST_HOST: localhost
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Query
   PROCESSLIST_TIME: 0
  PROCESSLIST_STATE: executing
   PROCESSLIST_INFO: select * from  performance_schema.threads where name ="thread/sql/one_connection"
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: Socket
       THREAD_OS_ID: 48638
     RESOURCE_GROUP: NULL
*************************** 2. row ***************************
          THREAD_ID: 55
               NAME: thread/sql/one_connection
               TYPE: FOREGROUND
     PROCESSLIST_ID: 10
   PROCESSLIST_USER: root
   PROCESSLIST_HOST: localhost
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Sleep
   PROCESSLIST_TIME: 141
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: Socket
       THREAD_OS_ID: 53792
     RESOURCE_GROUP: NULL
2 rows in set (0.00 sec)


可以看到另一个链接的processing state 是  PROCESSLIST_STATE: NULL



我们让另外一个线程sleep 
 select sleep(2);
+----------+
| sleep(2) |
+----------+
|        0 |
+----------+
1 row in set (2.06 sec)


 select count(*) from performance_schema.threads group by(PROCESSLIST_STATE) \G
*************************** 1. row ***************************
count(*): 39
*************************** 2. row ***************************
count(*): 1
*************************** 3. row ***************************
count(*): 1
*************************** 4. row ***************************
count(*): 1
*************************** 5. row ***************************
count(*): 1
*************************** 6. row ***************************
count(*): 1
6 rows in set (0.00 sec)



 select PROCESSLIST_STATE,count(*) from performance_schema.threads group by(PROCESSLIST_STATE) \G
*************************** 1. row ***************************
PROCESSLIST_STATE: NULL
         count(*): 39
*************************** 2. row ***************************
PROCESSLIST_STATE: waiting for handler commit
         count(*): 1
*************************** 3. row ***************************
PROCESSLIST_STATE: Waiting on empty queue
         count(*): 1
*************************** 4. row ***************************
PROCESSLIST_STATE: Suspending
         count(*): 1
*************************** 5. row ***************************
PROCESSLIST_STATE: executing
         count(*): 1
*************************** 6. row ***************************
PROCESSLIST_STATE: User sleep
         count(*): 1
6 rows in set (0.00 sec)


多了一个状态userSleep


mysql> show processlist \G
*************************** 1. row ***************************
     Id: 5
   User: event_scheduler
   Host: localhost
     db: NULL
Command: Daemon
   Time: 4644
  State: Waiting on empty queue
   Info: NULL
*************************** 2. row ***************************
     Id: 9
   User: root
   Host: localhost
     db: NULL
Command: Query
   Time: 0
  State: init
   Info: show processlist
2 rows in set (0.00 sec)




select * from performance_schema.threads where PROCESSLIST_ID in(5,9)\G
*************************** 1. row ***************************
          THREAD_ID: 48
               NAME: thread/sql/event_scheduler
               TYPE: FOREGROUND
     PROCESSLIST_ID: 5
   PROCESSLIST_USER: event_scheduler
   PROCESSLIST_HOST: localhost
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Daemon
   PROCESSLIST_TIME: 4694
  PROCESSLIST_STATE: Waiting on empty queue
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16253
     RESOURCE_GROUP: NULL
*************************** 2. row ***************************
          THREAD_ID: 54
               NAME: thread/sql/one_connection
               TYPE: FOREGROUND
     PROCESSLIST_ID: 9
   PROCESSLIST_USER: root
   PROCESSLIST_HOST: localhost
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Query
   PROCESSLIST_TIME: 0
  PROCESSLIST_STATE: executing
   PROCESSLIST_INFO: select * from performance_schema.threads where PROCESSLIST_ID in(5,9)
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: Socket
       THREAD_OS_ID: 48638
     RESOURCE_GROUP: NULL
2 rows in set (0.00 sec)




  select PROCESSLIST_ID,count(*) from performance_schema.threads group by(PROCESSLIST_ID) \G

*************************** 1. row ***************************
PROCESSLIST_ID: NULL
      count(*): 40
*************************** 2. row ***************************
PROCESSLIST_ID: 5
      count(*): 1
*************************** 3. row ***************************
PROCESSLIST_ID: 7
      count(*): 1
*************************** 4. row ***************************
PROCESSLIST_ID: 9
      count(*): 1
4 rows in set (0.00 sec)




 select PROCESSLIST_STATE,count(*) from performance_schema.threads where PROCESSLIST_ID is NULL group by(PROCESSL
*************************** 1. row ***************************
PROCESSLIST_STATE: NULL
         count(*): 39
*************************** 2. row ***************************
PROCESSLIST_STATE: waiting for handler commit
         count(*): 1
2 rows in set (0.00 sec)



select * from  performance_schema.threads where PROCESSLIST_STATE = "waiting for handler commit"\G
*************************** 1. row ***************************
          THREAD_ID: 39
               NAME: thread/innodb/clone_gtid_thread
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: 4928
  PROCESSLIST_STATE: waiting for handler commit
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16248
     RESOURCE_GROUP: NULL
1 row in set (0.00 sec)


https://www.cnblogs.com/bianxj/articles/9605067.html



select type,count(*) from  performance_schema.threads group by type\G
*************************** 1. row ***************************
    type: BACKGROUND
count(*): 40
*************************** 2. row ***************************
    type: FOREGROUND
count(*): 3
2 rows in set (0.00 sec)


select * from  performance_schema.threads where type="FOREGROUND"\G
*************************** 1. row ***************************
          THREAD_ID: 48
               NAME: thread/sql/event_scheduler
               TYPE: FOREGROUND
     PROCESSLIST_ID: 5
   PROCESSLIST_USER: event_scheduler
   PROCESSLIST_HOST: localhost
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Daemon
   PROCESSLIST_TIME: 5320
  PROCESSLIST_STATE: Waiting on empty queue
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16253
     RESOURCE_GROUP: NULL
*************************** 2. row ***************************
          THREAD_ID: 52
               NAME: thread/sql/compress_gtid_table
               TYPE: FOREGROUND
     PROCESSLIST_ID: 7
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Daemon
   PROCESSLIST_TIME: 5320
  PROCESSLIST_STATE: Suspending
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16257
     RESOURCE_GROUP: NULL
*************************** 3. row ***************************
          THREAD_ID: 71
               NAME: thread/sql/one_connection
               TYPE: FOREGROUND
     PROCESSLIST_ID: 26
   PROCESSLIST_USER: root
   PROCESSLIST_HOST: localhost
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Query
   PROCESSLIST_TIME: 0
  PROCESSLIST_STATE: executing
   PROCESSLIST_INFO: select * from  performance_schema.threads where type="FOREGROUND"
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: Socket
       THREAD_OS_ID: 53792
     RESOURCE_GROUP: NULL
3 rows in set (0.00 sec)




show processlist;
+----+-----------------+-----------+------+---------+------+------------------------+------------------+
| Id | User            | Host      | db   | Command | Time | State                  | Info             |
+----+-----------------+-----------+------+---------+------+------------------------+------------------+
|  5 | event_scheduler | localhost | NULL | Daemon  | 5338 | Waiting on empty queue | NULL             |
| 26 | root            | localhost | NULL | Query   |    0 | init                   | show processlist |
+----+-----------------+-----------+------+---------+------+------------------------+------------------+
2 rows in set (0.00 sec)





 select THREAD_OS_ID,count(*) from  performance_schema.threads group by THREAD_OS_ID\G
*************************** 1. row ***************************
THREAD_OS_ID: 14370
    count(*): 1
*************************** 2. row ***************************
THREAD_OS_ID: 15698
    count(*): 1
*************************** 3. row ***************************
THREAD_OS_ID: 15820
    count(*): 1
*************************** 4. row ***************************
THREAD_OS_ID: 15821
    count(*): 1
*************************** 5. row ***************************
THREAD_OS_ID: 15822
    count(*): 1
*************************** 6. row ***************************
THREAD_OS_ID: 15823
    count(*): 1
*************************** 7. row ***************************
THREAD_OS_ID: 15824
    count(*): 1
*************************** 8. row ***************************
THREAD_OS_ID: 15825
    count(*): 1
*************************** 9. row ***************************
THREAD_OS_ID: 15826
    count(*): 1
*************************** 10. row ***************************
THREAD_OS_ID: 15827
    count(*): 1
*************************** 11. row ***************************
THREAD_OS_ID: 15828
    count(*): 1
*************************** 12. row ***************************
THREAD_OS_ID: 15829
    count(*): 1
*************************** 13. row ***************************
THREAD_OS_ID: 15830
    count(*): 1
*************************** 14. row ***************************
THREAD_OS_ID: 15832
    count(*): 1
*************************** 15. row ***************************
THREAD_OS_ID: 15833
    count(*): 1
*************************** 16. row ***************************
THREAD_OS_ID: 15834
    count(*): 1
*************************** 17. row ***************************
THREAD_OS_ID: 15835
    count(*): 1
*************************** 18. row ***************************
THREAD_OS_ID: 15836
    count(*): 1
*************************** 19. row ***************************
THREAD_OS_ID: 16102
    count(*): 1
*************************** 20. row ***************************
THREAD_OS_ID: 16103
    count(*): 1
*************************** 21. row ***************************
THREAD_OS_ID: 16104
    count(*): 1
*************************** 22. row ***************************
THREAD_OS_ID: 16236
    count(*): 1
*************************** 23. row ***************************
THREAD_OS_ID: 16237
    count(*): 1
*************************** 24. row ***************************
THREAD_OS_ID: 16238
    count(*): 1
*************************** 25. row ***************************
THREAD_OS_ID: 16239
    count(*): 1
*************************** 26. row ***************************
THREAD_OS_ID: 16241
    count(*): 1
*************************** 27. row ***************************
THREAD_OS_ID: 16242
    count(*): 1
*************************** 28. row ***************************
THREAD_OS_ID: 16243
    count(*): 1
*************************** 29. row ***************************
THREAD_OS_ID: 16247
    count(*): 1
*************************** 30. row ***************************
THREAD_OS_ID: 16248
    count(*): 1
*************************** 31. row ***************************
THREAD_OS_ID: 16249
    count(*): 2
*************************** 32. row ***************************
THREAD_OS_ID: 16250
    count(*): 2
*************************** 33. row ***************************
THREAD_OS_ID: 16251
    count(*): 2
*************************** 34. row ***************************
THREAD_OS_ID: 16252
    count(*): 2
*************************** 35. row ***************************
THREAD_OS_ID: 16253
    count(*): 1
*************************** 36. row ***************************
THREAD_OS_ID: 16254
    count(*): 1
*************************** 37. row ***************************
THREAD_OS_ID: 16255
    count(*): 1
*************************** 38. row ***************************
THREAD_OS_ID: 16257
    count(*): 1
*************************** 39. row ***************************
THREAD_OS_ID: 53792
    count(*): 1
39 rows in set (0.00 sec)



select THREAD_OS_ID,count(*) from  performance_schema.threads group by THREAD_OS_ID having count(*)>1\G
*************************** 1. row ***************************
THREAD_OS_ID: 16249
    count(*): 2
*************************** 2. row ***************************
THREAD_OS_ID: 16250
    count(*): 2
*************************** 3. row ***************************
THREAD_OS_ID: 16251
    count(*): 2
*************************** 4. row ***************************
THREAD_OS_ID: 16252
    count(*): 2
4 rows in set (0.00 sec)









select * from  performance_schema.threads where THREAD_OS_ID in (16249,16250,16251,16252)\G
*************************** 1. row ***************************
          THREAD_ID: 40
               NAME: thread/innodb/srv_purge_thread
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: 5577
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16249
     RESOURCE_GROUP: NULL
*************************** 2. row ***************************
          THREAD_ID: 42
               NAME: thread/innodb/srv_purge_thread
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: NULL
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16249
     RESOURCE_GROUP: NULL
*************************** 3. row ***************************
          THREAD_ID: 41
               NAME: thread/innodb/srv_worker_thread
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: 5577
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16250
     RESOURCE_GROUP: NULL
*************************** 4. row ***************************
          THREAD_ID: 43
               NAME: thread/innodb/srv_worker_thread
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: NULL
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16250
     RESOURCE_GROUP: NULL
*************************** 5. row ***************************
          THREAD_ID: 44
               NAME: thread/innodb/srv_worker_thread
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: 5577
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16251
     RESOURCE_GROUP: NULL
*************************** 6. row ***************************
          THREAD_ID: 46
               NAME: thread/innodb/srv_worker_thread
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: NULL
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16251
     RESOURCE_GROUP: NULL
*************************** 7. row ***************************
          THREAD_ID: 45
               NAME: thread/innodb/srv_worker_thread
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: 5577
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16252
     RESOURCE_GROUP: NULL
*************************** 8. row ***************************
          THREAD_ID: 47
               NAME: thread/innodb/srv_worker_thread
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: NULL
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 16252
     RESOURCE_GROUP: NULL
8 rows in set (0.00 sec)



其中6个srv_worker_thread 用了三个系统THREAD_OS_ID Id
两个srv_purge_thread 用了一个系统THREAD_OS_ID Id



select count(distinct(THREAD_OS_ID)) from  performance_schema.threads;
+-------------------------------+
| count(distinct(THREAD_OS_ID)) |
+-------------------------------+
|                            39 |
+-------------------------------+
1 row in set (0.00 sec)


再起一个mysqlclient

 % ps -AvM 1733 |grep 1733 |wc -l
      41


select count(distinct(THREAD_OS_ID)) from  performance_schema.threads;
+-------------------------------+
| count(distinct(THREAD_OS_ID)) |
+-------------------------------+
|                            41 |
+-------------------------------+
1 row in set (0.00 sec)




退出一个client
mysql> select count(distinct(THREAD_OS_ID)) from  performance_schema.threads;
+-------------------------------+
| count(distinct(THREAD_OS_ID)) |
+-------------------------------+
|                            40 |
+-------------------------------+
1 row in set (0.00 sec)

再腿出一个
 select count(distinct(THREAD_OS_ID)) from  performance_schema.threads;
+-------------------------------+
| count(distinct(THREAD_OS_ID)) |
+-------------------------------+
|                            39 |
+-------------------------------+
1 row in set (0.00 sec)


其实就是39个os线程
还有两个被缓存了
所以
% ps -AvM 1733 |grep 1733 |wc -l
      41

41=39+2


show variables like "%thread_cache_size%";
+-------------------+-------+
| Variable_name     | Value |
+-------------------+-------+
| thread_cache_size | 9     |
+-------------------+-------+
1 row in set (0.00 sec)


改成2
 set global thread_cache_size=2;
Query OK, 0 rows affected (0.00 sec)


% ps -AvM 1733 |grep 1733 |wc -l
      41


select count(distinct(THREAD_OS_ID)) from  performance_schema.threads;
+-------------------------------+
| count(distinct(THREAD_OS_ID)) |
+-------------------------------+
|                            39 |
+-------------------------------+
1 row in set (0.00 sec)

可以看到缓存两个