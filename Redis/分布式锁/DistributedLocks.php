<?php




namespace warehouse\publicfun;
use common\PublicFunc;
use Common\RedisService;
use Xaircraft\Exception\ExceptionHelper;

class DistributedLocks
{

    private $retryDelay = 200; //重试延迟 毫秒
    private $retryCount = 1;  //重试数
    private $servers = array();

    function __construct( )
    {
        $this->servers = RedisService::getRedis();
        //var_dump($this->servers->info());
    }

    public static function Create()
    {
        return new  self();
    }

    public function lock($lockKey, $ttl)
    {
        $token = uniqid();
        $retry = $this->retryCount;
        do {
            if ($this->lockInstance($lockKey, $token, $ttl)) {
                return [
                    'lockKey' => $lockKey,
                    'token' => $token,
                ];
            }
            // 等待随机重试
            $delay = mt_rand(floor($this->retryDelay / 2), $this->retryDelay);
            usleep($delay * 1000);
            $retry--;
        } while ($retry > 0);
       ExceptionHelper::ThrowIfNullOrEmpty('','加锁错误请稍等');
    }

    /**
     * 解锁
     * @param $lockKey
     * @param $token
     */
    public function unlock($lockKey)
    {
        return  $this->servers->del($lockKey);

    }


    /**
     * 上锁
     * @param $lockKey
     * @param $token
     * @param $lockValue
     * @return mixed
     */
    private function lockInstance($lockKey, $token, $lockValue)
    {
        $lockValue = time() + $lockValue;
        $lock =$this->servers->setnx($lockKey, $lockValue);
        /**
         * 满足两个条件中的一个即可进行操作
         * 1、上面一步创建锁成功;
         * 2、   1）判断锁的值（时间戳）是否小于当前时间    $redis->get()
         *      2）同时给锁设置新值成功    $redis->getset()
         */
        if (!empty($lock) || ($this->servers->get($lockKey) < time() && $this->servers->getSet($lockKey, $lockValue) < time())){
            $this->servers->expire($lockKey,$lockValue);
            return true;
        }
       return false;
    }

}