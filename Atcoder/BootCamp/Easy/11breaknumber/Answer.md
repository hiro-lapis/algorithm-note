[問題文](https://atcoder.jp/contests/abc068/tasks/abc068_b)
git branch --con
## 問題文の把握

- 高橋君は 2 で割れる数が好き
- 整数 N が与えら得る時、1 <= N の間の数字で、最も 2 で割れる回数が多いものを求めよ
- 答えは必ず 1 つになる

- 制約
- 1 <= N <= 100

```
7
=> 4

25
=> 16 (2^4)

1
=> 1
```

## ロジック

各数字を再帰(ループでもいい)で 2 で割り、その回数で返す
気をつけるべき点は、入力値 1 の時のケース
1 は 2 で割り切れないが、出力例は 1 になっている
この点は、回答を扱う変数の初期値を 1 にして対応する

回答で出力される値は、2^n の値
1 <= N の範囲内の数値のうち最も大きい 2^n の値が一意で求められる

### pseudo code

```
var n
var max, answer

for i=1 to n
    var count
    divide(i, &count)
    if tmp >= max
        max = tmp
print(max)

func divide(num, &count)
    if num % 2 !== 0
        return 0
    if num == 2
        *count++
        return 1
    *count++
    return divide(num / 2)
```

### tips

time complexity はO(n log(n))

別解として、回答が必ず 2^n になる点を利用して以下のようなロジックも組める

```
var n
var a = 1
for i=1 < n {
    i *= 2
    if i <= n
        a = i
}
print(i)
```

2 を二乗した値と N を比較する処理をループ
i * 2 が N より低いうちは、回答値として保存
この場合のtime complexity は　O(log n) となる
