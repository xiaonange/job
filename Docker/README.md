1.简介
Docker 使用 Google 公司推出的 Go 语言 进行开发实现，基于 Linux 内核的 cgroup，namespace，以及 OverlayFS 类的 Union FS 等技术，
对进程进行封装隔离，属于 操作系统层面的虚拟化技术。由于隔离的进程独立于宿主和其它的隔离的进程，因此也称其为容器。最初实现是基于 LXC，
从 0.7 版本以后开始去除 LXC，转而使用自行开发的 libcontainer，从 1.11 开始，则进一步演进为使用 runC 和 containerd。

2.三个概念
镜像（Image）
容器（Container）
仓库（Repository）

3.常用命令
docker build -t   源目录 name:version  
docker  run -d/守护模式  -p 800:80
docker push 上传
docker pull 下拉
docker login 
docker images 查看当前的镜像列表
docker ps  查看当前运行的docker容器
docker stop/start/restart   name/仓库id


其他概念：
数据卷
网络

 
