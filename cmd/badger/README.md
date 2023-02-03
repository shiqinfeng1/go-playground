# 高性能的K/V数据库
- LevelDB 也是由Google的牛人 Jeffrey Dean 和 Sanjay Ghemawat创建的，被多个NoSql数据库用作底层的存储引擎。 
- RocksDB fork自LevelDB，但为多核和SSD做了很多的优化， 增加了一些有用的特性，比如Bloom filters、事务、TTL等，被Facebook、Yahoo!和Linkedin等公司使用。
- badger 为了更好地性能，专门为SSD所优化，将key和value分别存储以减少I/O放大。
- 随机读，Badger至少要比RocksDB快3.5倍，对于值的大小从128B到16KB，数据加载Badger比LevelDB快0.86 - 14倍。
- Badger分离的key和value,只有key存在LSM tree中， value存在WAL中，叫做value log。