[問題文](https://atcoder.jp/contests/abc152/tasks/abc152_c)

#tag

## 問題文の把握

- 1,...N という順列 P1,...Pn が与えらえる
- 下記条件を満たす整数 i を数えよ 整数 i (1 <= i <= N) の個数を数えよ
  - 任意の整数 j (1 <= j <= i) に対して Pi <= Pj

- 制約
- 1 <= N <= 2*10^5

```
N
P1,...Pn

5
4 2 1 5 3
=> 3

case2
4
4 3 2 1 
=> 4

case3
6
1 2 3 4 5 6
```

## ロジック
問題文がわかりにくい  
噛み砕いていうと、P1~P[i]までの数列で i 番目の数字が最小値であるようなときは、条件を満たすということ  

よって、最小値を更新する回数をカウントするロジックを組めば良い

### pseudo code


```
var n int
var list []int

var min = INF
var answer int

for i to n
  tmp = stdIn
  if min > tmp
    answer++ 
    min = tmp
    
print(answer)
```

### tips

- time complexity O(n)

