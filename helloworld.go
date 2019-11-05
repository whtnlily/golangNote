package main

import (
	"bytes"
	"errors"
	"log" //log打印比fmt打印会打印出日期和时间
	"reflect"
	"strconv"
	"time"
	"unicode/utf8"
)
import "./common/config"
import "./winterface"
import "./bean"
import "./manager"

var age int
var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//切片与数组的区别：不指定元素个数的数组就是切片。否则就是数组
var wlist []string

//数组
var myWife [20]string

var caim Wife

const (
	CBS = 83
	PI = 3.1415926
	AB  //不赋值，默认是继承上一项值3.1415926
)

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"
/**主函数*/
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	vartest()

	pointPointTest()

	structTest()

	arraytest()

	rangeTest()

	mapTest()

	funtest()

	interfaceTest()

	structTest1()

	channelTest()

	channelTest1()

	chanTest()

	stringTest()

	slinceTest()

	chanTimeout()

	checkChannelIsFull()

	checkType()

	Defertest()

	Timeout()

	reflectest(caim)
}

func funtest() {
	log.Println(fibonacci(10))
	log.Println("int->float32:", float32(16)/float32(3))
	printLine()
}

func arraytest() {
	log.Println(Wife{"caim", 26, []int{115, 62, 120}, 9})
	caim = Wife{"caimin", 26, []int{112, 60, 118}, 9}
	log.Println(caim)
	printLine()
	log.Println(append(wlist, "zuoyq", "wenyj", "liliq", "zhouyp", "chenzy", "dingx"))
	printLine()
	log.Println(wlist[2:5])
	//切片
	printLine()

}

func vartest() {
	log.Println("const default PI :", PI, ",AB:", AB)
	printLine()
	wlist = []string{"hsn", "caim", "liut", "haowq", "xiaoy"}
	/**/
	log.Println("Hello World!")
	printLine()
	test()
	g, h := 123, "hello"
	log.Println(x, y, a, b, c, d, e, f, g, h, config.PackageName)
	log.Println(config.NetworkType, ",getNetworkType:", config.GetNetworkType())
	log.Println("iota:", config.A, config.B, config.C, config.D, config.E, config.F, config.G, config.H, config.I)
	log.Println("iota1:", config.J, config.K, config.L, config.M)
	printLine()
	pointTest()
	printLine()
	log.Println(wlist)
	printLine()
	wlistfor(wlist)
	printLine()
	//切片末尾增加元素
	wlist = append(wlist, "zhanghj")
	wlistfor(wlist)
	printLine()
	//在切片头增加元素
	wlist = append([]string{"weiy"}, wlist...)
	wlistfor(wlist)
	printLine()
}

func structTest1() {
	liutsan := bean.SanWei{76, 58, 65}
	var liutwife = bean.WifeBean{
		"liut",
		28,
		8.2,
		liutsan,
	}
	log.Println("wife liut info:", liutwife)
	manager.AddWife(liutwife.Name, liutwife)
	printLine()
}

/**：=赋值运算符*/
func test() {
	age = 26
	name := "husn"
	level := 9
	log.Println("age =", age, ",name =", name, ",level =", level)
	printLine()
}

/*指针*/
func pointTest() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int //指针类型

	/* 运算符实例 */
	log.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	log.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	log.Printf("第 3 行 - c 变量类型为 = %T\n", c)
	log.Printf("第 4 行 - ptr 变量类型为 = %T\n", ptr)

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址，&取a的地址 */
	log.Printf("a 的值为  %d\n", a)
	log.Printf("*ptr指针为a的地址  %d\n", ptr)
	log.Printf("*ptr = &a 为 %d\n", *ptr) //*代表取ptr这个地址的值

}

/*for循环*/
func wlistfor(mylist []string) {
	for i := 0; i < len(mylist); i++ {
		log.Println("mywf", i, ":", mylist[i])
	}
}

/**指针的指针*/
func pointPointTest() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a //取a的地址

	/* 指向指针 ptr 地址 */
	pptr = &ptr //取ptr的地址

	/* 获取 pptr 的值 */
	log.Printf("变量 a = %d\n", a)
	log.Printf("指针变量 *ptr = %d\n", *ptr)
	log.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
	printLine()
}

/**结构体，可以看成java中的类*/
type Wife struct {
	Name  string
	Age   int
	Sanw  []int
	Level int
}

