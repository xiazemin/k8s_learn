 select thread_id, thread_os_id, name from performance_schema.threads ;
+-----------+--------------+---------------------------------------------+
| thread_id | thread_os_id | name                                        |
+-----------+--------------+---------------------------------------------+
|         1 |         6716 | thread/sql/main                             |
|         2 |         7490 | thread/mysys/thread_timer_notifier          |
|         4 |         7604 | thread/innodb/io_ibuf_thread                |
|         5 |         7605 | thread/innodb/io_log_thread                 |
|         6 |         7606 | thread/innodb/io_read_thread                |
|         7 |         7607 | thread/innodb/io_read_thread                |
|         8 |         7608 | thread/innodb/io_read_thread                |
|         9 |         7609 | thread/innodb/io_read_thread                |
|        10 |         7610 | thread/innodb/io_write_thread               |
|        11 |         7611 | thread/innodb/io_write_thread               |
|        12 |         7612 | thread/innodb/io_write_thread               |
|        13 |         7613 | thread/innodb/io_write_thread               |
|        14 |         7614 | thread/innodb/page_flush_coordinator_thread |
|        15 |         7655 | thread/innodb/log_checkpointer_thread       |
|        16 |         7656 | thread/innodb/log_flush_notifier_thread     |
|        17 |         7657 | thread/innodb/log_flusher_thread            |
|        18 |         7658 | thread/innodb/log_write_notifier_thread     |
|        19 |         7659 | thread/innodb/log_writer_thread             |
|        24 |         7986 | thread/innodb/srv_lock_timeout_thread       |
|        25 |         7987 | thread/innodb/srv_error_monitor_thread      |
|        26 |         7988 | thread/innodb/srv_monitor_thread            |
|        27 |         8044 | thread/innodb/buf_resize_thread             |
|        28 |         8045 | thread/innodb/srv_master_thread             |
|        29 |         8046 | thread/innodb/dict_stats_thread             |
|        30 |         8047 | thread/innodb/fts_optimize_thread           |
|        31 |         8093 | thread/mysqlx/worker                        |
|        32 |         8094 | thread/mysqlx/worker                        |
|        33 |         8095 | thread/mysqlx/acceptor_network              |
|        39 |         8384 | thread/innodb/buf_dump_thread               |
|        40 |         8385 | thread/innodb/clone_gtid_thread             |
|        41 |         8386 | thread/innodb/srv_purge_thread              |
|        42 |         8386 | thread/innodb/srv_purge_thread              |
|        43 |         8387 | thread/innodb/srv_worker_thread             |
|        44 |         8388 | thread/innodb/srv_worker_thread             |
|        45 |         8387 | thread/innodb/srv_worker_thread             |
|        46 |         8389 | thread/innodb/srv_worker_thread             |
|        47 |         8388 | thread/innodb/srv_worker_thread             |
|        48 |         8389 | thread/innodb/srv_worker_thread             |
|        49 |         8390 | thread/sql/event_scheduler                  |
|        50 |         8391 | thread/sql/signal_handler                   |
|        51 |         8392 | thread/mysqlx/acceptor_network              |
|        53 |         8394 | thread/sql/compress_gtid_table              |
|    270178 |      1680712 | thread/sql/one_connection                   |
+-----------+--------------+---------------------------------------------+
43 rows in set (0.00 sec)

每连接一个mysqLclient 
thread/sql/one_connection    加一

SELECT * FROM performance_schema.socket_instances\G
*************************** 1. row ***************************
           EVENT_NAME: wait/io/socket/mysqlx/tcpip_socket
OBJECT_INSTANCE_BEGIN: 4715380736
            THREAD_ID: 51
            SOCKET_ID: 29
                   IP: 127.0.0.1
                 PORT: 33060
                STATE: ACTIVE
*************************** 2. row ***************************
           EVENT_NAME: wait/io/socket/mysqlx/unix_socket
OBJECT_INSTANCE_BEGIN: 4715381048
            THREAD_ID: 51
            SOCKET_ID: 30
                   IP:
                 PORT: 0
                STATE: ACTIVE
*************************** 3. row ***************************
           EVENT_NAME: wait/io/socket/sql/server_tcpip_socket
OBJECT_INSTANCE_BEGIN: 4715381360
            THREAD_ID: 1
            SOCKET_ID: 39
                   IP: 127.0.0.1
                 PORT: 3306
                STATE: ACTIVE
*************************** 4. row ***************************
           EVENT_NAME: wait/io/socket/sql/server_unix_socket
OBJECT_INSTANCE_BEGIN: 4715381672
            THREAD_ID: 1
            SOCKET_ID: 41
                   IP:
                 PORT: 0
                STATE: ACTIVE
*************************** 5. row ***************************
           EVENT_NAME: wait/io/socket/sql/client_connection
OBJECT_INSTANCE_BEGIN: 4715450000
            THREAD_ID: 270178
            SOCKET_ID: 42
                   IP:
                 PORT: 0
                STATE: IDLE
*************************** 6. row ***************************
           EVENT_NAME: wait/io/socket/sql/client_connection
OBJECT_INSTANCE_BEGIN: 4715451560
            THREAD_ID: 270183
            SOCKET_ID: 46
                   IP:
                 PORT: 0
                STATE: ACTIVE
6 rows in set (0.02 sec)


mysql> \! cat /var/run/mysqld/mysqld.pid
cat: /var/run/mysqld/mysqld.pid: No such file or directory