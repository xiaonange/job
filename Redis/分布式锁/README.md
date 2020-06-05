目前分布式锁可归并为两种
1.基于单Redis节点的分布式锁
单节点redis
允许锁的偶尔失效，那么使用单Redis节点的锁方案就足够了，简单而且效率高

2.多节点分布式锁的算法Redlock
redis 之父提出，不完善实际还是要用
N 个完全独立的 Redis 节点，循环写入
增加了token，谁创建谁销毁


核心实现办法：
setnx  单线程互斥性
get  获取当前时间
getset 获取单前时间，并设置新的时间
expire 设置键有效期
ttl  获取当前键剩下有效时间 
del 删除缓存

文档参考引用：基于Redis的分布式锁到底安全吗（上）？（https://mp.weixin.qq.com/s/JTsJCDuasgIJ0j95K8Ay8w）
             基于Redis的分布式锁到底安全吗（下）（ https://cloud.tencent.com/developer/article/1418246）
             How to do distributed locking（https://martin.kleppmann.com/2016/02/08/how-to-do-distributed-locking.html）

