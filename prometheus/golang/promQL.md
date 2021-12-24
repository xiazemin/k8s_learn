https://www.qikqiak.com/k8strain/monitor/promql/

PromQL 是 Prometheus 内置的数据查询语言，其提供对时间序列数据丰富的查询，聚合以及逻辑运算能力的支持。并且被广泛应用在 Prometheus 的日常应用当中，包括对数据查询、可视化、告警处理。可以这么说，PromQL 是 Prometheus 所有应用场景的基础，理解和掌握 PromQL 是我们使用 Prometheus 必备的技能。


时间序列¶
前面我们通过 node-exporter 暴露的 metrics 服务，Prometheus 可以采集到当前主机所有监控指标的样本数据

# HELP node_cpu_seconds_total Seconds the cpus spent in each mode.
# TYPE node_cpu_seconds_total counter
node_cpu_seconds_total{cpu="0",mode="idle"} 6.62885731e+06
# HELP node_load1 1m load average.
# TYPE node_load1 gauge
node_load1 2.29

其中非 # 开头的每一行表示当前 node-exporter 采集到的一个监控样本：node_cpu_seconds_total 和 node_load1 表明了当前指标的名称、大括号中的标签则反映了当前样本的一些特征和维度、浮点数则是该监控样本的具体值。

Prometheus 会将所有采集到的样本数据以时间序列的方式保存在内存数据库中，并且定时保存到硬盘上。时间序列是按照时间戳和值的序列顺序存放的，我们称之为向量(vector)，每条时间序列通过指标名称(metrics name)和一组标签集(labelset)命名。

在时间序列中的每一个点称为一个样本（sample），样本由以下三部分组成：

指标(metric)：metric name 和描述当前样本特征的 labelsets
时间戳(timestamp)：一个精确到毫秒的时间戳
样本值(value)： 一个 float64 的浮点型数据表示当前样本的值


<--------------- metric ---------------------><-timestamp -><-value->
http_request_total{status="200", method="GET"}@1434417560938 => 94355


在形式上，所有的指标(Metric)都通过如下格式表示：


<metric name>{<label name> = <label value>, ...}
指标的名称(metric name)可以反映被监控样本的含义（比如，http_request_total - 表示当前系统接收到的 HTTP 请求总量）。指标名称只能由 ASCII 字符、数字、下划线以及冒号组成并必须符合正则表达式[a-zA-Z_:][a-zA-Z0-9_:]*。

标签(label)反映了当前样本的特征维度，通过这些维度 Prometheus 可以对样本数据进行过滤，聚合等。标签的名称只能由 ASCII 字符、数字以及下划线组成并满足正则表达式 [a-zA-Z_][a-zA-Z0-9_]*。

每个不同的 metric_name和 label 组合都称为时间序列，在 Prometheus 的表达式语言中，表达式或子表达式包括以下四种类型之一：


瞬时向量（Instant vector）：一组时间序列，每个时间序列包含单个样本，它们共享相同的时间戳。也就是说，表达式的返回值中只会包含该时间序列中的最新的一个样本值。而相应的这样的表达式称之为瞬时向量表达式。
区间向量（Range vector）：一组时间序列，每个时间序列包含一段时间范围内的样本数据，这些是通过将时间选择器附加到方括号中的瞬时向量（例如[5m]5分钟）而生成的。
标量（Scalar）：一个简单的数字浮点值。
字符串（String）：一个简单的字符串值。
所有这些指标都是 Prometheus 定期从 metrics 接口那里采集过来的。采集的间隔时间的设置由 prometheus.yaml 配置中的 scrape_interval 指定。最多抓取间隔为30秒，这意味着至少每30秒就会有一个带有新时间戳记录的新数据点，这个值可能会更改，也可能不会更改，但是每隔 scrape_interval 都会产生一个新的数据点。


Prometheus 定义了4种不同的指标类型：Counter（计数器）、Gauge（仪表盘）、Histogram（直方图）、Summary（摘要）。

在 node-exporter 返回的样本数据中，其注释中也包含了该样本的类型。例如：


# HELP node_cpu_seconds_total Seconds the cpus spent in each mode.
# TYPE node_cpu_seconds_total counter
node_cpu_seconds_total{cpu="cpu0",mode="idle"} 362812.7890625


PromQL 内置的聚合操作和函数可以让用户对这些数据进行进一步的分析，例如，通过 rate() 函数获取 HTTP 请求量的增长率：


rate(http_requests_total[5m])
查询当前系统中，访问量前 10 的 HTTP 请求：


topk(10, http_requests_total)


对于 Gauge 类型的监控指标，通过 PromQL 内置函数 delta() 可以获取样本在一段时间范围内的变化情况。例如，计算 CPU 温度在两个小时内的差异：


delta(cpu_temp_celsius{host="zeus"}[2h])
还可以直接使用 predict_linear() 对数据的变化趋势进行预测。例如，预测系统磁盘空间在4个小时之后的剩余情况：


predict_linear(node_filesystem_free_bytes[1h], 4 * 3600)


Summary 类型的指标相似之处在于 Histogram 类型的样本同样会反应当前指标的记录的总数(以 _count 作为后缀)以及其值的总量（以 _sum 作为后缀）。不同在于 Histogram 指标直接反应了在不同区间内样本的个数，区间通过标签 le 进行定义。



