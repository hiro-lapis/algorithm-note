[問題文](https://atcoder.jp/contests/abc320/tasks/abc320_c)

#tag

## 問題文の把握

- there are 3 slots
- array of `i` th reel has `S` strings
- Si has `M` length strings
- you can stop reel after `t` seconds, or not
- if stop at `t` seconds, then reel stop at `t mod M + 1` string
- you want stop all reel at same strings, as early as possible. get early seconds
- if you cant stop all reel at same strings, output `-1`

- 制約

- M is int
- 1 <= M <= 100

```
case1
10
1937458062
8124690357
2385760149
=> 
2nd reel stop at 0 sec, then show 2 
3rd reel stop at 2 sec, then show 2 
1st reel stop at 6 sec, then show 2 
thus, earlyest time is 6sec

case2
20
01234567890123456789
01234567890123456789
01234567890123456789

case3
5
11111
22222
33333
=> 1

```

## ロジック

slots number range is 0 ~ 9,
so you get each number's earliest align,
and compare to most earliest seconds.  
As a point, **there is cant stop reels at the same time**,   
if there is number align at the same time, you skip this loop 

### pseudo code

```
S = list of slot 1~3
M = length of reel
ANS = INF
for 0 to 3
    for i (0 to 10)
        for j (0 to 10)
            if i == j continue
        for k (0 to 10)
            if i == j || j == k continue
            if S[0][i%m] == S[1][j%m] && S[1][j%m] == S[2][k%m]
                if ans > max(i, j, k)
                    ans = max(i, j, k)

if ans == INF
    ans = -1
print(ans)
```

### tips

time complexity is O(3M) => O(n) 

