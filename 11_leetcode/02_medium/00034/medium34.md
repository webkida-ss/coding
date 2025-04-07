# Find First and Last Position of Element in Sorted Array
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

## 解説
ソートされた整数の配列と目標の値が与えられ、その目標の値が配列内で出現する最初と最後の位置を見つけるというものです。目標の値が配列内に存在しない場合は、[-1, -1]を返します。  
この問題のポイントは、二分探索を使用して目標の値を効率的に見つけることです。二分探索は、ソートされた配列内で特定の値を見つけるための高速なアルゴリズムで、各ステップで配列の半分を無視することができます。

## 計算量

### 時間計算量: O(log n)
二分探索

### 空間計算量: O(1)
このアルゴリズムが定数の追加空間しか必要としないためです。