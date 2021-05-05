# current-limit

使用go语言实现各种限流算法

### 限流类型
代码实现基于请求进行限流(限制QPS等),基于资源限流等类型在原有代码上修改即可

### 限流粒度
- 单机版
- 集群版(基于redis实现)

### 外部包引用
- phachon/go-logger 用来做日志实现 
- github.com/garyburd/redigo/redis 实现redis操作

