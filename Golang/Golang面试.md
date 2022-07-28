1.new 和make的差别
 答：new返回的是一个类型分配的内存的指针，make用来对map,切片，channel的分配和初始化

2.go内存管理
答：垃圾回收：三色并发标记法，避免了标记法的卡顿

3.函数变量传递用指针还是用值
答：看场景，需要改变上层的值，传指针，减少拷贝。只是单纯参数，正常传变量。指针传递可能引起内存逃逸

4.go的内存逃逸
答：内存逃逸指的是栈上的分配空间不够，或者是原本在栈上申请的空间变为在推上申请。
   1. 指针逃逸  函数返回结果，变量还复用
   2.申请的栈太大
   3.动态类型逃逸 interface
   4.闭包引用对象逃逸
   逃逸分析可以减少gc操作，栈上分配内存比在堆中分配内存有更高的效率，栈上分配的内存不需要GC处理。
   逃逸分析命令：go build -gcflags '-m'

5.go的协程，调度，抢占
答：GMP模型的M就是线程，M默认根据cpu数量限定,内核线程。P processor 调度器，G 协程创建G协程，
    相比java的多线程，开销小，更轻量级。io操作，等待channel,系统调用等待回调，G状态会变化,挂起等待，M继续接受其他的协程G
    通过信号协作的方式
    goroutine抢占时机，只有在垃圾回收的时候和栈扫描的时候

6.go协程有哪些状态
答：Gidle（ˈaɪdl）：表示刚刚分配了Goroutine内存但没有进行初始化，此时状态为空闲状态
    Grunnable（ /'rʌnəbl/ ）：可运行状态，在p的本地队列中等待执行。
    Grunning：代表正在执行程序指令，即Goroutine正在运行，此时Goroutine不在p运行队列中，并取得m、p运行资源，即与m和p完成绑定关系，Goroutine完全使用所在m（线程）的栈空间。
    Gsyscall：也说明Goroutine，但不是执行用户层代码，也是正在执行系统调用，此时Goroutine不在p运行队列中，完全拥有m的栈空间。
    Gwaiting：此时Goroutine没有执行用户代码，即不在p运行队列，只是某个时刻打个标
    Gdead：数量减少时多余的P会变为此状态,类似休眠。协程处理完毕会挂起来待用
    schedule保证了不会退出
    deadlock会进行死锁检测
   参考资料： https://blog.csdn.net/QQ1130141391/article/details/96350019

7.go线程和协程大小
答：1.5版本 线程2M，协程是2KB

8.linux有哪几种线程模型
答：1.多对一(M:1)的用户级线程模型
    2.一对一(1:1)的内核级线程模型
    3.多对多(M:N)的两级线程模型

9.go的线程状态
答：自旋和非自旋

10.go线程oom  内存泄漏
答：一般是代碼邏輯問題，使用pprof，或者火焰圖來查證分析處理，可能原因有：
   1.频繁申请重复对象
   2.并发大，Goroutine 数过多，GC 压力增大，GC 缓慢（1.2以後的go的內存回收不會立即釋放，需要5分鐘）
   3.部署環境導致的內存（linux內核版本，go版本）
   參考資料：https://studygolang.com/articles/29149?fr=sidebar

 11.go裡面的錯誤處理
 答：常規還是一行代碼，三行判斷。也可以用defer和 recover 來實現其他語言的try catch

12.Go实现的互斥锁有两种模式，分别是正常模式和饥饿模式
答：為了防止最後面的wait的Goroutine一直處於等待狀態，当waiter超过 1ms 没有获取到锁，它就会将当前互斥锁切换到饥饿模式，防止等待队列中的waiter被饿死。

13.gin的參數怎麼校驗
答：通過構造一個結構體，定義tag標籤，通過反射來做解析

14.go的interface
答：底層是一個有函數的iface和无函数的eface

15：go的反射實現
答：reflect有兩個結構體一個是type,一個是value.通過reflect.typeOf,reflect.valueOf去獲取對應的值

16.go怎麼通過字符串調用辦法
答：初始化含有辦法的結構體，通過reflect.ValueOf(data).MethodByName("").Call([]reflect.Value{})

