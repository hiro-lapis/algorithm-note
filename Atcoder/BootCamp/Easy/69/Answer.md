[問題文](https://atcoder.jp/contests/agc037/tasks/agc037_a)

#dp
#文字列


## 問題文の把握

- 英小文字Sが与えられる
- Sを空でない K 個のユニークな文字列に分割した場合の最大数を求めよ
- 分割した文字を Si とすると、Si ≠ Si+1 となる(1 <= i <= K-1)
- Si, Si+1,... Sk を結合するとSになる

- 制約
- 

```
case1

aabbaa
=> 4
a, ab, ba, a = 4
aa, b, ba, a = 4
a, ab, b, aa = 4
a, abb, aa = 3

case2
aaaccacabaababc
=> 12

aaa, cc, ac, ab, aa, ba, b, c = 8
a, aa, cc, ac, ab, a,a, ba, b, c = 8

```

## ロジック

文字列の出現回数を


### pseudo code


```
```

### tips

