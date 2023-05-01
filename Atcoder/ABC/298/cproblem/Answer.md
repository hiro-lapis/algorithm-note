[問題文](https://atcoder.jp/contests/abc298/tasks/abc298_c)

#配列

## 問題文の把握

- 1~Nの番号がついた空の箱とカードがある
- Qコのクエリが与えられるので、順番に処理せよ(結果も入力値に依存する)
  - `1 i j`: カード1枚に i を書き込み、ハコをjに入れる
  - `2 i j`: i 番目の箱に入っているカードの数を昇順で答える
  - `3 i`: 数 i が書かれたカードが入っている箱の番号をasc で答える 

## ロジック

入力のi,jの使われ方がクエリの値によって変わる
カードの値は、クエリ1:i, クエリ3:i
箱の値は、クエリ1:j,クエリ2:i




### pseudo code


```
```

### tips
