# Pow(x, n)
https://leetcode.com/problems/powx-n/


## 解説
この問題のポイントは、効率的にべき乗を計算する方法を理解することです。具体的には、再帰と分割統治法を使用して、計算量を減らすことが重要です。

### 再帰
べき乗の計算を再帰的に行います。つまり、pow(x, n)はx * pow(x, n-1)と表現できます。しかし、この方法だけでは計算量が大きくなりすぎるため、さらに効率化する必要があります。

### 分割統治法
べき乗の計算を半分に分割して行います。つまり、nが偶数の場合、pow(x, n)はpow(x*x, n/2)と表現できます。これにより、計算量が大幅に減少します。

この式x^n = (x^2)^(n/2)は、べき乗の法則に基づいています。この法則によれば、(a^m)^n = a^(m*n)です。したがって、x^nは(x^2)^(n/2)と等価になります。

この式を利用することで、計算量を半分にすることができます。具体的には、nが偶数の場合、x^nを直接計算するのではなく、(x^2)^(n/2)を計算します。これにより、再帰の深さ（つまり、計算量）が半分になります。

例えば、x^4を計算する場合、直接x*x*x*xを計算するのではなく、(x^2)^2を計算します。これにより、乗算の回数が2回から1回に減少します。

※ 3の倍数や4の倍数でも根本的には同じで、分割統治法のテクニックは活用できるがきりがないあまりやらない

### 負のべき乗
nが負の場合、pow(x, n)は1 / pow(x, -n)と表現できます。これにより、負のべき乗も効率的に計算できます。


## 計算量

### 時間計算量: O(log n)
理由は、各再帰ステップで問題のサイズ（この場合はべき乗の指数n）が約半分になるからです。具体的には、nが偶数の場合、pow(x, n)はpow(x*x, n/2)となり、nが奇数の場合でも、pow(x, n-1)となるため、問題のサイズは約半分になります。　　
したがって、このアルゴリズムは高速であり、大きなnに対しても効率的に動作します。

### 空間計算量: O(log n)
理由は、再帰の深さが問題のサイズ（この場合はべき乗の指数n）の対数に比例するからです。具体的には、各再帰ステップで問題のサイズが約半分になるため、再帰の深さはlog nとなります。　　
したがって、このアルゴリズムはメモリ効率も良く、大きなnに対しても効率的に動作します。