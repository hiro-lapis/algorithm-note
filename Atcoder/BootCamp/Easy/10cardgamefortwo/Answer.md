[問題文](https://atcoder.jp/contests/abc088/tasks/abc088_b)

## 問題文の把握

- N 枚のカードを使ってゲームをする
- i 枚目のカードの数字は ai であり、合計 N 枚ある  
- Alice, Bob の2人で、場に出されたカードを交互にとる
  - 順番は Alice => Bob の順
- カードを取り終えたらゲーム終了。とったカードの数の大きい方が勝者となる
- それぞれ最適な戦略をとった場合に Alice は Bob に何点差で勝つか求めよ

- 1 <= N <= 100
- 1 <= ai(各カード) <= 100

```
N
a1 a2 ... aN

例
2
3 1
=> 2

3
2 7 4
=>5
```

## ロジック

Alice も Bob も最善の戦略をとるということは、 数字の降順で交互にカードをとっていくことになる 　
カードの数字が何番目に大きいかは、場にあるカード全ての数字がわかる必要がある  

よって、
「N 枚すべてのカードを DESC で並び替える。その後、N 回だけ得点計算処理を実行  
降順に並び替えたカードを頭から 足し算, 引き算の順に交互に計算。最後に結果を出力」する


### pseudo code


```
var n, sum int
var cards [n]
for i to n
    cards[i] = stdIn
cards.sortDesc()

for j=0 to n
    if j % 2 == 0
        sum += cards[j]
    else  
        sum += cards[j]
print(sum)
```

### tips

N 枚のカードの入力を受け入れるために配列変数を作成するが、この時に `var cards [n]int` とすると構文エラーになる(定数式でないといけないと言われる)  
この場合はスライス `var cards = make([]int, n)` を定義する

