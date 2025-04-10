# Add Two Numbers
https://leetcode.com/problems/add-two-numbers/

ダミーヘッドあり

## 解説

### 上がり
2つのノードを加算した際の桁上がり（キャリー）を考慮する必要があります。キャリーは次の桁の加算に影響を与えます。

### 異なる長さのリスト
2つのリストの長さは異なる場合があります。片方のリストが終了しても、もう片方のリストのノードが残っているか、キャリーが残っている場合、その処理を続ける必要があります。

### 新しいリストの生成
加算結果を保存する新しいリストを生成する際、効率的にノードを追加するためのポインタを保持しておくことが重要です。

## 計算量

### 時間計算量
主に2つのリストを1回ずつ走査します（同時に走査）。
したがって、時間計算量は O(max(m,n)) です。ここで、m と n はそれぞれ l1 と l2 の長さを表します。

### 空間計算量
新しいリストを生成するためのスペースが必要です。最悪の場合、桁上がりが最後まで発生するため、出力リストの長さは入力リストの最大の長さ+1となります。したがって、空間計算量も O(max(m,n)) です。

## 逆順
連結リストを使用する場合、先頭に新しいノードを追加するのは簡単ですが、末尾に追加するのは難しいことが多いです（特に単方向の連結リストの場合）。数値を逆順に格納することで、低い桁（1の位）から順にノードを連結リストに追加することが容易になります。

例えば、数値 342 を連結リストに変換する場合、最初に 2 を持つノードを作成し、次に 4 を持つノードを追加し、最後に 3 を持つノードを追加します。この操作は先頭から順に追加しているため、効率的に行えます。


逆順にする主な意味は、このような連結リストの特性を利用して、数値の各桁を効率的にリストに追加するためです。

この問題においては、先頭に追加していないのでこのメリットは当てはまらない。
1の位から足し合わせて計算を行うのをより直感的にするため、逆順にして1の位を前に持ってくる。

## 連結リスト
連結リストはデータ構造の一つで、データの要素がノードとして連結されている特徴を持っています。以下は連結リストの主な特徴と、それに伴う利点と欠点です。

### 基本的な特徴:
- ノード: 各要素はノードとして格納され、ノードはデータの値（value）と次のノードへの参照（またはポインタ）を持ちます。
- 先頭: 連結リストは通常、先頭となるノードを持ちます。このノードからスタートして、次のノードへの参照を辿ることで、リスト全体を走査できます。
- 方向性: 連結リストには「単方向連結リスト」と「双方向連結リスト」の2つの主なタイプがあります。単方向の場合、各ノードは次のノードへの参照のみを持つのに対し、双方向の場合、前のノードと次のノードの両方への参照を持ちます。

### 利点:
- 動的サイズ: 配列とは異なり、連結リストのサイズは固定されておらず、ランタイム中にノードの追加や削除が容易です。
- メモリ効率: 必要に応じてノードを追加できるため、連結リストはメモリの確保と解放が柔軟です。
### 欠点:
- ランダムアクセス: 配列とは異なり、連結リストの特定の要素にランダムにアクセスすることは効率的ではありません。要素へのアクセスは先頭から順にノードを辿る必要があるため、O(n) の時間がかかります。
- メモリオーバーヘッド: 各ノードはデータの値だけでなく、次のノード（および場合によっては前のノード）への参照も持っているため、メモリのオーバーヘッドが発生します。

総じて、連結リストは要素の追加や削除が多いシナリオや、固定のサイズが不明な場合に有用ですが、ランダムアクセスや頻繁な読み取り操作が必要な場合には、他のデータ構造（例えば、配列や動的配列）がより適している場合があります。
