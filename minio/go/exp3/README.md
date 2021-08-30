PresignedGetObject(bucketName, objectName string, expiry time.Duration, reqParams url.Values) (*url.URL, error)
生成一个用于HTTP GET操作的presigned URL。浏览器/移动客户端可以在即使存储桶为私有的情况下也可以通过这个URL进行下载。这个presigned URL可以有一个过期时间，默认是7天。



http://docs.minio.org.cn/docs/master/golang-client-api-reference#PresignedGetObject

PresignedPutObject(bucketName, objectName string, expiry time.Duration) (*url.URL, error)
生成一个用于HTTP GET操作的presigned URL。浏览器/移动客户端可以在即使存储桶为私有的情况下也可以通过这个URL进行下载。这个presigned URL可以有一个过期时间，默认是7天。

注意：你可以通过只指定对象名称上传到S3。

resignedHeadObject(bucketName, objectName string, expiry time.Duration, reqParams url.Values) (*url.URL, error)


PresignedPostPolicy(PostPolicy) (*url.URL, map[string]string, error)
允许给POST操作的presigned URL设置策略条件。这些策略包括比如，接收对象上传的存储桶名称，名称前缀，过期策略。


2.4 文件链接获取
getObjectURL()获取桶内文件的url地址，如果设置了只写权限，用户直接访问地址是查看不了的
getObgect()可以获取桶内对应文件的流对象，进行文件流传输
presignedGetObject()返回的是进行加密算法的地址，通过它可以直接访问文件



presignedGetObject 用于获取桶内文件的get下载链接
presignedPutObject 用于获取桶内文件的上传链接
presignedPostPolicy 用于设置返回和访问链接的编解码算法

https://www.freesion.com/article/23021407163/
