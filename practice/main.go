// GOの文法をテストする

package main //このファイル内をimportする時の名前になる．

import (
	"fmt" // 命名規則
	"regexp"
	"time"

	"./mylib"
)

// 変数でも関数でも，
// GOでは，アンダースコアではなく，MixedCapsかmixedCapsとして書かなければならない．
//　ここで，頭文字が大文字の変数や関数は，外部ファイルから呼び出すことができる．(publicになる)
// 頭文字が小文字の場合は，purivateとなる．

//////////////////////// 変数 //////////////////////////////
// var 変数名 変数の型
// のように記述する
// varによって変数を宣言し、型の明示が必要である。(ex1)
// 初期値を渡した状態で変数を宣言すると型の明示を省略が可能。(ex2)
// 関数内では:=を利用することでより短いコードで変数の宣言を行うことが可能。(ex3)
// >関数内では，変数を宣言したのに使わないとエラーになる．
//  グローバル変数は問題ない．
var text1 string //ex1
var text2 = "文字" //ex1

func testVar() {
	text3 := "JS" //ex3
	var text4 = "lalala"
	fmt.Println(text1, text2, text3, text4)
}

// 定数
// 以下のように記述する
// const 変数名 = 値
// 定数の場合は，使用しなくてもエラーにならない
// 定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使用可能。

func testConst() {
	const text1 = "定数"
	// const text2 string//これはNG
	const text3 string = "aaaa" //これはOK
	fmt.Println(text1)
}

// 関数
// 関数は，以下のように記述する．
// func  <関数名>([引数]) [戻り値の型] {
//     [関数の本体]
// }

func printTest(text string) {
	fmt.Printf(text)
}

// 出力は，型を指定
func getText() string {
	return "lalala"
}

// 複数出力波動する？
func getTextMulti() (string, string) {
	return "lalala", "hahaha"
}

// // 関数のオーバーロードはできるか？>>できない
// func getText(text string) string {
// 	return text
// }

func testFunc() {
	text0 := getText()
	text1, text2 := getTextMulti()

	fmt.Println(text0)
	fmt.Println(text1, text2)
}

// 条件分岐
// if文
// if文の()は許可されない

func testIf() {
	var num int = 2
	if num == 1 {
		fmt.Printf("分岐１")
	}
	if num == 2 {
		fmt.Printf("分岐2")
	}
}

// Switch文
// breakは必要ない

func testSwitch() {
	var num int = 3
	switch num {
	case 1:
		fmt.Printf("分岐1")
	case 2:
		fmt.Printf("分岐2")
	default:
		fmt.Printf("その他")
	}

}

// Defer（遅延実行）
// その関数の終わりで実行される関数を設定する機能
func testDefer() {
	defer fmt.Printf("end")        //これが最後に実行される
	defer fmt.Printf("before end") //最後から二番目
	fmt.Printf("１")
	fmt.Printf("2")
	fmt.Printf("3")
	fmt.Printf("4")
	fmt.Printf("5")

}

// ポインタとアドレス
// 変数に&をつけると変数のアドレスを取得可能

func testPointer() {
	var lang string
	lang = "Go"
	fmt.Println(&lang) //=> 0x1040c128

	// ポインタにアドレスを格納
	var lang_p = &lang
	fmt.Println(*lang_p) //変数の中身を確認
	fmt.Println(lang_p)  // アドレスを確認

	// 値の書き換え
	*lang_p = "Go To Future"
	fmt.Println(lang)
}

/////////////////////////// 配列 //////////////////////////////////
// 配列とは、同じ型を持つ値（要素）を並べたもの。(ex1)
// 複数の宣言方法がある。(ex2)
// 最初に宣言した配列のサイズを変えることはできない。(ex3)

func testArray() {
	//宣言方法(1)
	var arr1 [2]string
	//宣言方法(2)
	var arr2 [2]string = [2]string{"Golang", "Ruby"}
	//宣言方法(3)
	var arr3 = [...]string{"Golang", "Ruby"}

	arr1[0] = "Golang"
	arr1[1] = "Ruby"
	// 配列を直接表示可能
	fmt.Println(arr1, arr2, arr3) //=> [Golang Ruby] [Golang Ruby] [Golang Ruby]

	// 要素の指定
	fmt.Println(arr1[1])

	// アドレスの表示
	fmt.Println(&arr1[1])

}

/////////////////////////// スライス ///////////////////////////
// 配列とは異なり長さ指定の必要なし。(参考コード① ex1)
// 別の配列から要素を取り出し参照する形での宣言やmake()を利用した宣言が可能。(参考コード① ex2)
// 配列とは異なり要素の追加が可能。(参考コード① ex3)
// 長さ(length)と容量(capacity)の両方を持っている。(参考コード② ex4)
// 型が一致している場合、他のスライスに代入することが可能。(参考コード② ex5)
// スライスのゼロ値はnil。(参考コード② ex6)
// appendをした時に容量オーバーを検知すると，自動で元の長さの倍のメモリ容量を確保する．
// ＞予めどのくらい追加されるかが分かっていた場合，多めに容量（makeの第3引数）を設定した方が良い．

