# job-
面试问题汇总
##Web
CSRF：token校验
sql注入

##PHP
opcache 字节缓存
ob_start ob系列函数 静态文件生成
设计模式：工厂模式，单例模式（instanceof ：数据库）

##Mysql
表级
行级锁 
乐观锁
高并发表数据的一致性
##Redis
分布式锁  锁的粒子

缓存雪崩  同时间失效，热点数据永不失效
缓存穿透  恶意请求，做好过滤，token校验
缓存击穿  单个，加锁写缓存即可
##Docket
常用命令

##kafaka
为什么选这个，为什么能处理这么多数据


##linux
crontab  定时器


##nginx
负载均衡
轮训方式：


##Golang
常用命令
性能测试工具

##计算机理论概念
进程拥有自己独立的堆和栈，既不共享堆，亦不共享栈，进程由操作系统调度。
线程拥有自己独立的栈和共享的堆，共享堆，不共享栈，线程也由操作系统调度。
协程和线程一样共享堆，不共享栈，协程由程序员在代码里调度