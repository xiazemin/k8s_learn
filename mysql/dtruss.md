% lsof -i tcp:3306
COMMAND  PID     USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
mysqld  1733 xiazemin   31u  IPv4 0x362a06740147157f      0t0  TCP localhost:mysql (LISTEN)

sudo dtruss -ap 1733
sudo dtrace -n 'syscall::open*:entry { printf("%s %s",execname,copyinstr(arg0)); }'

sudo dtrace -n 'syscall::open*:entry { printf("%s %s",execname,copyinstr(arg0)); }' -p 1733