/**结构体*/
func structTest() {
	myW := Wife{"husn", 30, []int{98, 60, 110}, 9}
	log.Println(myW)
	log.Println("husn sanw:", myW.Sanw)
	printLine()
}

/**range测试，类似python中的in*/
func rangeTest() {
	myWife = [20]string{"husn", "caim", "liut", "haowq", "xiaoy", "zuoyq", "wenyj", "liliq", "zhouyp", "chenzy", "dingx", "weiy", "huyy", "yangxx", "gongyl", "liulm"}
	for i, wifi := range myWife {
		log.Println("wifi", i, wifi)
	}
	printLine()
}

/**map集合*/
func mapTest() {
	wifelist := make(map[string]float32)
	wifelist["husn"] = 9.8
	wifelist["caim"] = 9.2
	wifelist["xiaoy"] = 7.8
	wifelist["liut"] = 8

	log.Println("map test :", wifelist)
	for name, level := range wifelist {
		log.Println(name, level)
	}
	printLine()
}

/**递归函数*/
func fibonacci(n int) int {

	if n < 2 {
		return n
	}
	if n == 10 {
		log.Println(errors.New("params error"))
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

type Husn struct{}

//(husn Husn)说明该方法是husn所有，
func (husn Husn) MaL() {
	log.Println("husn mal with wanht")
}

type Caim struct{
	level int
}

func (caim Caim) MaL() {
	log.Println("caim mal with wanht")
	caim.level = 9
	log.Println("caim level:",caim.level)
}

func (husn Husn) Merry() {
	log.Println("husn merry with wanht")
}

func (caim Caim) Merry() {
	log.Println("caim merry with wanht")

}

/**接口定义*/
var wife Winterface.IWife

func interfaceTest() {
	wife = new(Caim)
	wife.Merry()
	wife.MaL()

	wife = new(Husn)
	wife.Merry()
	wife.MaL()

	printLine()
}

/*
*channel管道
* 1、带缓冲区的是异步的，不带缓冲区的是同步的
* 2、
 */
func channelTest() {
	var ch = make(chan int, 1) //带缓冲区的channel是异步的 make(chan int,2)

	go func() {
		ch <- 2
		//log.Println("channel 1 <- :",<-ch)
		//log.Println("channel 2 <- :",<-ch)
	}()
	//ch <- 2
	//ch <- 3
	////log.Println(<-ch)
	////data := <-ch
	log.Println("channel test:", <-ch)
	printLine()
}

/**channel管道*/
func channelTest1() {
	ch := make(chan bool) //不带缓冲区的channel是同步的。同步的情况下，channel在的发送和接收不能在同一个协程中。
	var data bool
	go func() {
		log.Println("come into goroutine")
		ch <- true //这里发送1
		log.Println("come into goroutine1")
		//下面的代码不会执行，因为在该协程中，ch接收了数据，必须要等待另外一个协程去接收
		data = <-ch //这里接收2
		log.Println("come into goroutine2 data:", data)
	}()

	select {
	case data = <-ch: //这里接收1
		log.Println("come into goroutine data:", data)
	}

	log.Println("do something else：", data)

	ch <- false //这里发送2
	//<-ch
	//log.Println("ch data:",<-ch)  //上面已经处理了<-ch，这里的<-ch会导致死锁
	close(ch)

	log.Println("end channel")
	printLine()
}

var cha = make(chan int)

func foo(i int) {
	cha <- i
}

func chanTest() {
	// 开启5个routine
	for i := 0; i < 5; i++ {
		go foo(i)
	}

	// 取出信道中的数据
	for i := 0; i < 5; i++ {
		log.Println(<-cha)
	}
	printLine()
}

func stringTest() {
	log.Println("中文字符串的个数:", utf8.RuneCountInString("中文字符串的个数"))
	title := "zhanghj"
	sub := " ily"
	thr := " ify"

	//字符串拼接
	var stringBuidler bytes.Buffer
	stringBuidler.WriteString(title)
	stringBuidler.WriteString(sub)
	stringBuidler.WriteString(thr)

	log.Println(stringBuidler.String())

	//字符类型
	var hj rune
	hj = '惠'
	log.Println("字符类型rune:", hj)
	printLine()
}

func printLine() {
	log.Println("=====================================")
}

//切片用法
func slinceTest() {
	log.Println("切片删除头元素1")
	//删除元素,删除头元素
	wlist = wlist[1:]
	wlistfor(wlist)
	printLine()
	//删除开头元素方案2:append
	log.Println("切片删除头元素append")
	wlist = append(wlist[:0], wlist[1:]...)
	wlistfor(wlist)
	printLine()
	//删除开头的元素方案3，copy(目标切片，源切片)返回目标切片的长度
	log.Println("切片删除头元素copy")
	wlist = wlist[:copy(wlist, wlist[1:])]
	wlistfor(wlist)
	printLine()
	//删除中间的元素::append第一个参数的的1代表要删除的元素
}

/**
* 1、chan的用法
* 2、select的用法：类似于switch，但是select的case只能是I/O操作，多与channel配合使用
* 3、超时的设计
 */
func chanTimeout() {
	timeout := make(chan bool, 1) //带缓冲，异步
	go func() {
		time.Sleep(3e9) // sleep 3 seconds
		timeout <- true
	}()
	ch := make(chan int)
	select {
	case <-ch: //因为ch没有输入值，这里会读取失败。然后走下一个case，但是这个时候timeout也没有输入，还是会读取失败，并且select没有default，所以会一直阻塞在这里，这个时候3秒过后，协程中timeout被传入了值。所以走打印
	case <-timeout:
		log.Println("timeout!")
	}
	printLine()
}

/**
* 使用select判断channel是否存满
 */
func checkChannelIsFull() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	select {
	case <-ch1:
		log.Println("ch1 pop one element")
	case <-ch2:
		log.Println("ch2 pop one element")
	default: //如果case1、case2均未执行，则说明ch1&ch2已满
		log.Println("default,ch1&ch2 all are full")
	}
	printLine()

	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default: //因为 ch 插入 1 的时候已经满了， 当 ch 要插入 2 的时候，发现 ch 已经满了（case1 阻塞住）， 则 select 执行 default 语句。 这样就可以实现对 channel 是否已满的检测， 而不是一直等待。
		log.Println("channel is full !")
	}
	printLine()
}

