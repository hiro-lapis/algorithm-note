[問題文]()

#dp
#yes/no
#漸化式

## dp,漸化式とは

dp は大きな問題を小さな問題に分割して解き、その小さな回答をメモしておくことで、大きな問題を効率的に解く手法  

例えば、13時に渋谷で用事があるとして、13時に渋谷にいられるか判断するにはどうすればいいか？  
前提として、13時に着く時刻に電車に乗れてなきゃいけない。(12:30とする)
定刻までに電車に乗るには、駅についてないといけない。(12:25)
駅に着けるよう家を出発しないといけない。(12:10とする)
家を出られるよう準備をおえないといけない(12:00とする)

上記の例でいうと、渋谷に時間通りいれるかは12:00に準備を終えられるかどうかにかかっている
これら情報を配列で管理して、効率よく欲しい値を求めるのがdp


漸化式は数式の一つで**数列において前の項の結果から、現在の項の値を求める方法**のこと
