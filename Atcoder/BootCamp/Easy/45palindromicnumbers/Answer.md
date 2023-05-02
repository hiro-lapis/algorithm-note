[問題文](https://atcoder.jp/contests/abc090/tasks/abc090_b)

#文字列

## 問題文の把握

- A~B の整数のうち、回文となる数の個数を求めよ
回文とは、前から読んでも後ろから読んでも同じ読みになる文字列のこと

- 制約
- 100000 <= A <= B <= 99999


## ロジック

まともに計算するとしたら以下のようなロジックになる

```
var a, b int
ans
for a to b
    tmp int
    if checkTmpIsKaibun
        ans++

print(ans)

checkTmpIsKaibun
    return tmp[0] == tmp[4] && tmp[1] == tmp[3]
```

計算にかかる時間は O(n) 数値の範囲から10万回におさまる
A B の数値の範囲を元に解答値を一発で出すロジックもありそうだが、
家庭用PCでの1s計算量が 10 億回なので、間に合う想定でこのロジックでいく

10億 => 10^8

### pseudo code


### tips

- 解答時間 10m
