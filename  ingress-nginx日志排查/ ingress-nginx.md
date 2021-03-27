% kubectl get pod -n ingress-nginx
NAME                                       READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create-gfv8s       0/1     Completed   0          22h
ingress-nginx-controller-84599cfff-l77nt   1/1     Running     3          22h


 kubectl logs ingress-nginx-controller-84599cfff-l77nt -n ingress-nginx

 192.168.64.6 - - [27/Mar/2021:01:55:51 +0000] "GET /apple HTTP/1.1" 404 146 "-" "curl/7.64.1" 78 0.005 [upstream-default-backend] [] 127.0.0.1:8181 146 0.005 404 b8fce6c55c43933bba328d25f0f2400b


 2021/03/27 01:58:00 [error] 318#318: *75161 lua entry thread aborted: memory allocation error: not enough memory
stack traceback:
coroutine 0:
	[C]: in function 'clone_tab'
	/etc/nginx/lua/monitor.lua:60: in function </etc/nginx/lua/monitor.lua:51>, context: ngx.timer
2021/03/27 01:58:00 [emerg] 687#687: io_setup() failed (38: Function not implemented)

% curl http://localhost/apple
apple%

可能是ingress-nginx 的bug，也可能是系统缺少插件，每次更新ingress 需要重启k8s才生效