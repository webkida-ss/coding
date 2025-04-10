https://leetcode.com/problems/two-sum/
問題:
整数の配列と目標値が与えられ、2つの数の合計が目標値となるような数値のインデックスを見つける。

キーとなるアイディア:

各数値について、目標値からその数値を引いた「補数」を計算する。
この補数が以前に見た数値であるかどうかを効率的に確認するには、ハッシュマップ（辞書）を使用する。
ハッシュマップを使って、既に処理した数値のインデックスを保存しておく。
補数がハッシュマップに存在すれば、答えのペアが見つかったということになる。
このアプローチを採用すると、配列を一度だけスキャンすることで答えを見つけることができ、計算量はO(n)になります。
