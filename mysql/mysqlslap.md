 % mysqlslap -uroot  --concurrency=100 --iterations=3 --auto-generate-sql --auto-generate-sql-load-type=mixed --auto-generate-sql-add-autoincrement --engine=innodb --number-of-queries=500000
^@^@^@^@Benchmark
	Running for engine innodb
	Average number of seconds to run all queries: 17.124 seconds
	Minimum number of seconds to run all queries: 16.405 seconds
	Maximum number of seconds to run all queries: 18.418 seconds
	Number of clients running queries: 100
	Average number of queries per client: 5000



 select count(*) from performance_schema.threads where PROCESSLIST_COMMAND='Query'\G
*************************** 1. row ***************************
count(*): 101
1 row in set (0.00 sec)

select count(*) from performance_schema.threads where PROCESSLIST_COMMAND!='Query'\G
*************************** 1. row ***************************
count(*): 3
1 row in set (0.00 sec)


 show processlist;
+------+-----------------+-----------+--------------------+---------+------+----------------------------+------------------------------------------------------------------------------------------------------+
| Id   | User            | Host      | db                 | Command | Time | State                      | Info                                                                                                 |
+------+-----------------+-----------+--------------------+---------+------+----------------------------+------------------------------------------------------------------------------------------------------+
|    5 | event_scheduler | localhost | NULL               | Daemon  |  896 | Waiting on empty queue     | NULL                                                                                                 |
|  222 | root            | localhost | performance_schema | Query   |    0 | init                       | show processlist                                                                                     |
|  324 | root            | localhost | mysqlslap          | Sleep   |    2 |                            | NULL                                                                                                 |
| 1925 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1926 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,1759592334,'3lkoxjtvgLu5xKHSTTtJuGE5F5QqmCcppCTmvFZScRZQgim93gSxwb24gKmI |
| 1927 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1928 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1929 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 1930 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 1931 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 1932 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1933 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1934 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,1759592334,'3lkoxjtvgLu5xKHSTTtJuGE5F5QqmCcppCTmvFZScRZQgim93gSxwb24gKmI |
| 1935 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 1936 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1937 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,1759592334,'3lkoxjtvgLu5xKHSTTtJuGE5F5QqmCcppCTmvFZScRZQgim93gSxwb24gKmI |
| 1938 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1939 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 1940 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1941 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 1942 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1943 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1944 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 1945 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1946 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1947 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 1948 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,1759592334,'3lkoxjtvgLu5xKHSTTtJuGE5F5QqmCcppCTmvFZScRZQgim93gSxwb24gKmI |
| 1949 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1950 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 1951 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 1952 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 1953 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 1954 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1955 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1956 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1957 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1958 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 1959 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 1960 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 1961 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 1962 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,1759592334,'3lkoxjtvgLu5xKHSTTtJuGE5F5QqmCcppCTmvFZScRZQgim93gSxwb24gKmI |
| 1963 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 1964 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,1759592334,'3lkoxjtvgLu5xKHSTTtJuGE5F5QqmCcppCTmvFZScRZQgim93gSxwb24gKmI |
| 1965 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,1759592334,'3lkoxjtvgLu5xKHSTTtJuGE5F5QqmCcppCTmvFZScRZQgim93gSxwb24gKmI |
| 1966 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1967 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 1968 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 1969 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 1970 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 1971 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 1972 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 1973 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 1974 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 1975 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1976 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 1977 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 1978 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1979 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 1980 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 1981 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1982 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1983 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 1984 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1985 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 1986 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,1759592334,'3lkoxjtvgLu5xKHSTTtJuGE5F5QqmCcppCTmvFZScRZQgim93gSxwb24gKmI |
| 1987 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 1988 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1989 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 1990 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1991 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 1992 | root            | localhost | mysqlslap          | Query   |    0 | Sending to client          | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 1993 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,1759592334,'3lkoxjtvgLu5xKHSTTtJuGE5F5QqmCcppCTmvFZScRZQgim93gSxwb24gKmI |
| 1994 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1995 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1996 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 1997 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 1998 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 1999 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 2000 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 2001 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 2002 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 2003 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,1759592334,'3lkoxjtvgLu5xKHSTTtJuGE5F5QqmCcppCTmvFZScRZQgim93gSxwb24gKmI |
| 2004 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 2005 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 2006 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 2007 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 2008 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,100669,'qnMdipW5KkXdTjGCh2PNzLoeR0527frpQDQ8uw67Ydk1K06uuNHtkxYBxT5w8plb |
| 2009 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 2010 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 2011 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 2012 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 2013 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 2014 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 2015 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 2016 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 2017 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,95275444,'bNIrBDBl81tjzdvuOpQRCXgX37xGtzLKEXBIcE3k7xK7aFtqxC99jqYnpTviK8 |
| 2018 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 2019 | root            | localhost | mysqlslap          | Sleep   |    0 |                            | NULL                                                                                                 |
| 2020 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
| 2021 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 2022 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,73673339,'BN3152Gza4GW7atxJKACYwJqDbFynLxqc0kh30YTwgz3FktQ43XTrqJ4PQ25fr |
| 2023 | root            | localhost | mysqlslap          | Query   |    0 | update                     | INSERT INTO t1 VALUES (NULL,364531492,'qMa5SuKo4M5OM7ldvisSc6WK9rsG9E8sSixocHdgfa5uiiNTGFxkDJ4EAwWC2 |
| 2024 | root            | localhost | mysqlslap          | Query   |    0 | waiting for handler commit | INSERT INTO t1 VALUES (NULL,866596855,'naQuzhMt1IrZIJMkbLAKBNNKKK2sCknzI5uHeGAgQuDd5SLgpN0smODyc7qor |
+------+-----------------+-----------+--------------------+---------+------+----------------------------+------------------------------------------------------------------------------------------------------+
103 rows in set (0.00 sec)


 select PROCESSLIST_COMMAND,count(*) from performance_schema.threads group by  PROCESSLIST_COMMAND\G
*************************** 1. row ***************************
PROCESSLIST_COMMAND: Query
           count(*): 51
*************************** 2. row ***************************
PROCESSLIST_COMMAND: NULL
           count(*): 40
*************************** 3. row ***************************
PROCESSLIST_COMMAND: Sleep
           count(*): 51
*************************** 4. row ***************************
PROCESSLIST_COMMAND: Daemon
           count(*): 2
4 rows in set (0.00 sec)



 select * from performance_schema.threads where PROCESSLIST_COMMAND!='Query'\G
*************************** 1. row ***************************
          THREAD_ID: 47
               NAME: thread/sql/event_scheduler
               TYPE: FOREGROUND
     PROCESSLIST_ID: 5
   PROCESSLIST_USER: event_scheduler
   PROCESSLIST_HOST: localhost
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Daemon
   PROCESSLIST_TIME: 1301
  PROCESSLIST_STATE: Waiting on empty queue
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 9759
     RESOURCE_GROUP: NULL
*************************** 2. row ***************************
          THREAD_ID: 51
               NAME: thread/sql/compress_gtid_table
               TYPE: FOREGROUND
     PROCESSLIST_ID: 7
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Daemon
   PROCESSLIST_TIME: 1301
  PROCESSLIST_STATE: Suspending
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 9763
     RESOURCE_GROUP: NULL
*************************** 3. row ***************************
          THREAD_ID: 3569
               NAME: thread/sql/one_connection
               TYPE: FOREGROUND
     PROCESSLIST_ID: 3525
   PROCESSLIST_USER: root
   PROCESSLIST_HOST: localhost
     PROCESSLIST_DB: mysqlslap
PROCESSLIST_COMMAND: Sleep
   PROCESSLIST_TIME: 12
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: Socket
       THREAD_OS_ID: 37081
     RESOURCE_GROUP: NULL
3 rows in set (0.00 sec)



实例3(自定义sql语句)
#mysqlslap -h127.0.0.1 -uroot -p123456789 --concurrency=100 --iterations=1 --create-schema=rudao --query='select * from serverlist;' --engine=innodb --number-of-queries=5000 --debug-info

 % mysqlslap -uroot --concurrency=100 --iterations=1 --create-schema=mytable --query='select * from edge;' --engine=innodb --number-of-queries=5000
Benchmark
	Running for engine innodb
	Average number of seconds to run all queries: 0.124 seconds
	Minimum number of seconds to run all queries: 0.124 seconds
	Maximum number of seconds to run all queries: 0.124 seconds
	Number of clients running queries: 100
	Average number of queries per client: 50



实例4(指定sql脚本)
#mysqlslap -h127.0.0.1 -uroot -p123456789 --concurrency=100 --iterations=1 --create-schema=rudao --query=/tmp/query.sql --engine=innodb --number-of-queries=5000 --debug-info


 % % mysqlslap -uroot --concurrency=100 --iterations=1 --create-schema=mytable --query=Downloads/sqls.json --engine=innodb --number-of-queries=5000
Benchmark
	Running for engine innodb
	Average number of seconds to run all queries: 0.124 seconds
	Minimum number of seconds to run all queries: 0.124 seconds
	Maximum number of seconds to run all queries: 0.124 seconds
	Number of clients running queries: 100
	Average number of queries per client: 50


https://www.cnblogs.com/fjping0606/p/5853325.html



常用参数 [options] 详细说明：

--auto-generate-sql, -a 自动生成测试表和数据，表示用mysqlslap工具自己生成的SQL脚本来测试并发压力。
--auto-generate-sql-load-type=type 测试语句的类型。代表要测试的环境是读操作还是写操作还是两者混合的。取值包括：read，key，write，update和mixed(默认)。
--auto-generate-sql-add-auto-increment 代表对生成的表自动添加auto_increment列，从5.1.18版本开始支持。
--number-char-cols=N, -x N 自动生成的测试表中包含多少个字符类型的列，默认1
--number-int-cols=N, -y N 自动生成的测试表中包含多少个数字类型的列，默认1
--number-of-queries=N 总的测试查询次数(并发客户数×每客户查询次数)
--query=name,-q 使用自定义脚本执行测试，例如可以调用自定义的一个存储过程或者sql语句来执行测试。
--create-schema 代表自定义的测试库名称，测试的schema，MySQL中schema也就是database。
--commint=N 多少条DML后提交一次。
--compress, -C 如果服务器和客户端支持都压缩，则压缩信息传递。
--concurrency=N, -c N 表示并发量，也就是模拟多少个客户端同时执行select。可指定多个值，以逗号或者--delimiter参数指定的值做为分隔符。例如：--concurrency=100,200,500。
--engine=engine_name, -e engine_name 代表要测试的引擎，可以有多个，用分隔符隔开。例如：--engines=myisam,innodb。
--iterations=N, -i N 测试执行的迭代次数，代表要在不同并发环境下，各自运行测试多少次。
--only-print 只打印测试语句而不实际执行。
--detach=N 执行N条语句后断开重连。
--debug-info, -T 打印内存和CPU的相关信息。


测试的过程需要生成测试表，插入测试数据，这个mysqlslap可以自动生成，默认生成一个mysqlslap的schema，如果已经存在则先删除。可以用--only-print来打印实际的测试过程，整个测试完成后不会在数据库中留下痕迹。

https://blog.csdn.net/lin443514407lin/article/details/73277774


如何不让mysqlslap删除生成后的数据

把这个账号的 DROP TABLE的权限取消。

或者修改一下mysqlslap的源代码，重新编译再执行。


缓存了连接
mysqlslap: Error when connecting to server: 1040 Too many connections


show variables like "max_connections";
+-----------------+-------+
| Variable_name   | Value |
+-----------------+-------+
| max_connections | 151   |
+-----------------+-------+
1 row in set (0.05 sec)




 % mysqlslap -uroot --concurrency=50 --iterations=1 --create-schema=svc_tree --query=Downloads/sqls.json --engine=innodb --number-of-queries=5000


  select PROCESSLIST_COMMAND,count(*) from performance_schema.threads group by  PROCESSLIST_COMMAND\G
*************************** 1. row ***************************
PROCESSLIST_COMMAND: NULL
           count(*): 40
*************************** 2. row ***************************
PROCESSLIST_COMMAND: Daemon
           count(*): 2
*************************** 3. row ***************************
PROCESSLIST_COMMAND: Query
           count(*): 51
*************************** 4. row ***************************
PROCESSLIST_COMMAND: Sleep
           count(*): 1
4 rows in set (0.00 sec)


Query:51 有50个并发连接，1个当前client连接
 select * from  performance_schema.threads where PROCESSLIST_COMMAND ='Sleep'\G
*************************** 1. row ***************************
          THREAD_ID: 6531
               NAME: thread/sql/one_connection
               TYPE: FOREGROUND
     PROCESSLIST_ID: 6487
   PROCESSLIST_USER: root
   PROCESSLIST_HOST: localhost
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Sleep
   PROCESSLIST_TIME: 235
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: Socket
       THREAD_OS_ID: 54089
     RESOURCE_GROUP: NULL
1 row in set (0.00 sec)


两个Daemon线程分别是compress_gtid_table和event_scheduler
select * from  performance_schema.threads where PROCESSLIST_COMMAND ='Daemon'\G
*************************** 1. row ***************************
          THREAD_ID: 47
               NAME: thread/sql/event_scheduler
               TYPE: FOREGROUND
     PROCESSLIST_ID: 5
   PROCESSLIST_USER: event_scheduler
   PROCESSLIST_HOST: localhost
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Daemon
   PROCESSLIST_TIME: 3417
  PROCESSLIST_STATE: Waiting on empty queue
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 9759
     RESOURCE_GROUP: NULL
*************************** 2. row ***************************
          THREAD_ID: 51
               NAME: thread/sql/compress_gtid_table
               TYPE: FOREGROUND
     PROCESSLIST_ID: 7
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: Daemon
   PROCESSLIST_TIME: 3417
  PROCESSLIST_STATE: Suspending
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 9763
     RESOURCE_GROUP: NULL
2 rows in set (0.00 sec)


40个NULL线程中33个是innodb相关的，两个thread/mysqlx/acceptor_network，一个thread/sql/main，一个thread/mysys/thread_timer_notifier，两个thread/mysqlx/worker，一个thread/sql/signal_handler

select * from  performance_schema.threads where PROCESSLIST_COMMAND IS NULL and NAME not like 'thread/innodb%'\G
*************************** 1. row ***************************
          THREAD_ID: 1
               NAME: thread/sql/main
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: mysql
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: 3651
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: NULL
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 8395
     RESOURCE_GROUP: NULL
*************************** 2. row ***************************
          THREAD_ID: 2
               NAME: thread/mysys/thread_timer_notifier
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: NULL
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 9554
     RESOURCE_GROUP: NULL
*************************** 3. row ***************************
          THREAD_ID: 31
               NAME: thread/mysqlx/worker
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: NULL
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 9734
     RESOURCE_GROUP: NULL
*************************** 4. row ***************************
          THREAD_ID: 32
               NAME: thread/mysqlx/worker
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: NULL
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 9735
     RESOURCE_GROUP: NULL
*************************** 5. row ***************************
          THREAD_ID: 33
               NAME: thread/mysqlx/acceptor_network
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: NULL
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 9736
     RESOURCE_GROUP: NULL
*************************** 6. row ***************************
          THREAD_ID: 48
               NAME: thread/sql/signal_handler
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: NULL
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 9760
     RESOURCE_GROUP: NULL
*************************** 7. row ***************************
          THREAD_ID: 49
               NAME: thread/mysqlx/acceptor_network
               TYPE: BACKGROUND
     PROCESSLIST_ID: NULL
   PROCESSLIST_USER: NULL
   PROCESSLIST_HOST: NULL
     PROCESSLIST_DB: NULL
PROCESSLIST_COMMAND: NULL
   PROCESSLIST_TIME: NULL
  PROCESSLIST_STATE: NULL
   PROCESSLIST_INFO: NULL
   PARENT_THREAD_ID: 1
               ROLE: NULL
       INSTRUMENTED: YES
            HISTORY: YES
    CONNECTION_TYPE: NULL
       THREAD_OS_ID: 9761
     RESOURCE_GROUP: NULL
7 rows in set (0.01 sec)




select name,count(*) from  performance_schema.threads where PROCESSLIST_COMMAND IS NULL and NAME  like 'thread/innodb%' group by name\G
*************************** 1. row ***************************
    name: thread/innodb/io_ibuf_thread
count(*): 1
*************************** 2. row ***************************
    name: thread/innodb/io_log_thread
count(*): 1
*************************** 3. row ***************************
    name: thread/innodb/io_read_thread
count(*): 4
*************************** 4. row ***************************
    name: thread/innodb/io_write_thread
count(*): 4
*************************** 5. row ***************************
    name: thread/innodb/page_flush_coordinator_thread
count(*): 1
*************************** 6. row ***************************
    name: thread/innodb/log_checkpointer_thread
count(*): 1
*************************** 7. row ***************************
    name: thread/innodb/log_flush_notifier_thread
count(*): 1
*************************** 8. row ***************************
    name: thread/innodb/log_flusher_thread
count(*): 1
*************************** 9. row ***************************
    name: thread/innodb/log_write_notifier_thread
count(*): 1
*************************** 10. row ***************************
    name: thread/innodb/log_writer_thread
count(*): 1
*************************** 11. row ***************************
    name: thread/innodb/srv_lock_timeout_thread
count(*): 1
*************************** 12. row ***************************
    name: thread/innodb/srv_error_monitor_thread
count(*): 1
*************************** 13. row ***************************
    name: thread/innodb/srv_monitor_thread
count(*): 1
*************************** 14. row ***************************
    name: thread/innodb/buf_resize_thread
count(*): 1
*************************** 15. row ***************************
    name: thread/innodb/srv_master_thread
count(*): 1
*************************** 16. row ***************************
    name: thread/innodb/dict_stats_thread
count(*): 1
*************************** 17. row ***************************
    name: thread/innodb/fts_optimize_thread
count(*): 1
*************************** 18. row ***************************
    name: thread/innodb/buf_dump_thread
count(*): 1
*************************** 19. row ***************************
    name: thread/innodb/clone_gtid_thread
count(*): 1
*************************** 20. row ***************************
    name: thread/innodb/srv_purge_thread
count(*): 2
*************************** 21. row ***************************
    name: thread/innodb/srv_worker_thread
count(*): 6
21 rows in set (0.00 sec)