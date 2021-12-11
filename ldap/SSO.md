单点登录英文全称Single Sign On，简称就是SSO。它的解释是：在多个应用系统中，只需要登录一次，就可以访问其他相互信任的应用系统。

普通的登录认证机制:
在浏览器（Browser）中访问一个应用，这个应用需要登录，我们填写完用户名和密码后，完成登录认证。这时，我们在这个用户的session中标记登录状态为yes（已登录），同时在浏览器（Browser）中写入Cookie，这个Cookie是这个用户的唯一标识。下次我们再访问这个应用的时候，请求中会带上这个Cookie，服务端会根据这个Cookie找到对应的session，通过session来判断这个用户是否登录。如果不做特殊配置，这个Cookie的名字叫做jsessionid，值在服务端（server）是唯一的。


同域下的单点登录:
一个企业一般情况下只有一个域名，通过二级域名区分不同的系统。比如我们有个域名叫做：a.com，同时有两个业务系统分别为：app1.a.com和app2.a.com。我们要做单点登录（SSO），需要一个登录系统，叫做：sso.a.com。


我们只要在sso.a.com登录，app1.a.com和app2.a.com就也登录了。通过上面的登陆认证机制，我们可以知道，在sso.a.com中登录了，其实是在sso.a.com的服务端的session中记录了登录状态，同时在浏览器端（Browser）的sso.a.com下写入了Cookie。那么我们怎么才能让app1.a.com和app2.a.com登录呢？这里有两个问题：

Cookie是不能跨域的，我们Cookie的domain属性是sso.a.com，在给app1.a.com和app2.a.com发送请求是带不上的。
sso、app1和app2是不同的应用，它们的session存在自己的应用内，是不共享的。


针对第一个问题，sso登录以后，可以将Cookie的域设置为顶域，即.a.com，这样所有子域的系统都可以访问到顶域的Cookie。我们在设置Cookie时，只能设置顶域和自己的域，不能设置其他的域。比如：我们不能在自己的系统中给baidu.com的域设置Cookie。

Cookie的问题解决了，我们再来看看session的问题。我们在sso系统登录了，这时再访问app1，Cookie也带到了app1的服务端（Server），app1的服务端怎么找到这个Cookie对应的Session呢？这里就要把3个系统的Session共享，如图所示。共享Session的解决方案有很多，例如：Spring-Session。这样第2个问题也解决了。

同域下的单点登录就实现了，但这还不是真正的单点登录。


不同域下的单点登录:
不同域之间Cookie是不共享的，怎么办？
这里我们就要说一说CAS流程了，这个流程是单点登录的标准流程。
具体流程如下：

用户访问app系统，app系统是需要登录的，但用户现在没有登录。
跳转到CAS server，即SSO登录系统，以后图中的CAS Server我们统一叫做SSO系统。 SSO系统也没有登录，弹出用户登录页。
用户填写用户名、密码，SSO系统进行认证后，将登录状态写入SSO的session，浏览器（Browser）中写入SSO域下的Cookie。
SSO系统登录完成后会生成一个ST（Service Ticket），然后跳转到app系统，同时将ST作为参数传递给app系统。
app系统拿到ST后，从后台向SSO发送请求，验证ST是否有效。
验证通过后，app系统将登录状态写入session并设置app域下的Cookie。
至此，跨域单点登录就完成了。以后我们再访问app系统时，app就是登录的。接下来，我们再看看访问app2系统时的流程。

用户访问app2系统，app2系统没有登录，跳转到SSO。
由于SSO已经登录了，不需要重新登录认证。
SSO生成ST，浏览器跳转到app2系统，并将ST作为参数传递给app2。
app2拿到ST，后台访问SSO，验证ST是否有效。
验证成功后，app2将登录状态写入session，并在app2域下写入Cookie。
这样，app2系统不需要走登录流程，就已经是登录了。SSO，app和app2在不同的域，它们之间的session不共享也是没问题的。
SSO系统登录后，跳回原业务系统时，带了个参数ST，业务系统还要拿ST再次访问SSO进行验证，觉得这个步骤有点多余。他想SSO登录认证通过后，通过回调地址将用户信息返回给原业务系统，原业务系统直接设置登录状态，这样流程简单，也完成了登录，不是很好吗？

其实这样问题时很严重的，如果我在SSO没有登录，而是直接在浏览器中敲入回调的地址，并带上伪造的用户信息，是不是业务系统也认为登录了呢？这是很可怕的。

https://www.jianshu.com/p/75edcc05acfd

开源的有OpenSSO、CAS ，微软的AD SSO，及基于kerberos 的SSO等等
https://www.cnblogs.com/EzrealLiu/p/5559255.html

https://zhuanlan.zhihu.com/p/66037342?ivk_sa=1024320u
https://github.com/ZhongFuCheng3y/athena

https://github.com/buzzfeed/sso


https://github.com/aldaris/opensso


https://github.com/CASMPostol

https://github.com/github/cas-overlay

https://blog.csdn.net/guyan0319/article/details/106697940/

https://github.com/guyan0319/go-sso/

https://studygolang.com/articles/31375


仅使用jwt实现单点登录会遇到两个问题

用户无法主动登出，即服务端发出token后，无法主动销毁token，用户还可以用通过token访问系统，本项目增加了缓存登出用户token到黑名单的方式，变相实现登出。
token续期问题，access_token携带有效期，有效期过了无法自动续期。本项目提供了续期接口（renewal），服务端在生成access_token同时还会生成refresh_token（有效期比access_token长），用户可以通过有效的refresh_token和access_token访问renewal接口重新获取新的refresh_token和access_token。

https://www.jb51.net/article/188548.htm

https://www.zhangshengrong.com/p/q0arAQ5k1x/

https://segmentfault.com/a/1190000022909135

https://github.com/apognu/gocas
https://github.com/aldaris/opensso

https://github.com/apereo/cas
