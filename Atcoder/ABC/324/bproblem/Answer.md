[問題文](https://atcoder.jp/contests/abc324/tasks/abc324_b)

#累積和

## 問題文の把握

特に難しいところはない

## ロジック
`N = 2^x * 3^y`  
x, y の組み合わせをまともに計算すると際限なくなってしまう  

そこで、因数分解の考え方を応用してみる。  
まず、具体例を見て気づく点として N は必ず偶数になる。
なので、 N を 2 で割り切れるだけ割っていくと、3 で割り切れる値になる。

```
144 = 2^4 * 3^2
144 / 2 = 72
72 / 2 = 36
36 / 2 = 18
18 / 2 = 9
9 / 3 = 3
3 / 3 = 1
```

N よって、2,3それぞれで割っていて、剰余がないならYesと出力していけば良い  
golangにおいては for文を使って実装することになるので、無限ループにならないよう条件を書くのを忘れずに  

```
while N % 2 == 0:
    N /= 2
    
while N % 3 == 0:
    N /= 3
    
if N == 1:
    print("Yes")
else:
    print("No")
```