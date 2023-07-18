# tenkiGetter

天気情報を取得します

[![Coverage Status](https://coveralls.io/repos/github/KaMatsubara/tenkiGetter/badge.svg?branch=main)](https://coveralls.io/github/KaMatsubara/tenkiGetter?branch=main)
[![codebeat badge](https://codebeat.co/badges/4e8d995f-16b5-4bc8-b135-641143c9a467)](https://codebeat.co/projects/github-com-kamatsubara-tenkigetter-main)
[![Go Report Card](https://goreportcard.com/badge/github.com/KaMatsubara/tenkiGetter)](https://goreportcard.com/report/github.com/KaMatsubara/tenkiGetter)
[![DOI](https://zenodo.org/badge/628910827.svg)](https://zenodo.org/badge/latestdoi/628910827)
## Description
気象庁から提供されているAPIを利用してCLIから天気情報を取得するアプリケーションです

## Usage
引数として都道府県名を与えると天気予報を取得できます
"都・道・府・県"は省略できません
```
tenkiGetter [オプション] <場所>

オプション
-h --help ヘルプを表示
-v --version　バージョンを表示
-d --day　短期天気概況の取得
-w --week　週間天気概況の取得
```

## Installation
```
brew install KaMatsubara/tap/tenkiGetter 
```


## About
気象庁のAPIから天気予報を取得しています
