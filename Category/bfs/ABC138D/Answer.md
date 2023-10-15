[問題文](https://atcoder.jp/contests/abc138/tasks/abc138_d)

#bfs

## 問題文の把握

- 1~N までの番号がついた N vertex をもつ root tree が与えられる。
- tree node の番号は 1 で、 i 番目の edge はそれぞれの vertex を結ぶ
- ここから、以下のような操作を Q 回行う
  - j(1 <= j <= Q) :頂点pj をROOT とする partial tree 全てのNODEのカウンターにxj をプラス
全ての操作をした後の、各NODEのカウンター値を求めよ

- 制約
- 2 <= N <= 2*10^5
- 1 <= Q <= 2*10^5
- 1 <= ai < bi <= N
- 1 <= pi <= N
- 1 <= xj <= N
- 与えられるグラフは tree 

```
N Q

a1 b1
aN-1 bN-1
p1 x1
pQ xQ

case1
4 3
1 2
2 3
2 4
2 10
1 100
3 1
=> 100 110 111 110

case2
6 2
1 2
1 3
2 4
3 6
2 5
1 10
1 10
=> 20 20 20 20 20 20

```

## ロジック

まずどんな条件下の問題かを整理する

### 全探索問題である

・木が与えらえる
=> 孤立した vertex は存在しない。よってcycleなどの問題ではない
=> root tree なので、頂点は1つで、下に行くほど枝が増える
=> weightyがない。undirected でもない。よって、最短経路問題ではない

### edge数は決まっている

```
N Q

a1 b1
aN-1 bN-1
p1 x1
pQ xQ
```
=> a b は edge がどこからどこへ伸びるかの入力値。数は常に N-1
=> p は各ループで root node として扱う vertex の index
=> x は加算するポイント

### visited list を扱う問題である

各頂点のカウンターの値を求めよ
=> visited list の要領で、各 vertex の探索が必要

case1について、
```
case1
4 3
1 2
2 3
2 4
2 10
1 100
3 1 * ここの処理で 111 となっていることから、p になった頂点自体加算対象になる
=> 100 110 111 110
```

これらを元にロジックを作成する

- 入力値 N を受け取り、探索済みslice(index N)を INT 型で作成
- edge 情報を受け取る二次元 slice を作成
  - 1 次元はint(vertex number), 2 次元は tuple(a, b で、aと1次元目の値は同じ)
- N-1 回のループで edge 情報を受け取る
- 入力値 Q を受け取り、j => Q の以下ループを実行
  - 入力値を tmpI tmpVal として受け取る
  - 探索済みslice[tmpI] に,tmpVal を加算
  - edge 情報の入った二次元 slice[tmp] として、対応する vertex のedge を参照
  - 探索済みslice[edge.b] に,tmpVal を加算
    - edge がなくなるまで続けたら、次のjループへ
- 探索を終えたら、for 文で探索済みsliceの値を出力


### pseudo code


```
```

### tips

#### グラフを表現するデータ構造の種類

- adjacent matrix(隣接行列)
2 次元配列を使って表現する
[1, 0, 0, 1] => vertex1 から繋がりのあるvertex
[0, 1, 0, 0] => vertex2 から繋がりのあるvertex
[0, 0, 1, 0] => vertex3 から繋がりのあるvertex
[1, 0, 0, 1] => vertex4 から繋がりのあるvertex

- adjacency list(隣接リスト)
今回の問題で使う

- adjacency set

