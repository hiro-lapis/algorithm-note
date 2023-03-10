[問題文](https://atcoder.jp/contests/abc292/tasks/abc292_c)

#整数論
#分割統治

## 問題文の把握

- 正整数 N が与えられる
- 正整数 (A, B, C, D) であって、AB+CD = N を満たすものの個数を求めよ

- 制約
- 2 <= N <= 2*10^5
制約により、解答値は 9 * 10^18 以下となる

```
case1
4 
=> 8
(1, 1, 1, 3) => 1*1 + 1*3 = 4
(1, 1, 3, 1)
(1, 3, 1, 1)
(3, 1, 1, 1)
(2, 1, 1, 2) => 2 + 2
(1, 2, 1, 2)
(1, 1, 2, 2)
(2, 1, 2, 1)


```


## ロジック

入力値の最大は 2*10^5 (20万) 値自体は大したことないが、ABCD それぞれに1,2,3,と当てはめていくと exponential になる  
地道に計算していたら、 9 * 10^18 回計算することになり TLE 確実
よって、計算回数を何らかの形で省略する必要がありそう

ここで、いくつかの数学的考察を挟んでいく

### 項のまとめあげ

AB + CD = N

4つの数字の記号として記載があるが、**足し合わせる値としては AB CD の2つ**にまとまっている
そこで、
AB を X 
CD を Y 
という値として固定し、その値となる組み合わせを考えるようにする

つまり、 N = 4 だとしたら

x + y
1 + 3
2 + 2
の組み合わせを考えることになる

### X,Y,N の関係

X + Y = N ということは、X 側の値がもとまることで自然と Y の値も決まるということ    
つまり、 
X + (N - X) = N  
X, Y を構成するAB, CD の組み合わせは複数あっても、X の値が決まった時点で Y の値は自然と定まる    
この意味において X の値は複数パターンあっても、Y は 1 パターンしか存在しえない  

### 組み合わせの性質

数学的性質上、**X, Y を構成する整数の組み合わせは、X, Y の約数個数**である
約数は X % A == 0 になるかどうかで判断できる
また、B の値は X / A で求められる

X = 3 だとしたら、約数は 1, 3 の 2 つ。A, B それぞれが1, 3 になるケースがあるので2 通りの組み合わせがある
X = 4 だとしたら、約数は 1, 2, 4 の 2 つ。A, B それぞれが1, 4 にになるケースがあるので2 通りの組み合わせがあり、  
さらにAB ともに 2, 2 になる組み合わせ 1通りの組み合わせがある

=> ABともに約数である時、A != B なら2通り、A == B なら 1 通りの組み合わせが存在する


### 探索範囲の最小化

組み合わせは for 文で計算していくことになるが、forの範囲を狭めることができる

100 という値は 10^2 である
X = 100 とした時に、A, B どちらかの値は 10 常に 平方根の 10 以下に収まる
ex. 10*10, 50*2, 20*5

よって、X = 100 の組み合わせについて探索する場合、
A: 1 -> 100 
B: 1 -> 10

この探索範囲で X % A or B == 0 かどうかを計算することで約数が計算できる

これと同様のことが、Y 側にもいえる

X, Y それぞれにこのループ処理を行い、X + Y は N になる


A のtime complexity: 1 to X
B のtime complexity: 1 to √X
C のtime complexity: 1 to Y
D のtime complexity: 1 to √Y
=> X*√X + Y*√Y
X + Y = N なので、いかに簡略化できる
=> N√N


```

```

いくつかcase を考えてみる

```
2
=> 1
(1,1,1,1)

3
=> 4
(1,1,1,2)
(1,1,2,1)
(1,2,1,1)
(2,1,1,1)

4 => 8

5
=> 14
(4,1,1,1) => 4通り
(3,1,2,1) => 8通り
(1,3,2,1)
(3,1,1,2) 
(1,3,1,2)
(2,1,3,1) 
(1,2,3,1)
(2,1,1,3)
(1,2,3,1)
(2,2,1,1) => 2通り
(1,1,2,2)  

6
=> 20
(5,1,1,1) => 4通り
(4,1,2,1) => 8通り
 ↑ 4 を 2 に分解できる
(2,2,2,1) => 4通り
(3,1,3,1) => 4通り

7
=> 28
(6,1,1,1) => 4通り
(3,2,1,1) => 4通り
(5,1,2,1) => 8通り
(4,1,3,1) => 8通り
(2,2,3,1) => 4通り

8 => 41
(7,1,1,1) => 4通り
(6,1,2,1) => 8通り *
(3,2,2,1) => 8通り * => 6 の部分を分解
(5,1,3,1) => 8通り
(4,1,4,1) => 4通り
(4,1,2,2) => 8通り * => 4 の部分を分解
(2,2,2,2) => 1通り

9 => 44
(8,1,1,1) => 4通り
(4,2,1,1) => 4通り
(7,1,2,1) => 8通り
(6,1,3,1) => 8通り
(3,2,3,1) => 8通り
(5,1,4,1) => 8通り
(5,1,2,2) => 4通り

```

### pseudo code


```
```

### tips

