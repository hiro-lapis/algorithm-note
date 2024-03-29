[問題文](https://atcoder.jp/contests/abc139/tasks/abc139_c)

#数列

## 問題文の把握

- N コの数字からなる数列がある
- 左から i 番目のマスの高さは Hi
- 好きなマスに立ち、右のマスが現在いるマスの数値以下なら右に移動する、という行動を繰り返すとして、
- 最大何回移動できるかを求めよ

- 制約
- 1 <= N <= 10^5
- 1 <= Hi <= 10^9

```
N
H1 H2 HN

case1
5
10 4 8 7 3
=> 2

case2
7
4 4 5 6 6 5 5
=> 3
```


## ロジック

for 文で Hi と Hi+1 を比較することで、移動開始地点を決めることはできそう  
しかし、単に数列の要素を左から比較していくと N! の計算量が発生するため、工夫が必要

```
7
4 4 5 6 6 5 5
```
さらにいうと、↑ のような数列で、冒頭の4から探索を開始して,4, 5と探索して 5 にぶつかった次の探索の開始地点は 6 からにすることがのぞましい  
2 つめの 4 から 4 => 5 と1回移動しても、最大回数には満たないため  
このあたりを考慮してロジックを組む  

### tips

- 解答時間 26m
