Changelog
=========

- バージョン見だしとして、単線の下線を使うようにした (#2)
- 先頭に `Changelog` の二重下線の見出しを挿入するようにした (#2)
- ユーザ名とレポジトリ名が省略された時、.git/config の情報を使うようにした (#3)

v0.3.0
------
Jan 7, 2025

- `-r REVISION` オプションを追加

v0.2.0
------
May 11, 2024

- 出力から CR を除いた
- 以下の形式をサポート
    - `type-rnote https://github.com/USER/REPO`
    - `type-rnote USER/REPO`
    - `type-rnote USER REPO`

v0.1.1
------
Apr 10, 2023

- `name` 項目が空の場合は `tag_name` を見出しに使うようにした

v0.1.0
------
Apr 10, 2023

初版
