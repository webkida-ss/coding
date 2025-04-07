# Longest Common Prefix
https://leetcode.com/problems/longest-common-prefix/


## ポイント

### 解法1：前から比較
文字列比較より前方から文字単位で比較した方が、文字列生成が少なくて済む。

### 解法2：削る方式
減点法的な考えでやるときは、接尾辞から削っていくイメージ。
今回の例でいうと、共通部分は基本的にはないので、どんどん削っていきあったら残す程度。

## 解説
初期接頭辞の選択: 最長の共通接頭辞は、最初の文字列より長くなることはありません。このため、最初の文字列を共通接頭辞のスタートとして選ぶのは合理的です。

接頭辞の縮小: 残りの各文字列を調べる際、現在の接頭辞と一致しない部分があれば、接頭辞を1文字ずつ短くしていきます。これにより、他のすべての文字列にも一致する最長の接頭辞を効率的に見つけることができます。

最小の文字列: 最短の文字列が共通接頭辞の最大の長さを決定します。したがって、最短の文字列を探して、それを初期の接頭辞として選択することも効率的なアプローチです。

## 計算量

### 時間計算量 (Time Complexity)

最悪のケースとして、最も短い文字列の長さを minLen とした場合、minLen 回までの繰り返し処理があります。
それぞれの繰り返しで、配列内の他のすべての文字列に対して indexOf 操作が行われます。indexOf の時間計算量は、比較される文字列の長さに比例するため、最悪のケースではO(n)です (n は文字列の長さ)。
したがって、全体の時間計算量は O(minLen * n * m) となります。ここで、n は最も長い文字列の長さ、m は配列の要素数を指します。

### 空間計算量 (Space Complexity)

このコードにおける追加的なメモリ使用は非常に限られています：

prefix 変数は入力文字列の1つを参照します。
ループ変数や一時的な変数は定数のスペースを使用します。
したがって、空間計算量は O(1) です。

結論:
このコードの時間計算量は O(minLen * n * m)、空間計算量は O(1) です。


## 改善案

「Longest Common Prefix」の問題をさらに最適化するための一般的な方法として、以下のアプローチが考えられます。

### ソートを利用する方法

文字列の配列を辞書順でソートします。そして、最初と最後の文字列のみを比較することで、共通のプレフィックスを探します。ソートされたリスト内で最初と最後の文字列のプレフィックスが最も短くなるため、これら二つを比較するだけで十分です。

```typescript
function longestCommonPrefix(strs: string[]): string {
    if (strs.length === 0) return "";

    strs.sort();

    const first = strs[0];
    const last = strs[strs.length - 1];
    let i = 0;

    while (i < first.length && i < last.length && first[i] === last[i]) {
        i++;
    }

    return first.substring(0, i);
}
```

### 二分探索を利用する方法

最も短い文字列の長さを最大のプレフィックスの長さと考え、二分探索で共通のプレフィックスの長さを探します。

```typescript
function isCommonPrefix(strs: string[], len: number): boolean {
    const str1 = strs[0].substring(0, len);
    for (let i = 1; i < strs.length; i++) {
        if (!strs[i].startsWith(str1)) {
            return false;
        }
    }
    return true;
}

function longestCommonPrefix(strs: string[]): string {
    if (strs.length === 0) return "";

    let minLen = Number.MAX_VALUE;
    for (const str of strs) {
        minLen = Math.min(minLen, str.length);
    }

    let low = 1;
    let high = minLen;

    while (low <= high) {
        const mid = (low + high) >> 1;

        if (isCommonPrefix(strs, mid)) {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }

    return strs[0].substring(0, (low + high) >> 1);
}
```

これらの方法の中でどれがベストかは、入力データや実際の使用ケースによって異なります。たとえば、ソートを利用する方法は、文字列の数が非常に多い場合に有利である可能性があります。一方、二分探索を利用する方法は、文字列の長さが非常に長い場合に有利である可能性があります。