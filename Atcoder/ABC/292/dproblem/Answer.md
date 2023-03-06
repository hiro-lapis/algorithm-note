[問題文](https://atcoder.jp/contests/abc292/tasks/abc292_d)

#graph
#connected

## 問題文の把握

- undirected graph として、N コの vertex M コの edge が与えられる
- edge i は <u, v> を結んでいる
- graph の connected である時、含まれる vertex 数と edge 数が日宇土市以下どうかを判定せよ

- 制約
- 1 <= N <= 2*10^5
- 0 <= M <= 2*10^5
- 1 <= ui <= vi <= N

```
N M
u1 v1
uM vM

case1
3 3
2 3
1 1
2 3
Yes
↓
1 
2 <-> 3
2,3を結ぶedgeが2つあり、それにより2 vertex　が結ばれている
1,1を結ぶedgeが1つあり、それにより1 vertex　が結ばれている
vertex(2) * edges(2) である
vertex(1) * edges(1) である

case2
5 5
1 2
2 3
3 4
3 5
1 5
Yes
↓
1 - 5
|   |
2 - 3 - 4

5 つ全てのvertex が繋がっており、かつedge 数は5
よって Yes

case3
13 16
7 9
7 11
3 8
1 13
11 11
6 11
8 13
2 11
3 3
8 12
9 11
1 11
5 13
3 12
6 9
1 10
=> No
```

## ロジック

graph の基本として non-cyclic connected であるときは edges = vertices - 1 である必要がある  
というルールがある
しかし、case1 のようにconnected でなくとも YES となりうるケースがありうる  

graph 理論では
「connected graph を形成するかどうか」
「vertex 同士が同じconnected group に属するかどうか(disjoint set)」 
という2つの理論があるが、case1 があることを考えると 後者の考え方を進めていった方が良さそうに思える

以下、disjoint set のロジックについて振り返ってみる

[disjoint set]  
- edges 数分だけ下記の処理
- edge が結ぶ vertices が同じグループに属するかどうかを比較
  - 異なるグループのときは union 。どちらかのグループに結びつける
  - 同じグループのときは何もしない

この処理は edge 数だけ行う処理なので、今回の問題でいえば M 回やることになる
そして、disjoint group set の数も同時に管理している

下記のように処理をカスタマイズすれば、問題を解けそう  

- 処理したedge 数をカウントする変数を定義
- 処理した変数を increment するが、自己参照edge は処理したedge 数に含めない
- 各ループの最後に「set 数とループ回数(処理したedge数)が同値かを比較する」

### pseudo code


```
```

### tips

