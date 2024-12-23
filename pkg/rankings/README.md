## rankings


weight 的影响因子有 送礼金额、送礼时间、用户等级

衰减

我们的排行榜的需求经常变化，怎么实现只需要配置即可xxx

用factory pattern吗


比如说直播间礼物排行、文章热度排行、游戏战力排行等不同场景，具体场景的需求也经常变化

如果有一个排行榜，用 zset，根据积分和时间来排序，积分高的，时间最近的拍前面，怎么实现？

很好想到“直接用zset的score进行排序，积分作为高位，时间作为低位”这个方案，但是想到一些存在的问题，比如说

- 通常我们的需求都是类似 [redis zset 多字段排序](https://blog.51cto.com/u_16213628/7599769) 这个帖子所说的“积登录时间排序，updated_at越大，该用户排名越靠前”，这个是天然符合score排序的。但是如果需求变了，我们需要updated_at越小，排名越靠前怎么实现？
- timestamp时间戳位数太长了，13位嘛，怎么减少数据量？


*[Review代码思考：排行榜同积分按时间排序优化方案 | Lua开发实战 - 知乎](https://zhuanlan.zhihu.com/p/380545260)*

```markdown
① 方案一：通过时间差计算出方案； 分数 + 结束时间戳(固定位数)
② 方案二：通过进行位操作； 原理是 分数 + 结束时间戳(固定位数)
③ 方案三：基于雪花算法思想实现
1位高位不用 + 22位表示积分 + 41位时间戳
分数在高位，时间戳在低位，这样就可以保证不管时间戳是多少，分数越大，那么值就越大，也就符合我们需求

22位是符合我们业务需求的值：最大支持 (2^21-1)

41位时间戳最大支持毫秒级：2^40-1
```

三种方案的思路其实大同小异，

- 方案一的核心是通过当前时间和活动截止时间的时间差。然后拿时间差直接拼接积分，作为score
- 方案二是位运算，我不懂懒得看了

通过看这篇文章想到以下：

- 这类需求是确定的，不太存在需求变动的问题。所以问题1不存在。
- timestamp时间戳位数不是问题，因为本身zset的score作为double，64位，无论score的具体值是多少，都会以64位的double进行存储，所以不存在压缩位数来节省内存的操作。


---


[ranking/ranking.go at main · bryant12138/ranking](https://github.com/bryant12138/ranking/blob/main/ranking.go)

[Mace411/rank: 排行榜](https://github.com/Mace411/rank)

[interview-rank/rank.go at main · ComEyt/interview-rank](https://github.com/ComEyt/interview-rank/blob/main/rank.go)

[skiplist/skiplist.go at master · refine1017/skiplist](https://github.com/refine1017/skiplist/blob/master/skiplist.go)

[sirius2alpha/scoreboard: 使用Redis在服务器上对用户的点击数排序，并返回排行榜](https://github.com/sirius2alpha/scoreboard)