17.go的鎖有哪幾種模式？對他的理解
答：sync.Mutex 互斥鎖，sync.RWMutex 讀寫鎖
    1.互斥鎖：兩種操作：获取锁和释放锁,兩種模式：正常模式和饥饿模式,為了防止最後面的wait的Goroutine一直處於等待狀態，当waiter超过 1ms 没有获取到锁，它就会将当前互斥锁切换到饥饿模式
        ，防止等待队列中的waiter被饿死。
     用互斥鎖實現的官方包：sync.once
    2.讀寫鎖：四種操作：读上锁 读解锁 写上锁
      用读写锁实现的办法:sync.map
    3.其他：sync.waitgroup裡面定義的辦法和互斥鎖一樣有兩個操作：Lock，unlock,它的Add操作依靠的是原子鎖atomic.AddUint64,輕量級的原子鎖

18：go的channel的用法，對他的理解
答：channel有两种：一种是无缓冲，一种是有缓冲。无缓冲有写入就堵塞等待读取，有缓冲，在缓冲数量未写满前，都可以继续写入，缓冲数量满了再堵塞等待。写入已关闭的channel将报panic
    channel是线程安全的

 19 .runtime

20.协程G数据结构
struct G
{
    uintptr    stackguard;    // 分段栈的可用空间下界
    uintptr    stackbase;    // 分段栈的栈基址
    Gobuf    sched;        //进程切换时，利用sched域来保存上下文
    uintptr    stack0;
    FuncVal*    fnstart;        // goroutine运行的函数
    void*    param;        // 用于传递参数，睡眠时其它goroutine设置param，唤醒时此goroutine可以获取
    int16    status;        // 状态    Gidle,Grunnable,Grunning,Gsyscall,Gwaiting,Gdead
    int64    goid;        // goroutine的id号
    G*    schedlink;
    M*    m;        // for debuggers, but offset not hard-coded
    M*    lockedm;    // G被锁定只能在这个m上运行
    uintptr    gopc;    // 创建这个goroutine的go表达式的pc
...
};

struct M
{
    G*    g0;        // 带有调度栈的goroutine
    G*    gsignal;    // signal-handling G 处理信号的goroutine
    void    (*mstartfn)(void);
    G*    curg;        // M中当前运行的goroutine
    P*    p;        // 关联P以执行Go代码 (如果没有执行Go代码则P为nil)
    P*    nextp;
    int32    id;
    int32    mallocing; //状态
    int32    throwing;
    int32    gcing;
    int32    locks;
    int32    helpgc;        //不为0表示此m在做帮忙gc。helpgc等于n只是一个编号
    bool    blockingsyscall;
    bool    spinning;
    Note    park;
    M*    alllink;    // 这个域用于链接allm
    M*    schedlink;
    MCache    *mcache;
    G*    lockedg;
    M*    nextwaitm;    // next M waiting for lock
    GCStats    gcstats;
...
};
struct P
{
    Lock;
    uint32    status;  // Pidle或Prunning等
    P*    link;
    uint32    schedtick;  // 每次调度时将它加一
    M*    m;    // 链接到它关联的M (nil if idle)
    MCache*    mcache;
    G*    runq[256];
    int32    runqhead;
    int32    runqtail;
    // Available G's (status == Gdead)
    G*    gfree;
    int32    gfreecnt;
    byte    pad[64];
};


21.channel 结构体
type hchan struct {
    qcount   uint   // channel 里的元素计数
    dataqsiz uint   // 可以缓冲的数量，如 ch := make(chan int, 10)。 此处的 10 即 dataqsiz
    elemsize uint16 // 要发送或接收的数据类型大小
    buf      unsafe.Pointer // 当 channel 设置了缓冲数量时，该 buf 指向一个存储缓冲数据的区域，该区域是一个循环队列的数据结构
    closed   uint32 // 关闭状态
    sendx    uint  // 当 channel 设置了缓冲数量时，数据区域即循环队列此时已发送数据的索引位置
    recvx    uint  // 当 channel 设置了缓冲数量时，数据区域即循环队列此时已接收数据的索引位置
    recvq    waitq // 想读取数据但又被阻塞住的 goroutine 队列
    sendq    waitq // 想发送数据但又被阻塞住的 goroutine 队列

    lock mutex
    ...
}




