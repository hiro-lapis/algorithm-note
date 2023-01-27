[問題文](https://atcoder.jp/contests/abc160/tasks/abc160_c)

## 問題文の把握

-   全周 K メートルの円形の湖の周りに、 N 件の家がある
-   i 番目の家は北端 Ai メートルの位置にある
-   いずれかの家から出発して N 件全部の家を尋ねる時、最短移動虚を求めよ

-   制約
-   2 <= K <= 10^6
-   2 <= N <= 2\*10^5
-   0 <= A1 <...AN < K

入力

```
K N
A1 A2...AN

// case1.
20 3
5 10 15
=>10

// case2.
20 3
0 5 15
=> 10
```

K=20 の場合、20 の次は 0 になるような circulation の構造になっている

## ロジック

入力値が 0 <= A1 <...AN < K であることから、スタートから向かった方向に一直線に歩いていけば最短距離になり、
いわゆるグラフ探索のような処理は不要であることがわかる  

ここでいくつかの入力例を考えてみる
```
10 3
1 2 9
=> 1の時9, 2 の時3, 9の時3
 2~9の区間が一番離れている(7)

20 3
6 12 13
=> 6の時7, 12 の時13, 13の時7
 6~13の区間(0経由)が一番離れている(7)


20 3
0 12 19
=> 0の時19, 12 の時8, 19の時13
 0~12の区間が一番離れている(12)

20 4
1 5 10 19
=> 1:16, 5:15 ,10:11, 19:11  
 10~19の区間が一番離れている(9)
```

このように、一番離れている区間(最長距離)について、逆に進めば最短経路でいけることがわかる  
この時、最短距離 = K - 最長距離 である  
この関係をロジックとして組むと、

- 移動距離の最大値変数を定義
- `n[i] - n[i+1]`(i ~ i+1 地点の距離)を計算。これまでより最大値が高かったら、最大値を更新
- これを N 回繰り返す
注意点としては、最後の家から北端を跨いで最初の家を尋ねる距離の計算
他とは計算式が異なる  
```
20 3
6 12 13
13 から 6 への移動距離
=> 13 から北端20 へ行く(7) + 北端から6 へ => 13
(K - a[N-1]) + a[0] となる
```

最後に K - 最長距離を出力すれば、回答となる

### pseudo code

```
var K, N
var A[]int

max = 0
for i to N
    if i == N
        if max > (K + A[N]) - A[0]
        
    else
        if max > (A[i+1] - A[i])
            max = (A[i+1] + A[i])
print(K - max)
```

### tips

計算回数は N に依存、O(n) である  