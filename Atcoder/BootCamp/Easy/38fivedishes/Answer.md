[問題文](https://atcoder.jp/contests/abc123/tasks/abc123_b)

#tag
#考察力


## 問題文の把握

- 5つの料理を注文する。それぞれの料理の調理にかかる時間はa, b, c, d, e である
- それぞれの料理に取り掛かる時間は、10の倍数の時間である
  - 最初に調理する料理が8mかかる場合 => 0 ~ 8m の間で調理。2番目の料理は10mから調理開始
  - 2番目に調理する料理が30mに完成し届いた場合 => 30m ジャストから次の料理を開始
- 調理開始は 0m　からである
- ある料理を注文したら、その料理が届くまで別の注文はできない
- 料理を注文する順番を任意とした場合、最短で全ての料理を注文するオーダー順を求めよ

- 制約
- A,B,C,D,E は 1~123 以下の整数

```
case1
29
20
7
35
120
=>215

case2
101
86
119
108
57
=>481

case3
123
123
123
123
123
=>643
```

## ロジック

case を見ると、最後の注文の届く時刻だけは 10の倍数に当てはまらないことがわかる  
また、case で示される注文の順番を変えるのも考えてみると、ラストオーダー以外の注文の順番を変えても結果には影響しない
このことをまとめ、最短オーダーをするための条件にすると下記のように考えられる  

- 条件1 10の倍数の値は途中で注文する(無駄のないオーダー時間になる)
- 条件2 最も時間のかかる注文は、最後に注文するのはNG
- 条件3 最後の注文だけは届く時間が10の切り上げにならない
=> 1の位が1以上で、かつ最も低い注文を最後にするのが最短。最後の注文以外は結果に影響しない

よって、下記のようなロジックを組むのがいい

- A~E の注文から、1の位が1以上で、かつ最も低い注文を探索
- 上記要素以外を、10の倍数切り上げの値で合算
- 上記要素と合算値を回答として出力
- 全て同じ注文時間だった時を考慮しての初期値の設定だけはしておくこと

### pseudo code


```
var int a, b, c, d, e
list [a, b, c, d, e]

min = 124 
minI = 0 
for i to list.length 
    mod = list[i] % 10
    if mod not 0 & mod < min 
        min = mod
        minI = i

var int answer 
for i to list
    if i not minI
        answer += floored list[i]

print(answer + list[minI])
```

### tips

回答時間: 40m

time complexity は O(1)