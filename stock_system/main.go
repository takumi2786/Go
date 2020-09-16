package main

import (
	"github.com/markcheno/go-quote"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
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

type MyPlot struct {
	*plot.Plot
}

func NewMyPlot() (*MyPlot, error) {
	p_, err := plot.New() //ベース構造体のコンストラクタを呼び出し．
	p := &MyPlot{Plot: p_}
	return p, err
}

// 新たな関数を追加．
// 自分で追加した関数．軸の数値と表示を別で指定する場合に使用
func (p *MyPlot) NominalX2(names []string, vals []float64) {
	p.X.Tick.Width = 0
	p.X.Tick.Length = 0
	p.X.Width = 0
	p.Y.Padding = 0
	ticks := make([]plot.Tick, len(names))
	for i, name := range names {
		ticks[i] = plot.Tick{float64(vals[i]), name}
	}
	p.X.Tick.Marker = plot.ConstantTicks(ticks)
}

// x軸データ，y軸データ，x軸データのラベルを渡すと，データをプロット．
// labelArrを入れた場合は，x軸の値の代わりに，ラベルが値としてかかれる．
func plotLine(xArr []float64, yArr []float64, dataNum int, labelArr []string) {
	// インスタンスを生成
	p, err := NewMyPlot()
	if err != nil {
		panic(err)
	}

	pts := make(plotter.XYs, dataNum)
	for i := 0; i < dataNum; i++ {
		pts[i].X = xArr[i]
		pts[i].Y = yArr[i]
	}

	// グラフを描画
	// p.NominalX(labelArr...)
	p.NominalX2(labelArr, xArr)

	err = plotutil.AddLinePoints(p.Plot, pts) //既存の関数を利用する徳は，元のstructを渡さなければならない．
	if err != nil {
		panic(err)
	}
	if err := p.Save(10*vg.Inch, 5*vg.Inch, "out.jpg"); err != nil {
		panic(err)
	}

}

func LoadStockData(mode int) {
	// mode:0 spiを取得．
	// mode:1 CSVから取得．

}

func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-12-01", quote.Daily, true)
	// Yahooのサイトからspyの値をスクレイピングで取得する関数
	// rsi2 := talib.Rsi(spy.Close, 2)
	// fmt.Println(rsi2)

	// 棒グラフの作成
	dateArr := spy.Date     //日付
	OpenArr := spy.Open     //開始時の値段
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

	plotLine(dateArrUN, OpenArr, dataNum, dateArrST)
}
