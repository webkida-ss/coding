#!/bin/bash

# 使い方
# bash setup.sh m 53

# 引数からdegreeとnoを受け取る
degree=$1
no=$2

# degreeに応じたディレクトリ名と文字列を設定
case $degree in
  e)
    dir="01_easy"
    degree_str="easy"
    ;;
  m)
    dir="02_medium"
    degree_str="medium"
    ;;
  h)
    dir="03_hard"
    degree_str="hard"
    ;;
  *)
    echo "Invalid degree: $degree"
    exit 1
    ;;
esac

# noを5桁にパディング
padded_no=$(printf "%05d" $no)

# コピー先のパスを設定
destination="11_leetcode/${dir}/${padded_no}"

# templateフォルダをdestinationにコピー
cp -r "template" "$destination"

# コピーしたフォルダ内のファイルをリネーム
for file in "$destination"/*; do
  # ファイルの拡張子を取得
  extension="${file##*.}"
  # 新しいファイル名を生成
  new_filename="${destination}/${degree_str}${no}.${extension}"
  # ファイルをリネーム
  mv "$file" "$new_filename"
done