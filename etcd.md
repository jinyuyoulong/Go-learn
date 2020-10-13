[TOC]

### 描述

etcd是CoreOS团队于2013年6月发起的开源项目，它的目标是构建一个高可用的分布式键值(key-value)数据库。etcd内部采用raft协议作为一致性算法，etcd基于Go语言实现。

etcd作为服务发现系统，有以下的特点

- 简单：安装配置简单，而且提供了HTTP API进行交互，使用也很简单
- 安全：支持SSL证书验证
- 快速：根据官方提供的benchmark数据，单实例支持每秒2k+读操作
- 可靠：采用raft算法，实现分布式系统数据的可用性和一致性
- etcd项目地址：https://github.com/coreos/etcd/

### 下载

```sh
$ docker search etcd
$ docker pull xxx/etcd
```

## ETCD参数说明

- **data-dir:**指定节点的数据存储目录，这些数据包括节点ID，集群ID，集群初始化配置，Snapshot文件，若未指定—wal-dir，还会存储WAL文件；
- **wal-dir:**指定节点的was文件的存储目录，若指定了该参数，wal文件会和其他数据文件分开存储。
- **name:** 节点名称
- **initial-advertise-peer-urls:** 告知集群其他节点url.(`对于集群内提供服务的url`)
- **listen-peer-urls:** 监听URL，用于与其他节点通讯
- **advertise-client-urls:** 告知客户端url, 也就是服务的url(`对外提供服务的utl`)
- **initial-cluster-token:** 集群的ID
- **initial-cluster:** 集群中所有节点

然后进入其中一个 Docker 主机：

```sh
$ docker exec -it etcd bin/sh
```

执行下面命令（查看集群成员）：

```sh
$ etcdctl member list
```