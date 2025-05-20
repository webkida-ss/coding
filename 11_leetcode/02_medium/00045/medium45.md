# Jump Game II
https://leetcode.com/problems/jump-game-ii/

※ 難しい

## 解説
この問題の解法は、貪欲法（Greedy Algorithm）を使用しています。貪欲法は、各ステップで最適な選択を行うことで全体の最適解を見つけるアルゴリズムです。

このコードでは、各ステップで最も遠くにジャンプできる位置（currentFarthest）を探し、その位置までジャンプすることを繰り返します。そして、そのジャンプ数（jumps）をカウントします。

具体的には、以下のような手順で解いています：

1. 各ステップで、現在の位置からジャンプできる最も遠い位置を探します（currentFarthestを更新）。
2. 現在の位置が前回のジャンプで到達できた最も遠い位置（currentEnd）に達したら、ジャンプ数を1増やします（jumps++）。
3. その後、次のジャンプの終点をcurrentFarthestに更新します（currentEnd = currentFarthest）。


## 計算量

### 時間計算量: O(n)
各要素を一度だけ見て、最も遠くにジャンプできる位置を探すからです。ここで、nは配列の長さを表します。

### 空間計算量: O(1)
このアルゴリズムは定数の追加空間しか必要としないからです（jumps, currentEnd, currentFarthestの3つの変数のみを使用）。