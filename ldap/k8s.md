为了方便管理和集成jenkins，k8s、harbor、jenkins均使用openLDAP统一认证。

https://blog.csdn.net/weixin_30457551/article/details/95853707

将openLDAP部署在k8s上，openLDAP可以在集群之外存在，不一定非要在k8s上部署openLDAP。

https://github.com/dotbalo/k8s

https://github.com/dotbalo/k8s/tree/master/openldap

https://github.com/nginxinc/nginx-ldap-auth

https://blog.csdn.net/weixin_33842304/article/details/94743413

LDAP 上自定义一个 token 的 schema，然后将在 k8s 集群上已创建好的 service account 的 token 跟 LDAP 绑定起来。这样，在进行 webhook 认证时，便能带着 token 去 LDAP 上进行认证。另外，还有些是对接上 OpenID connect token，然后再这些服务提供系统那里对接上 LDAP。

使用 webhook 进行 LDAP 认证，平常的用户名密码认证行不通的主要原因是 k8s apiserver 进行认证请求时，带过来只有用户名和 bearer token，没有密码。为了能够使用用户名密码和 bearer token 去 ldap 进行认证，这里做了些 trick 的东西。

客户端通过 nginx 访问 apiserver，在 nginx 一层里，配置了 auth-request，将 basic auth 的请求发送给后端的 ldap 认证代理，ldap 认证代理认证通过后，会随机生成一段 bearer token，并通过相应头部告诉给 nginx。nginx 收到这个响应头部后，就配置使用 bearer token 访问 apiserver。也就是说，请求到达 apiserver 的时候，使用的是 bearer token 这种方式进行认证。apiserver 收到请求后，会同时出发内部认证机制和 webhook 认证机制，webhook 认证在这里同样配置了 nginx 的 ldap 认证代理。由于 ldap 认证代理保存了所有经过 ldap 认证的用户的 bearer token，因此便可以通过 apiserver 带过来的 bearer token 成功认证用户！


location / {
       auth_request /auth-proxy;       # 这里获取从ldap认证服务返回来的bearer token 
       auth_request_set $bearer_token $upstream_http_x_token;
       proxy_set_header   Host             $host;
       proxy_set_header X-Real-IP $remote_addr;
       proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;       # 设置bearer token认证
       proxy_set_header Authorization "Bearer $bearer_token";
       proxy_ssl_verify       off;
       proxy_ssl_session_reuse on;
       proxy_pass         https://dashboard-k8s;
    }
 
    location = /auth-proxy {
        internal;
        proxy_pass http://k8s-ldap-backend/ldap-auth;
        proxy_pass_request_body off;
        proxy_set_header Content-Length "";
        proxy_cache ldap_auth_cache;
        proxy_cache_valid 200 30m;
        proxy_cache_key "$http_authorization$cookie_nginxauth";
    }

https://www.imooc.com/article/275600

实现方式其实很简单，首先写一个登录界面与公司的AD进行打通获取用户与组，然后将用户或者组与k8s集群中的 service account 进行关联就实现了对应的rbac与登录token，最后在登录后实现一个反向代理服务即可完成。

https://segmentfault.com/a/1190000022908617


https://blog.inkubate.io/access-your-kubernetes-cluster-with-your-active-directory-credentials/

https://blog.poychang.net/ldap-introduction/

https://www.cnblogs.com/dukuan/p/9983899.html

https://www.jianshu.com/p/e8a9d649f617?from=timeline@

https://icicimov.github.io/blog/virtualization/Kubernetes-LDAP-Authentication/
https://medium.com/@pmvk/step-by-step-guide-to-integrate-ldap-with-kubernetes-1f3fe1ec644e

k8s 部署 openldap
https://itwenti.com/?p=1159

https://cloud.tencent.com/developer/article/1684908

https://github.com/apprenda-kismatic/kubernetes-ldap

　　1） 与容器云相比，容器云的权限控制更为完善，但是实现基于用户的验证，能够区分openLDAP中不一样的用户不一样的权限，而k8s被统一成了kubernetes-dashboard这个用户，也有多是本身没有配置成功，后期须要再次确认。

　　2） 能够限制openLDAP中的用户只访问某些namespace，须要自行定义权限。

　　3） 公司没有须要无需让非ops员工访问k8s。

　　4） k8s使用openLDAP登陆非必须。

　　5） 网上也有大神使用schema登陆的:https://icicimov.github.io/blog/virtualization/Kubernetes-LDAP-Authentication/，看了一下不是本身须要的。

https://www.shangmayuan.com/a/d9b7d002dc5a4cc382c6efd7.html

https://www.iteye.com/blog/wx1569466809-2485346

https://www.bbsmax.com/A/l1dyK6gbJe/

https://www.136.la/net/show-44074.html

https://github.com/nginxinc/nginx-ldap-auth
https://github.com/vesse/passport-ldapauth
