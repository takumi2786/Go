package main

import (
	"./stock"
	"github.com/markcheno/go-quote"
)

// GOによる株価の取得
// spyとは？
// SPYは、アメリカの指数連動型投資信託の一種で、株価指標であるS&P500の株価指数に連動するように組成されています。
// 要は，日経平均みたいなもの？
// rsiとは？
// 相場の売られすぎ，買われすぎなどをはかる指標．
// 低いほど〜，高いほど〜

// Goによるグラフの描画
// https://qiita.com/RuyPKG/items/0a569953e9e24870f527

func main() {
	// Yahooのサイトからspyの値をスクレイピングで取得する関数
	// quote, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-12-01", quote.Daily, true)
	// CSVからの読み込み
	quote, _ := quote.NewQuoteFromCSVFileDateFormat("test", "./stock_data/EskeyEL_2020.csv", "2006-01-02")

	// fmt.Println(quote.CSV())

	// 棒グラフの作成
	dateArr := quote.Date   //日付
	OpenArr := quote.Open   //開始時の値段
	dataNum := len(dateArr) //データ数

	// UNIX時間に変換
	dateArrUN := make([]float64, dataNum)
	for i := 0; i < dataNum; i++ {
		dateArrUN[i] = float64(dateArr[i].Unix())
	}

	// タイムスタンプを文字列に変換
	dateArrST := make([]string, dataNum)
	const layout = "2006/01/02" //これは，この日付じゃないとバグる．
	for i := 0; i < dataNum; i++ {
		if i%30 == 0 {
			dateArrST[i] = dateArr[i].Format(layout)
		} else {
			dateArrST[i] = ""
		}
	}

	stock.PlotLine(dateArrUN, OpenArr, dataNum, dateArrST)
}
