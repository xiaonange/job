1.简介
Docker 使用 Google 公司推出的 Go 语言 进行开发实现，基于 Linux 内核的 cgroup，namespace，以及 OverlayFS 类的 Union FS 等技术，
对进程进行封装隔离，属于 操作系统层面的虚拟化技术。由于隔离的进程独立于宿主和其它的隔离的进程，因此也称其为容器。最初实现是基于 LXC，
从 0.7 版本以后开始去除 LXC，转而使用自行开发的 libcontainer，从 1.11 开始，则进一步演进为使用 runC 和 containerd。

2.三个概念
镜像（Image）
容器（Container）
仓库（Repository）

3.常用命令
容器生命周期管理
run -d 启动守护模式
start/stop/restart  容器操作
kill 杀掉进程
rm  删除一个容器
pause/unpause
create 
exec
容器操作
ps
inspect
top
attach
events
logs
wait
export
port
容器rootfs命令
commit
cp
diff
镜像仓库
login
pull
push
search
本地镜像管理
images
rmi
tag
build
history
save
load
import
info|version
info
version

其他概念：
数据卷
网络

创建实名数据卷
docker volume create my-vol
创建匿名卷
docker volume create
查看数据卷列表
docker volume ls
查看具体的数据卷
docker volume inspect my-vol
删除数据卷
docker volume rm my-vol

 
