为了在兼容旧版本的同时不断升级新的API，Kubernetes支持多种API版本，每种API版本都有不同的API路径，例如/api/v1或 /apis/extensions/v1beta1

Alpha级别：Kubernetes中新功能的早期候选者。这些可能包含错误，并且不能保证将来能正常工作
Beta级别：在API版本名称意味着测试已取得进展过去的Alpha级别水平，并且该功能最终将被列入Kubernetes“测试”。尽管它的工作方式可能会发生变化，并且对象的定义方式可能会完全发生变化，但功能本身很有可能以某种形式将其纳入Kubernetes
Stable级别：稳定版本，该版本名称命名方式：vX这里X是一个整数

https://blog.csdn.net/yangshihuz/article/details/112648675

如Deployment:

1.6版本之前 apiVsersion：extensions/v1beta1
1.6版本到1.9版本之间：apps/v1beta1
1.9版本之后:apps/v1

https://blog.csdn.net/xixihahalelehehe/article/details/112789211

apps是指应用的场景，例如可以使用deployment，service，namespace等，几乎涵盖了申明资源，不过也可以不写，进行通用匹配资源，但也有些资源需要特殊申明，具体可以通过命令查看。

https://www.cnblogs.com/zhouzhifei/p/12082909.html