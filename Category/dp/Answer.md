[問題文](https://atcoder.jp/contests/abc211/tasks/abc211_c)

https://www.youtube.com/watch?v=Zu9S_kJ-7tk&ab_channel=AtCoderLive

#DP
#動的計画法

## 問題文の把握

- 文字列Sが与えられる
- 8文字を抽出しアンダーラインをひき、その文字が `c`, `h`, `o`, `k`, `u`, `d`, `a`, `i` となる組み合わせを求めよ
- 回答が大きくなる可能性があるので(10^9 +7)で割った値を出力せよ
制約

- 8 <= N <= 10^5

```
chchokudai
=> 3

chokudaichokudaichokudai
=> 45

```

## ロジック

仮に文字列を1文字づつ取り出して全ての組み合わせを試すとなると、O(8N^N) の計算量  
よって、何らかの高速化を行う必要がある  


### pseudo code


```
```

### tips

