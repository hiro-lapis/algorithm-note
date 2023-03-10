[問題文](https://atcoder.jp/contests/abc157/tasks/abc157_b)

## 問題文の把握

-   3 x 3 のビンゴゲームで、縦列は i, 横は j で表す
-   縦横から指定されるマスは A
-   N コの数 bN が選ばれ、その数がビンゴカードにあった場合に印がつく
-   N コの数が選ばれた時点でビンゴが達成するかを判定する。達成は縦・横・斜めのいずれかで 3 マス 1 列並んでいる時
-   1 <= Ai, j <= 100

-   入力について

```
A1,1 A1,2 A1,3
A2,1 A2,3 A2,3
N
b1
bN
```

最初にビンゴカードの数字 AN が与えられる　　
N は 1~10 のため 9 コ以上は与えられない
カードのマス目は 3 \* 3 のため、A の数は 9 である

ビンゴで当たり番号は BN である　　
N は 1~10 のため 9 コ以上は与えられない

## ロジック

最初に A の二次元配列を受け取り、カードの各マス目の値を受け取る  
list(A と同じ要素数の bool 値をもつ配列)を定義
＊初期値は全て false にする
ヒットした時に対応するマス目(index) を true に更新する    
N を受け取ったら、N の数だけ for ループを回し、以下の処理を行う  
- 入力値 b を受け取り、A の中に同じ値があるかチェック
- 同じ値があったら、A のどの index にあるかを取得
- list 内, A の index と同じ index の値を true に更新  
ループが終わったら、
縦、横、斜めのいずれかで揃っているか判定して Yes/No を出力する

### pseudo code

```
var A = [9]
var N = stdIn
var list = [9]false
for i to N
    var b = stdIn
    key = getKey b in A
    if key valid
        list[key] = true

if hit is row || column || diagonal
    print("Yes")
else
    print("No")
```

### tips
time complexity は、N 回だけ入力値を受け取ってヒット判定を行うため、O(N) である  
理屈の上では、入力値 N が少ない場合において、ビンゴが完成しているか判定する処理もボトルネックになりうるが、ビンゴカードが 3*3 の 9 マスであるため、定数時間であり、大した処理時間にはならない  
