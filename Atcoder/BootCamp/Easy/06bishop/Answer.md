[問題文](https://atcoder.jp/contests/abc121/tasks/abc121_b)

## 問題文の把握

- 縦 H マス、横 W マスの盤面
- 角の動きをするコマが左上隅のマスにおいてある
- コマが 0 回以上、好きなだけ動ける
- 到達できるマスの数を答える
- マスの座標を上から r1 番目, 左から c1 番目とするなら、r1+c1 = r2+c2, r1-c1 = r2-c2 のどちらか一方が成り立つ時に移動が成立する

## ロジック

H=4, W=4 だとして、(1列:2, 2列:2, 3列:2, 4列:2) = 8 マス
H=4, W=5 だとして、(1列:3, 2列:2, 3列:3, 4列:2) = 10 マス
H=4, W=6 だとして、(1列:3, 2列:3, 3列:3, 4列:3) = 12 マス
H=5, W=4 だとして、(1列:2, 2列:3, 3列:2, 4列:3) = 10 マス
H=5, W=5 だとして、(1列:3, 2列:2, 3列:3, 4列:2, 5列:3) = 12 マス
H=3, W=3 だとして、(1列:3, 2列:2, 3列:3, 4列:2, 5列:3) = 12 マス

1 列(横)単位で見た場合、
- 奇数列で移動できるマスの数は W/2 の切り上げ
- 偶数列で移動できるマスの数は W/2 の切り捨て

である
また、マスの全体数(H*W)との関係を見るに、
H,W どちらも偶数 = 移動しうるマスは H*W /2  
H,W どちらかが偶数 = 移動しうるマスは H*W /2
H,W どちらも奇数 = 移動しうるマスは (H*W /2) + 1
となることがわかる

この性質を利用して、
「H*W/2 の値を計算。H, W がどちらも奇数の時はさらに +1 した値を結果として出力する」  
という計算をする

追記：H=1 もしくは W=1 の時はエッジケースとなり、上記式が成り立たない点を考慮する必要がある
### pseudo code

```
var H, W
var num = 0
if H == 1 || W == 1
    print(1)
    return
    
if H % 2 ==1 && W % 2
    num = 1
    
print(H*W/2 +1)
```

### tips

計算量は O(1)