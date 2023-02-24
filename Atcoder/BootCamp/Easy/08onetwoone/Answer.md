[問題文](https://atcoder.jp/contests/abc086/tasks/abc086_b)

## 問題文の把握
- 整数 a, b を受け取る
- a, b を繋げた数が平方数かどうかを判定せよ
- 1 <= a,b <= 100

平方数とは、ある数を二乗した数値のこと  
また、二乗すると平方数になる値のことを平方根という  
ex. 4 は 2 の平方数であり、2 は 4 の平方根である

## ロジック

至ってシンプルな問題  
「入力値を受け取ったら文字列として連結し、数値に変換。  
その後変換した値の平方根を求め、整数になりうるなるなら Yes そうでなければ No と出力」  
する
また、平方根が整数になりうるかどうかは floor 関数を使うといい  
小数点以下が 0 でない時は値が繰り上がるため、floor 関数適用前後の値が同値の時は整数ということがわかる

### pseudo code

```
var a, b 

c = string a + string b
c conversion to int

root = c root
if root is int
    print(Yes)
else
    print(No)
```

### tips
