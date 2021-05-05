# current-limit

使用go语言实现各种限流算法

### 限流类型
代码实现基于请求进行限流(限制QPS等),基于资源限流等类型在原有代码上修改即可

### 限流粒度
- 单机版
- 集群版(基于redis实现)

### 实现
#### 固定窗口限流器 
- 规定时间范围判断次数 [实现1](./window/util/fixation1.go)
- 规定次数范围判断时间 [实现2](./window/util/fixation2.go)


#### 滑动窗口限流器
- 单机版本 [实现](./window/util/slide.go)
- 集群版本 [实现](./window/util/slide_redis.go)

### 外部包引用
- phachon/go-logger 用来做日志实现 
- github.com/garyburd/redigo/redis 实现redis操作

