# Peers-Touch-Go

P2P communication lib in Go and based on IPFS. It supports Sub/Pub msg and file.   

Why I create this?

We need to build a world without AI, BigData. Human in this world respect each other's private info. No man analyze our data in the place we don't know.

No one can pry into our personal life and trace any private things.

So now you know, I hate the big companies hold all data of ours. 

## Structure

Inspired by Go-Micro and Orbit-db.

![architecture](./doc/images/architecture.png)

```
--Log
--Peer
----Codec
----Pubsub
----Registry
----Store
----Transport
-------------Touch
------------------Node
```

### Node

Node is the smallest indivisible unit for exchanging data in Peers. It includes `Transport`, `Broker`, `Store`, `Registry` for the outside and `Config`, `Logger` for the inside.

# 中文

基于P2P传输协议（IPFS）的远程通信框架，提供去中心化的通信组件，支持消息订阅、发布，支持远程文件传输。

### Node 类型

- 节点类型
  - 中继节点：网关，公网节点，用于NAT节点之间的通信中继，不干涉与不猜测消息内容，透明但有条件（如安全、验证等）过桥传输
  - 注册节点：公网节点，用于NAT节点的与中继节点的注册
  - 终端节点：不限网络的节点，数据的来源与最终去向

部分参考Go-Micro及Orbit-db的设计

