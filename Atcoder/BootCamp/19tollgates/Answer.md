[問題文](https://atcoder.jp/contests/abc094/tasks/abc094_b)

## 問題文の把握

- N+1 のマスが左右に一列で並んでいる
- マスにはそれぞれ 0,1...N と番号が振られている
- 最初に X マスにいる
  - 隣あうマスには cost 0 で移動できる
  - マス 0, N に着くとゴールになる
  - i = 1,2....M について、Ai には料金所がある。料金所 Ai に移動するには cost = 1がかかる
  - マス 0, N には料金所がないことが保証されている
- ゴールまでにかかる cost 最小値を求めよ

- 制約
- 1 <= N,M <= 100
- 1 <= X <= N-1
- 1 <= A1 < A2 ... Am <= N-1
- Ai ≠ X

```
N M X
A1 A2 AM

5 3 3
1 2 4
=> 1
```

・料金所と自分の初期配置されるマスは重ならない
・M の数だけ Ai の入力値が増える
・A1, AM の入力値は、移動cost がかかる key の情報

## ロジック

最小値を求める場合、初期配置 index よりも大きい・小さい Ai の数を count し、
小さい方を解答として出力すればいい




### pseudo code


```
var n m x
var aList
for i to m
    tmp = stdIn 
    aList[i] = tmp
    
before = count under x
after = count above x

ans = min(before, after)
print(ans)
```

### tips

