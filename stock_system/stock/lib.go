package stock

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Plot structに、独自関数を組み込みにより追加。
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
func PlotLine(xArr []float64, yArr []float64, dataNum int, labelArr []string) {
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