func testSlices() {
	//ex1:配列とは異なり長さ指定の必要なし
	var slice1 []string                     //スライス(1)
	var slice2 = []string{"Ruby", "Golang"} //スライス(2)
	//ex2:配列から要素を取り出し参照する形での宣言が可能
	var arr [2]string = [2]string{"Ruby", "Golang"}
	var slice3 = arr[0:2] //スライス(3)

	var slice4 = append(slice3, "JavaScript") //sliceに"Java Script"を追加

	// makeを使った宣言//これは何？
	var slice5 = make([]int, 2, 3) //スライス(4)
	for i := 0; i < 7; i++ {
		if i < 2 {
			slice5[i] = i //はじめに設定したサイズ以下なら可能
		} else {
			slice5 = append(slice5, i) //これは可能
		}
	}

	fmt.Println(slice1, slice2, slice3, slice4, slice5)
	fmt.Println("len:", len(slice5))
	fmt.Println("cap:", cap(slice5))

	// スライスの一部を切り取る
	// 簡易スライス式
	var slice6 = slice5[2:4]
	fmt.Println(slice6, cap(slice6)) //>>要素２なのにcapがでかい．故に容量の無駄
	// 完全スライス式
	// [x:y:z]
	// x:始点y:終点z:メモリ確保の終点
	var slice7 = slice5[2:4:4] //これだとcapは2
	fmt.Println(slice7, cap(slice7))

}

/////////////////////////// MAP（連想配列）///////////////////////////
func testMap() {
	//①組み込み関数make()を利用して宣言
	//make(map[キーの型]値の型, キャパシティの初期値)
	//make(map[キーの型]値の型)
	map1 := make(map[string]string) //①
	map1["Name"] = "Mike"
	map1["Gender"] = "Male"

	//②初期値を指定して宣言
	//var 変数名 map[key]value = map[key]value{key1: value1, key2: value2, ..., keyN: valueN}
	var map2 = map[string]int{"Age": 25, "UserId": 2}

	fmt.Println("map1:", map1)
	fmt.Println("map2:", map2)

}

///////////////////////////　配列要素でのfor///////////////////////////
func testRangeLoop() {
	//スライスとマップの作成
	var slice1 = []string{"Golang", "Ruby", "Python"}
	var map1 = map[string]string{"Lang1": "Golang", "Lang2": "Ruby", "lang3": "Python"}

	//ex1:スライスやマップに使用すると反復毎に2つの変数を返す。
	//ex2:スライスの場合、1つ目の変数は `インデックス(index)`で、`2つ目は要素(value)`である。
	for index, value := range slice1 {
		fmt.Println(index, value)
	}

	for index, value := range map1 {
		fmt.Println(index, value)
	}

	//ex4:インデックスや値は、 _ へ代入することで省略することが可能。
	for _, value := range map1 {
		fmt.Println(value)
		//=> Golang
		//=> Ruby
	}

}

// func main() {
// 	// testArray()
// 	// testSlices()
// 	// testMap()
// 	// testRangeLoop()

// }

/////////////////////////// Structs(構造体)///////////////////////////
// classに似た役割を提供する。(関連する変数をひとまとめにする。)
// typeとstructを使用して定義する。(参考コード① ex1)
// 複数の初期化方法が存在する。(参考コード① ex2)
// 構造体内にメソッドmethodを定義できる。(参考コード② ex3)
// 継承に似た機能として構造体の埋め込みが可能(参考コード③ ex4)

//初期化方法①:変数定義後にフィールドを設定する
// 構造体の定義

type Person struct {
	name string
	age  int
}

//Personのメソッドを定義
// ele Personの部分で，所属先を指定
func (ele Person) intro(arg string) string {
	return arg + " I am" + " " + ele.name
}

// コンストラクタ
// 構造体の機能としてのコンストラクタはない．
// そのため，自分で作る必要がある．
// この書き方が慣例的によく使われるらしい．
// 参考：https://qiita.com/gold-kou/items/4494f8b69b8fa53d5e93
func NewPerson(name string, age int) *Person {
	// 以下の2つはほとんど同じ.2つ目の方がわかりやすい．
	// p := new(Person)
	p := &Person{name: name, age: age}
	// ここで何らかの処理を書けば，一般的なコンストラクタと同じ使い方ができる（ほぼ）

	return p
}

func testPerson() {
	// 初期化方１：インスタンス生成後にプロパティを定義
	// var mike Person
	// mike.name = "Mike"
	// mike.age = 23

	// プロパティ名を指定して初期化
	// var jack = Person{name: "Jack", age: 41}

	// 順番で初期化
	// var jon = Person{"Jon", 41}

	// メゾッドの呼び出し
	// fmt.Println(jon.intro("Hello!"))

	// コンストラクタ で初期化
	jony := NewPerson("jony", 38)
	fmt.Println(jony.intro("Hello!"))
}

