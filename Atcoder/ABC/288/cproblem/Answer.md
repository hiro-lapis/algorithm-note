[問題文](https://atcoder.jp/contests/abc288/tasks/abc288_c)

## 問題文の把握

- N vertex, M edge からなるsimple undirected graph が与えられる
- i 番目の edge は Ai-Bi を結んでいる
- グラフから 0 本以上のいくつかの edge を削除して graph が cyclic にならないようにしたい
- この時に削除する edge の最小数を求めよ

- 制約
- 1 <= N <= 2*10^5
- 1 <= M <= 2*10^5
- 1 <= Ai,Bi <= N
- 与えられる graph は simple

```
N M
A1, B1
A2, B2
AN, BN

AM BM

case1
6 7
1 2
1 3
2 3
4 2
6 5
4 6
4 5

         ①
        /  \
       ② - ③
       |
       ④ - ⑤
        \  /
         ⑥
上記のようなgraphが作成される
cyclicにならないようにするには、最低2本消せばいい

case2
4 2
1 2
3 4
=> 0
conncted でないため、削除する必要がない  
```



## ロジック

単に探索するだけなら visited list を作成し、「from/toどちらかが探索済でなければ、削除するedgeである」
というロジックでedgeをカウントするだけでいいが、case1 のような時に edge case を考慮する必要がある

```
case1
6 7
1 2
1 3
2 3
4 2
6 5
4 6 - ココ
4 5

         ①
        /  \
       ② - ③
       |
       ④ - ⑤
        \  /
         ⑥
```
「from/toどちらかが探索済でなければ、削除するedgeである」というロジックだと、
ココ と書かれているedge も削除すべきedge と判定され 3(誤答)になってしまう  

解説を見た限りではいくつかのアプローチがあるようだが、GIT で学んだ disjoint set union を使ったアプローチが非常にわかりやすかったので、これを使ってとく  

- 削除すべきedge数を管理する変数 count を定義

- M コのedge が結びつける vertex U,Vが同じグループに属するかをチェック
  - 属していたらグルーピング
  - 属してなかったら、その edge は不要なので count を increment
- count を出力

なお、今回の問題において、MSTが完成するかどうかは重要ではないので、Edgeを全て処理することだけ考えればいい  

[解答](https://atcoder.jp/contests/abc288/submissions/38633996)

- disjoint set union, union find について

GIT の復習になるが、dsu はMinimum Spanning Tree (最小限のedge からなるconnected graph の作成)で使われる algorithm

- 最小限の edge からなる graph を作りたい場合、まずバラバラの vertex からなる disjoint set(非連結セット)を作る　　
- 次に edge が結びつける vertex U,V を比較して、vertex 同士を結びつけていく
  - vertex U,V が異なるグループなら、union(結合)する
    - グルーピングにおいて、U,Vどちらかをparent node に設定する(index値が小さいをparentにするのが多い)
  - 同じグループなら結合しない (MSTにおいては不要なedge であるという判定が可能)
- 上記を全ての edge を取り出して行う
- 一通り探査を終えたあと、dsu にはグループを見て、グループが 1 つにまとまっていて(connected)、 かつ N-1(最小限のedge数からなる) 場合 MST といえるであると判定できる  

上記の通り、disjoint set はデータ構造、ds を使って MST を構成するためのロジックが union find となる  

[参考:Kruskal's Algorithm: Disjoint Sets](https://learning.edx.org/course/course-v1:GTx+CS1332xIV+2T2022/block-v1:GTx+CS1332xIV+2T2022+type@sequential+block@962fb5f148374d92bf4a9ce31d44cebd/block-v1:GTx+CS1332xIV+2T2022+type@vertical+block@596ce493152a496490c350d942a42f39)

### pseudo code

```
var n, m int

dsu = make disjoint set of n

var redundantCount = 0
for i to m
    receive edge input 
    if same group edge.U and edge.V in dsu
        increment redundantCount
    else 
        union edge.U and edge.V

print(redundantCount)
```

### tips

- time compelexity

m コの edge を受け取る箇所は計算量 m  
edge の U,V それぞれのsame group 判定、union 処理は 1回あたり 1 (constant time) で、これを m 回行う。計算量は m  
よって、計算量は O(2m) => O(m) となる  
＊ kruskal を使うと O(m log(n)) とかに改善できる

## 公式解説

- 最終的に残す edge 数を L とすると、答えは M - L
- graph の連結部分の個数を S とする
- acyclic な graph になるには、N - 1 vertices のedge数である必要がある
- つまり、L(最終的に残るedge数) = N - S ()

###  disjoint set を使った解法

disjoint set (非連結セット)による解答もバリエーションとして有効
- 全てのvertexを連結してないvertex集合のgraphとして作成
- vertex 同士のroot vertexを比較し、異なるgroup の場合は、どちらかに取り込む形で同じグループにmerge
