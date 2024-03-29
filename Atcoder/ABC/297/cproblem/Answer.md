[問題文]()

#文字列

## 問題文の把握

- H コ、長さ W の `.` `T` からなる文字列 S1,S2,...SH がある
- 以下の操作を 0~何回でも行う
  - 1 <=i<=H, 1 <=j<=W-1, を満たす整数で Si[j], Si[j+1] の文字も `T` である文字を対象に、
  - Si[j] の文字を `P` で置換し、Si[j+1]を `C` で置換する
- 最大回数上記の操作をするとして、終了後のS1,S2,SH としてありうるものの一例を出力せよ

- 制約
- 1 <= H <= 100
- 2 <= W <= 100

```
case1
2 3
TTT
T.T
↓
PCT
T.T

case2
3 5
TTT..
.TTT.
TTTTT
↓
PCT..
.PCT.
PCTPC


case3

```

## ロジック

条件文が回りくどい感じなので、caseを例に噛み砕いてみる

・Si[j], Si[j+1] の文字も `T` である => T が2回続く文字列 S について
・Si[j] の文字を `P` で置換し、Si[j+1]を `C` で置換する => PC にする

なんのことはない、`TT` という文字があったら`PC` にする、ということ

```
case1
2 3
TTT
T.T
↓
PCT
T.T
```

ロジックについて考える時の考慮すべき点としては、都度都度文字列置換の対象にした方がよさそうなところ  
途中で置換対象になった`j+1`番目の文字は、その後も継続する処理で`j`番目の文字になる  
判定のみ行って、後々一括で置換する、とかしようとすると誤った判定が入る可能性がある  
`TTT` みたいな文字で `TPC` みたいになったり `PPC` みたいにならないようにする

文字列数や文字数自体は少ないため、処理速度を気にする必要はない  

### pseudo code


```
```

### tips

