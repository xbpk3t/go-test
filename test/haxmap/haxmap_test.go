package haxmap

// - url: https://github.com/orcaman/concurrent-map
// des: 怎么实现 concurrent-map? (怎么用map实现并发读写? 比较 "mutex+map" 和 sync.Map) 默认使用 32个shards，分片并发，减少锁粒度，也就是把一把锁分成几把，每把锁控制一个分片。 # [一个线程安全的泛型支持map库](https://colobu.com/2022/09/04/a-thread-safe-and-generic-supported-map/)
// rel:
// - url: https://github.com/cornelk/hashmap
// - url: https://github.com/alphadose/haxmap
// des: 相比于concurrent-map，key和value都支持generics
//
//
// - url: https://github.com/elliotchance/orderedmap
// des: 怎么实现 map 的有序查找，且支持 add、支持 delete、支持迭代？构造一个辅助 slice
// rel:
// - url: https://github.com/iancoleman/orderedmap
// des: Use slice as linear in struct.
//
// - url: https://github.com/chainbound/shardmap
// des: 怎么优化map? shardmap = HashMap + generics # implement hashmap using generics. Can be used as a reference to implement hashmap.
