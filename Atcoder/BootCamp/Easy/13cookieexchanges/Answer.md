[問題文](https://atcoder.jp/contests/agc014/tasks/agc014_a)

## 問題文の把握
- 子供3人が A,B,C のクッキーを持っている
- それぞれ順番に、下記の行為をする
  - 3人は同時に、各々が持っているクッキーを半分づつに分けて、残りの2人に渡す
  - 半分シェアを繰り返す
  - 誰かの持っているクッキーが奇数個になったら、その時点で終了
- クッキーの交換は何回できるか
- ただし、無限にできる場合もある

1 <= A,B,C <= 10^9

```
A B C 

4 12 20
=> 3
1: 16 12 8
2: 10 12 14
3: 13 12 11

2 2 2
=> -1 (無限)
```

## ロジック

いくつか例を試してみると、ABC が**同値の偶数**の時に無限に交換できることがわかる
基本はA B C それぞれを二等分して他に割り振るループを作ればいい  
ループの最初に各要素が偶数であることを判定するのを忘れずに

### pseudo code


```
var a b c int

if a %2 == 0 && a == b && b == c
    print(-1)
    return

var count
while a,b,c is even
    count++
    aHalf = a / 2
    bHalf = b / 2
    cHalf = c / 2
    
    a = (bHalf + cHalf)
    b = (aHalf + cHalf)
    c = (aHalf + bHalf)
print(count)
```

### tips

time complexity は...不明
ABCがそれぞれ 2^n の値である時に、分割回数が多くなる  