//　埋め込み（Embed）
// >Goに継承はないが，埋め込みがある．
type MasterStudent struct {
	*Person
	researchArea string //研究分野
	paperNum     int    //論文投稿数
}

// MasterStudentのコンストラクタ
func NewMasterStudent(researchArea string, paperNum int, name string, age int) *MasterStudent {
	MS := &MasterStudent{
		researchArea: researchArea,
		paperNum:     paperNum,
		Person:       NewPerson(name, age)}
	return MS
}

func testEnbed() {
	aMasterStudent := NewMasterStudent("control theory", 31, "Saburo", 24)
	fmt.Println(aMasterStudent.intro("Hello!"))
}

///////////////////////////インターフェース///////////////////////////
//任意の型が「どのようなメゾッドを実装すべきか」を規程するためのもの．
// ＞構造体やクラス？
// 参考：https://medium.com/since-i-want-to-start-blog-that-looks-like-men-do/%E5%88%9D%E5%BF%83%E8%80%85%E3%81%AB%E9%80%81%E3%82%8A%E3%81%9F%E3%81%84interface%E3%81%AE%E4%BD%BF%E3%81%84%E6%96%B9-golang-48eba361c3b4

type Student struct {
	Name   string
	Number int
	Grade  int
}
type Teacher struct {
	Name string
}

type Worker interface {
	getEmail() string
}

// Worker型のインターフェースを継承するには，StudentとTeacherはgetEmail()を実装する必要がある．
func (s Student) getEmail() string {
	return s.Name + "@student.ed.jp"
}
func (t Teacher) getEmail() string {
	return t.Name + "@teacher.ed.jp"
}

// 完全に共通な処理は，Worker内に実装することで，共通処理として実現することができる．
// >Worker内のクラスとして実現できない？
// >>できない．つまり，次のように書くとエラー
// >>さすがGoogle，分かってる．

// func (p Worker) sendEmail_() string {//インターフェースに関数を実装することはできない
// 	from := p.getEmail()
// 	var context = `送信元 : ` + from + `
// これはテスト用のメールです。
// よろしくお願いします。
// `
// 	return context
// }

func sendEmail(p Worker) string {
	from := p.getEmail()
	context := `送信元 : ` + from + `
これはテスト用のメールです。
よろしくお願いします。
`
	return context
}

func testInterface() {
	var s Worker // 変数sをWorker型で宣言
	var t Worker // 変数tをWorker型で宣言
	s = Student{ //StudentがgetEmailを持たない場合，ここでエラー
		Name:   "Yamada",
		Number: 999,
		Grade:  5,
	}
	t = Teacher{
		Name: "Tsubomi",
	}

	// var emailS = s.getEmail()
	// var emailT = t.getEmail()

	// fmt.Println(emailS)
	// fmt.Println(emailT)

	cxtStu := sendEmail(s)
	fmt.Println(cxtStu)
	cxtTea := sendEmail(t)
	fmt.Println(cxtTea)

}

// ポインタ レシーバと値レシーバ
type Vertex struct {
	X, Y int
}

// 値レシーバ
func (v Vertex) Area() int { //こうするとVertexの一部になる．
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) { //値渡しとポインタレシーバで効果が異なる．
	v.X = v.X * i
	v.Y = v.Y * i
	// 値レシーバとなるかどうかは，関数の書き方だけで決まる？
}

func Area(v Vertex) int { //これをVertexに含めたい
	return v.X * v.Y
}

func testReceiver() {
	v := Vertex{3, 4}
	// fmt.Println(Area(v))

	v.Scale(10)
	fmt.Println(v.Area())
}

// Q1
func (v Vertex) Plus() int {
	return v.X + v.Y
}

func (v Vertex) String() string {
	return fmt.Sprintf("X is %d,Y is %d", v.X, v.Y)
}

// 便利な標準パッケージ
// time
func testTime() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
}

// regex
// 要は，正規表現
func testRegex() {
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	// ms := r.MatchString("apple")
	st := r.FindStringSubmatch("apple")
	// 正規表現がマッチした順番で，スライスに格納される．
	fmt.Println(st[1])
}

// オプショナル引数を持つ関数の作成
type OptVals struct {
	Val string
}

type option func(*OptVals) //optionという，ハンドラ？を作成．引数にGreetOptsを持つ．

// GreetingWord引数を設定する関数
func SetOptVal(v string) option {
	return func(g *OptVals) {
		g.Val = v
	}
}

func Greet(name string, opts ...option) {
	// デフォルトパラメータを定義
	g := &OptVals{
		Val: "Hello",
	}

	// ユーザーから渡された値だけ上書き
	for _, opt := range opts {
		opt(g)
	}

	fmt.Printf("%s, %s!\n", g.Val, name)
}

func testOptVal() {
	Greet("gopher")                   // Hello, gopher!
	Greet("gopher", SetOptVal("Hey")) // Hey, gopher!// 可変長なので，複数渡すこともできる．
}

// ライブラリの自作
// 基本は、packageごとにディレクトリを分ける。
// そのpackage名(ディレクトリ名)によって、中身の関数や構造体を呼び出すことが可能。

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))
}