/**
* 判断数据类型方法
*/
func checkType() {
	Params := make([]interface{}, 3)
	Params[0] = 88                   // 整型
	Params[1] = "wanht lmf all"         // 字符串
	Params[2] = Wife{Name:"xxx",Age:26,Sanw:[]int{126,73,88},Level:9} // 自定义结构体类型

	// Comma-ok断言
	for index, v := range Params {
		if _, ok := v.(int); ok {
			log.Printf("Params[%d] 是int类型 \n", index)
		} else if _, ok := v.(string); ok {
			log.Printf("Params[%d] 是字符串类型\n", index)
		} else if _, ok := v.(Wife); ok {
			log.Printf("Params[%d] 是自定义结构体类型\n", index)
		} else {
			log.Printf("list[%d] 未知类型\n", index)
		}
	}

	printLine()
	// switch判断
	for index, v := range Params {
		switch  value := v.(type) {
		case int:
			log.Printf("Params[%d] 是int类型, 值：%d \n", index,value)
		case string:
			log.Printf("Params[%d] 是字符串类型, 值：%s\n", index,value)
		case Wife:
			log.Printf("Params[%d] 是Wife类型, 值：%s\n", index, value.Name+","+strconv.Itoa(value.Age))
		default:
			log.Printf("list[%d] 未知类型\n", index)
		}

	}

	printLine()
}

/**defer*/
func Defertest() (i int) {
	defer func() {
		log.Println("defertest1 i:",i)
		i++
		log.Println("defertest2 i:",i)
		printLine()
	}()
	return 1
}

func Timeout()  {
	ch := make(chan int)
	select {
	case <-ch:
	case <-time.After(time.Second*1):
		log.Println("超时测试：超时了")
	}
	printLine()
}

/**反射*/
func reflectest(mWife Wife)  {
	//获取目标对象
	t := reflect.TypeOf(mWife)
	log.Println("<<<reflect t type,",t.Name())
	//获取目标对象的值类型
	v := reflect.ValueOf(mWife)
	for i := 0;i < v.NumField();i++{
		key := t.Field(i)
		value := v.Field(i).Interface()
		log.Println("<<<第",i+1,"个字段是：",key.Name,":",key.Type,"=",value)
	}
	printLine()

	// 通过.Kind()来判断对比的值是否是struct类型
	if k := t.Kind(); k == reflect.Struct {
		log.Println("<<<struct类型判断")
	}

	num := 1;
	numType := reflect.TypeOf(num)
	if k := numType.Kind(); k == reflect.Int {
		log.Println("<<<数值类型判断")
	}
	printLine()
}