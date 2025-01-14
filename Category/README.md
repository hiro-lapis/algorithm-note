ここでは基礎的なデータ構造とアルゴリズムの確認ができる  

実装を元にしたロジックの確認は各項目を参照するとして、  
概念が掴みにくい場合は下記のビジュアライズツールを使うとわかりやすい    
[CS 1332 Data Structures and Algorithms Visualizations](https://csvistool.com/)  

- Search
  - [binary search](./binarysearch): sort済配列を対象にした配列要素の高速な検索
  - [linear search](./linearsearch): 配列要素の順次検索(全探索)
- String
  - [analyze syntax](./syntax)構文解析。文字の特定のパターンを検索する
- Graph
  - [Disjoint Set Union](./disjointsetunion): graph vertex のグルーピング
  - [BFS](./bfs): 幅優先探索
  - [DFS](./dfs): 深さ優先探索
  - [Dijkstra](./dijkstra): 最短経路問題
- Math
  - https://qiita.com/drken/items/a14e9af0ca2d857dad23
  - [prime number](./prime): 素数判定
  - [prime factorization](./primefactorization): 素因数分解
- List manipulation
  - [prefix sum](./prefix_sum) 累積和(配列の任意の区間の総和を求める)
  - [imos](./imos) いもす法。累積和を応用した加算テクニック
  - [shakutori](./imos) しゃくとり法配列の添字2つを交互に動かして効率的に走査する手法
  - [division scheduling](./imos) 配列の特定区間。貪欲法を使う
- DP 動的計画法(小さな問題の解決をメモして大きな問題を効率的にとく)
  - 最大最小: その範囲での最大最小
  - yes/no: その条件下でのtrue/false
  - 組み合わせ: その条件下での組み合わせ数
  - [chokudai](./dp/chokudai) 文字列中のキーワード組み合わせパターン数検索
  - [70_100to105](../Atcoder/BootCamp/Easy/70_100to105) ナップサック問題


## Requirements

各種グレードになるために必要なテクニック  

- Brown
  - 全探索, 順列全探索
  - 累積和,imos法,しゃくとり法,区間スケジューリング問題,二分探索,めぐる式二分探索,
- Green

## Problems

練習できるサイト

[Atcoder problems](https://kenkoooo.com/atcoder/#/table/)
[AtCoderの問題を分類しました【カテゴリ】 zenn](https://zenn.dev/koyanagihitoshi/books/atcoder-classification-6)
問題をカテゴリ分類してくれている
[EDPC Educational DP Contest](https://atcoder.jp/contests/dp/tasks)   
DP問題ばかりのコンテスト
[競技プログラミングにおける動的計画法問題まとめ はまやん](https://blog.hamayanhamayan.com/entry/2017/02/27/021246)
DPに関する考察まとめ
[AtCoder 問題カテゴリー分類](https://qiita.com/c-yan/items/56a051d826b873b4f78d#imos-%E6%B3%95-imos-method)
カテゴリー幅広く網羅している