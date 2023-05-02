[問題文](https://atcoder.jp/contests/abc298/tasks/abc298_b)

#多次元配列
#配列の組み替え

## 問題文の把握

- 0,1からなる N x N の行列 A, B が与えれる
- A i,j の要素をAi,j という風に表記する
- 下記の回転操作を行うことで、Ai,jである箇所全てにおいて、同様にBi,jであるかどうかを判定せよ
  - 1 <= i,j <= N を満たす全ての整数の組 (i,j) について同時に、A(N+1-j, i) で置き換える

- 制約
- 1 <= N <=  100
- A,B の要素は 1,0

```
N 
A1,1 A1,2 A1,N
AN,1, ...AN,N

B1,1 B1,2 B1,N
BN,1, ...BN,N

case1

3
0 1 1
1 0 0
0 1 0 * 
1 1 0
0 0 1
1 1 1
=> Yes

* を区切りとして、A,Bが与えれる
2回回転操作を行うと、A行列上で1になっている箇所は全てB上でも1になる


```


## ロジック

- 回転操作

本文中で書かれている回転操作というものについて考えてみる
case1　で書かれている内容は以下

```
case1

0 1 1
1 0 0
0 1 0
↓
0 1 0
1 0 1 
0 0 1
↓
0 1 0
0 0 1
1 1 0
```

時計回りに要素が組み替えられているのがわかる  
時計回りであるということは、4回目の操作で初期の配置に戻るということなので、  

回転操作は 0 ~ 3回の内に収まることがわかる

・回転操作のロジック

`整数の組 (i,j) について同時に、A(N+1-j, i) で置き換え` 
↑のロジックを二重 for 文で回せばOK

### pseudo code

```
var n 

var a = [n][n]int
var b = [n][n]int
for i=0 to n
  for j to n
    a[i][j] = stdIn
    
for i=0 to n
  for j to n
    b[i][j] = stdIn
  
for i=0 <= 4
  if i == 4 
    print(No)
    break
  if a == b 
    print(Yes)
    break
  a = rotate(a)

func rotate(n, a)


```

### tips

