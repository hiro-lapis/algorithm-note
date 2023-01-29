[問題文](https://atcoder.jp/contests/agc027/tasks/agc027_a)

## 問題文の把握

- N 人の子供がいて、それぞれに 1, 2, ... N の番号が振られている
- x 個のお菓子を子供たちにあげる。この際、お菓子をもらわない子供がいてもいい
- 各 i(1 <= i <= N) について、子供 i はちょうど ai このお菓子をもらうと喜ぶ
- 喜ぶ子供の人数の最大値を求めよ

- 制約
- 2 <= N <= 100
- 1 <= x <= 10^9
- 1 <= ai <=10^9

```
N x
a1 a2 ...aN

case1
3 70
20 30 10
=> 2
(20 30 20) と配れば 2 人の子供がよろこむ

case2
3 10
20 30 10
=> 1
(0 0 10) と配れば 1 人の子供がよろこむ

case3
4 1111
1 10 100 1000
=> 4
(1 10 100 1000)
```
## ロジック

満足する子供の**人数**を最大化したい場合、少ないお菓子で満足する子供に優先して配ればいい  

ポイントは x コのお菓子を全て配り切らないといけない  
ということ

- a1~aN までの子供の配列を、昇順にソート
- ループで a1 から子供の欲しがるお菓子の数の分だけ x を差し引く
- 差し引いて x が 0 以上なら満足する子供の人数(count)を増やす
- ループの最後で x ≠ aN なら、count は増やさない

### pseudo code

```
var N, x
var list[N] int
list.sortAsc()
var tmp = x

for i to N
    if x > tmp
        break
    if i == N -1 & tmp > list[i]
        break
    
    tmp = tmp - list[i]
    if tmp >= 0
        count++
    else 
        break    
print(count)        
```

### tips

sort を N 回、その後に お菓子を配るループで N 回なので、
time complexity は O(n^2)  

解答にかかった時間: 20m

テキストに考察を書き込むのにざっくり10分  
考察書き込みなしで解いて、解答時間が fix した後に考察を書く、という手順でもいいかも  

- Go でINT値配列を昇順にソート  

- [sort.Slice()](https://pkg.go.dev/sort@go1.19.5#Slice) で要素間の比較をして並べ替える
- 組み込みの [sort.Ints()](https://pkg.go.dev/sort@go1.19.5#Ints) 
を使う方法がある  

```
sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })

もしくは、sort.Ints(list) でもいい  
```

- 降順にソート
昇順の逆をやる方法と、組み込み関数を組み合わせる方法がある

```
sort.Slice(list, func(i, j int) bool { return list[i] > list[j] })

// IntSlice() でinterface convert し、逆順にソート配列を作って slice に変換
sort.Sort(sort.Reverse(sort.IntSlice(list)))
```

- Go で文字列配列を asc/desc ソート