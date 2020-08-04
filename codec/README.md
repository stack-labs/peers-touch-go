# Codec

## Design

### Peer

Connected two Peers codec each other's data with different codec implements.

### 设计

### 端点解码

端点与端点之间的解码使用不同的编（解）码器，这是为了保证每个连接了的点之间通信在解码层面是特定的，在默认情况，使用同一JSON编码器已足够。但是对于有安全加密需要两点之前，往往彼此在创建连接前就已经声明好加密密钥、方式等。

# Acknowledgements

1. Copy some design of Go-Micro





