可以使用标签进行过滤查询，标签过滤器支持4种运算符：

= 等于
!= 不等于
=~ 匹配正则表达式
!~ 与正则表达式不匹配


标签过滤器都位于指标名称后面的{}内，比如过滤 master 节点的 CPU 使用数据可用如下查询语句：
node_cpu_seconds_total{instance="ydzs-master"}

此外我们还可以使用多个标签过滤器，以逗号分隔。多个标签过滤器之间是 AND 的关系，所以使用多个标签进行过滤，返回的指标数据必须和所有标签过滤器匹配。

例如如下查询语句将返回所有以 ydzs-为前缀的节点的并且是 idle 模式下面的节点 CPU 使用时长指标：


node_cpu_seconds_total{instance=~"ydzs-.*", mode="idle"}


我们可以通过将时间范围选择器（[]）附加到查询语句中，指定为每个返回的区间向量样本值中提取多长的时间范围。每个时间戳的值都是按时间倒序记录在时间序列中的，该值是从时间范围内的时间戳获取的对应的值。

时间范围通过数字来表示，单位可以使用以下其中之一的时间单位：

s - 秒
m - 分钟
h - 小时
d - 天
w - 周
y - 年


可以看到上面的两个时间序列都有4个值，这是因为我们 Prometheus 中配置的抓取间隔是15秒，所以，我们从图中的 @ 符号后面的时间戳可以看出，它们之间的间隔基本上就是15秒。

rate(): 计算整个时间范围内区间向量中时间序列的每秒平均增长率
irate(): 仅使用时间范围中的最后两个数据点来计算区间向量中时间序列的每秒平均增长率，irate 只能用于绘制快速变化的序列，在长期趋势分析或者告警中更推荐使用 rate 函数
increase(): 计算所选时间范围内时间序列的增量，它基本上是速率乘以时间范围选择器中的秒数

有的时候可能想要查看5分钟前或者昨天一天的区间内的样本数据，这个时候我们就需要用到位移操作了，位移操作的关键字是 offset，比如我们可以查询30分钟之前的 master 节点 CPU 的空闲指标数据：


node_cpu_seconds_total{instance="ydzs-master", mode="idle"} offset 30m

可以应用于多个时间序列或标量值的常规计算、比较和逻辑运算。

比如如下的两个瞬时向量：


node_cpu_seconds_total{instance="ydzs-master", cpu="0", mode="idle"}
和
node_cpu_seconds_total{instance="ydzs-node1", cpu="0", mode="idle"}
如果我们对这两个序列做加法运算来尝试获取 master 和 node1 节点的总的空闲 CPU 时长，则不会返回任何内容了：
这是因为这两个时间序列没有完全匹配标签。我们可以使用 on 关键字指定只希望在 mode 标签上进行匹配，就可以计算出结果来

node_cpu_seconds_total{instance="ydzs-master", cpu="0", mode="idle"}+on(mode)node_cpu_seconds_total{instance="ydzs-node1", cpu="0", mode="idle"}

如果我们真的想要获取节点的 CPU 总时长，我们完全不用这么操作，使用 sum 操作要简单得多：


sum(node_cpu_seconds_total{mode="idle"}) by (instance)

on 关键字只能用于一对一的匹配中，如果是多对一或者一对多的匹配情况下，就不行了
要解决这个问题，我们可以使用 group_left 或group_right 关键字。这两个关键字将匹配分别转换为多对一或一对多匹配。左侧和右侧表示基数较高的一侧。因此，group_left 意味着左侧的多个序列可以与右侧的单个序列匹配。结果是，返回的瞬时向量包含基数较高的一侧的所有标签，即使它们与右侧的任何标签都不匹配。

例如如下所示的查询语句就可以正常获取到结果，而且获取到的时间序列数据包含所有的标签:


container_cpu_user_seconds_total{namespace="kube-system"} * on (pod) group_left() kube_pod_info

瞬时向量和标量结合¶
此外我们还可以将瞬时向量和标量值相结合，这个很简单，就是简单的数学计算，比如：


node_cpu_seconds_total{instance="ydzs-master"} * 10
会为瞬时向量中每个序列的每个值都剩以10。这对于计算比率和百分比得时候非常有用。

除了 * 之外，其他常用的算数运算符当然也支持：+、-、*、/、%、^。
还有其他的比较运算符：==、!=、>、<、>=、<=。
逻辑运算符：and、or、unless，不过逻辑运算符只能用于瞬时向量之间。



https://www.bookstack.cn/read/prometheus-book/promql-prometheus-promql-functions.md

cpu_temperature_celsius*10
sum(cpu_temperature_celsius)
increase(hd_errors_total[1s])
rate(hd_errors_total[5s])



Legend format
使用名称或模式控制时间系列的名称，例如，{{hostname}}将替换为标签hostname的标签值
https://segmentfault.com/a/1190000016237454


据我所知,目前无法格式化Grafana中的图例(有一个开放的PR),但是当你使用Prometheus时,你可以使用它的label_replace()功能,例如:

label_replace(my_vector, "short_hostname", "$1", "hostname", "(.*):.*")
这应该给你:

Legend Format = {{short_hostname}}

Result = myhostname.mydomain.com

