https://tangxusc.github.io/blog/2019/05/code-generator%E4%BD%BF%E7%94%A8/

code-generator提供了以下工具为kubernetes中的资源生成代码:

deepcopy-gen: 生成深度拷贝方法,避免性能开销

client-gen:为资源生成标准的操作方法(get,list,create,update,patch,delete,deleteCollection,watch)

informer-gen: 生成informer,提供事件机制来相应kubernetes的event

lister-gen: 为get和list方法提供只读缓存层

其中informer和listers是构建controller的基础,kubebuilder也是基于informer的机制生成的代码.

code-generator还专门整合了这些gen,形成了generate-groups.sh和generate-internal-groups.sh这两个脚本.

https://segmentfault.com/a/1190000039706356

https://segmentfault.com/a/1190000023097945

,在我们的源代码中出现了很多

doc.go

// +k8s:deepcopy-gen=package,register
// +groupName=samplecontroller.k8s.io
types.go

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
出现了这样的tag,这些tag到底是什么意思呢,有什么作用呢?

分类
其实code-generator将tag分为了两种,

Global tags: 全局的tag,放在具体版本的doc.go文件中

Local tags: 本地的tag,放在types.go文件中的具体的struct上.

tag的使用语法为:

// +tag-name 
或
// +tag-name=value
并且 这些注释块必须分开,这也是源代码中 注释存在分割的原因. 

Global
全局的tag是写在doc.go中的,典型的内容如下:

// +k8s:deepcopy-gen=package


// Package v1 is the v1 version of the API.
// +groupName=example.com
package v1
注意: 空行不能省

+k8s:deepcopy-gen=: 它告诉deepcopy默认为该包中的每一个类型创建deepcopy方法

如果不需要深度复制,可以选择关闭此功能// +k8s:deepcopy-gen=false

如果不启用包级别的深度复制,那么就需要在每个类型上加入深度复制// +k8s:deepcopy-gen=true

+groupName: 定义group的名称,注意别弄错了.

注意 这里是 +k8s:deepcopy-gen=,最后是 = ,和local中的区别开来.

local
本地的tag直接写在类型上,典型的内容如下:

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FooSpec   `json:"spec"`
	Status FooStatus `json:"status"`
}
可以看到local支持两种tag

+genclient: 此标签是告诉client-gen,为此类型创建clientset,但也有以下几种用法.

1,对于集群范围内的资源(没有namespace限制的),需要使用// +genclient:nonNamespaced,生成的clientset中的namespace()方法就不再需要传入参数

2,使用子资源分离的,例如/status分离的,则需要使用+genclient:noStatus,来避免更新到status资源(当然代码的struct中也没有status)

3,对于其他的值,这里不做过多的解释,请参照

// +genclient:noVerbs
// +genclient:onlyVerbs=create,delete
// +genclient:skipVerbs=get,list,create,update,patch,delete,deleteCollection,watch
// +genclient:method=Create,verb=create,result=k8s.io/apimachinery/pkg/apis/meta/v1.Status
+k8s:deepcopy-gen:interfaces=: 为struct生成实现 tag值的DeepCopyXxx方法

例如:// +k8s:deepcopy-gen:interfaces=example.com/pkg/apis/example.SomeInterface

将生成 DeepCopySomeInterface() SomeInterface方法


http://bingerambo.com/posts/2021/05/k8s%E4%B8%AD%E7%9A%84crd%E5%BC%80%E5%8F%91/

https://blog.csdn.net/u010918487/article/details/104961052


https://cloud.redhat.com/blog/kubernetes-deep-dive-code-generation-customresources


https://blog.crazytaxii.com/posts/k8s_code_generation_4_customresources/

https://www.xieys.club/code-generator-crd/


https://juejin.cn/post/6844903783298777096
