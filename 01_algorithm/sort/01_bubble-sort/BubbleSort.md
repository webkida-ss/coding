# バブルソート
https://www.youtube.com/watch?v=WY-_iKLNhA0

## 概要
左から1つずつ左右を比較していき、大きい方を右に移動させることで、最大値を右に移動させるアルゴリズム。
1度まで行ってからは、最大値が確定するため、最大値を除いて同様の処理を繰り返す。

## 処理フロー

```
2,6,7,3
↓
2,6,3 ,7
↓
2,3, 6,7
↓
2,3,6,7
```