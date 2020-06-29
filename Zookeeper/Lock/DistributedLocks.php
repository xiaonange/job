<?php

class zkCli {
    protected static $zk;
    protected static $myNode;
    protected static $isNotifyed;
    protected static $root;

    public static function getZkInstance($conf, $root){
        try{

            if(isset(self::$zk)){
                return self::$zk;
            }

            $zk = new \Zookeeper($conf['host'] . ':' . $conf['port']);
            if(!$zk){
                throw new \Exception('connect zookeeper error');
            }

            self::$zk = $zk;
            self::$root = $root;

            return $zk;
        } catch (\ZookeeperException $e){
            die($e->getMessage());
        } catch (\Exception $e){
            die($e->getMessage());
        }
    }

    // 获取锁
    public static function tryGetDistributedLock($lockKey, $value){
        try{
            // 创建根节点
            self::createRootPath($value);
            // 创建临时顺序节点
            self::createSubPath(self::$root . $lockKey, $value);
            // 获取锁
            return self::getLock();

        } catch (\ZookeeperException $e){
            return false;
        } catch (\Exception $e){
            return false;
        }
    }

    // 释放锁
    public static function releaseDistributedLock(){
        if(self::$zk->delete(self::$myNode)){
            return true;
        }else{
            return false;
        }
    }

    public static function createRootPath($value){
        $aclArray = [
            [
                'perms'  => Zookeeper::PERM_ALL,
                'scheme' => 'world',
                'id'     => 'anyone',
            ]
        ];
        // 判断根节点是否存在
        if(false == self::$zk->exists(self::$root)){
            // 创建根节点
            $result = self::$zk->create(self::$root, $value, $aclArray);
            if(false == $result){
                throw new \Exception('create '.self::$root.' fail');
            }
        }

        return true;
    }

    public static function createSubPath($path, $value){
        // 全部权限
        $aclArray = [
            [
                'perms'  => Zookeeper::PERM_ALL,
                'scheme' => 'world',
                'id'     => 'anyone',
            ]
        ];
        /**
         * flags :
         * 0 和 null 永久节点，
         * Zookeeper::EPHEMERAL临时，
         * Zookeeper::SEQUENCE顺序，
         * Zookeeper::EPHEMERAL | Zookeeper::SEQUENCE 临时顺序
         */
        self::$myNode = self::$zk->create($path, $value, $aclArray, Zookeeper::EPHEMERAL | Zookeeper::SEQUENCE);
        if(false == self::$myNode){
            throw new \Exception('create -s -e '.$path.' fail');
        }
        echo 'my node is ' . self::$myNode.'-----------'.PHP_EOL;

        return true;
    }

    public function getLock(){
        // 获取子节点列表从小到大，显然不可能为空，至少有一个节点
        $res = self::checkMyNodeOrBefore();
        if($res === true){
            return true;
        }else{
            self::$isNotifyed = false;// 初始化状态值
            // 考虑监听失败的情况：当我正要监听before之前，它被清除了，监听失败返回 false
            $result = self::$zk->get($res, [zkCli::class, 'watcher']);
            while(!$result){
                $res1 = self::checkMyNodeOrBefore();
                if($res1 === true){
                    return true;
                }else{
                    $result = self::$zk->get($res1, [zkCli::class, 'watcher']);
                }
            }

            // 阻塞，等待watcher被执行，watcher执行完回到这里
            while(!self::$isNotifyed){
                echo '.';
                usleep(500000); // 500ms
            }

            return true;
        }
    }

    /**
     * 通知回调处理
     * @param $type 变化类型 Zookeeper::CREATED_EVENT, Zookeeper::DELETED_EVENT, Zookeeper::CHANGED_EVENT
     * @param $state
     * @param $key 监听的path
     */
    public static function watcher($type, $state, $key){
        echo PHP_EOL.$key.' notifyed ....'.PHP_EOL;
        self::$isNotifyed = true;
        self::getLock();
    }

    public static function checkMyNodeOrBefore(){
        $list = self::$zk->getChildren(self::$root);
        sort($list);
        $root = self::$root;
        array_walk($list, function(&$val) use ($root){
            $val = $root . '/' . $val;
        });

        if($list[0] == self::$myNode){
            echo 'get locak node '.self::$myNode.'....'.PHP_EOL;
            return true;
        }else{
            // 找到上一个节点
            $index = array_search(self::$myNode, $list);
            $before = $list[$index - 1];
            echo 'before node '.$before.'.........'.PHP_EOL;
            return $before;
        }
    }
}


function zkLock($resourceId){
    $conf = ['host'=>'127.0.0.1', 'port'=>2181];
    $root = '/lockKey_' . $resourceId;
    $lockKey = '/lock_';
    $value = 'a';

    $client = zkCli::getZkInstance($conf, $root);
    $re = zkCli::tryGetDistributedLock($lockKey, $value);

    if($re){
        echo 'get lock success'.PHP_EOL;
    }else{
        echo 'get lock fail'.PHP_EOL;
        return ;
    }

    try {

        doSomething();

    } catch(\Exception $e) {

        echo $e->getMessage() . PHP_EOL;

    } finally {

        $re = zkCli::releaseDistributedLock();
        if($re){
            echo 'release lock success'.PHP_EOL;
        }else{
            echo 'release lock fail'.PHP_EOL;
        }

        return ;
    }
}

function doSomething(){
    $n = rand(1, 20);
    switch($n){
        case 1:
            sleep(15);// 模拟超时
            break;
        case 2:
            throw new \Exception('system throw message...');// 模拟程序中止
            break;
        case 3:
            die('system crashed...');// 模拟程序崩溃
            break;
        default:
            sleep(13);// 正常处理过程
    }
}

// 执行
zkLock(0);
