# 基于 Bitcask 模型，兼容 Redis 数据结构和协议的高性能 KV 存储引擎 设计细节

1. 采用 Key/Value 的数据模型，实现数据存储和检索的快速、稳定、高效
2. 存储模型：采用 Bitcask 存储模型，具备高吞吐量和低读写放大的特征
3. 持久化：实现了数据的持久化，确保数据的可靠性和可恢复性
4. 索引：多种内存索引结构，高效、快速数据访问
5. 并发控制：使用锁机制，确保数据的一致性和并发访问的正确性
6. 编程语言：采用 Go 编写，兼顾高性能以及编码简洁性