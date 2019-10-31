package config

const PackageName = "ily.hsn.wht"

const (
	A = iota   //0
	B          //1
	C          //2
	D = "ha"   //独立值，iota += 1
	E          //"ha"   iota += 1
	F = 100    //iota +=1
	G          //100  iota +=1
	H = iota   //7,恢复计数
	I          //8
)

const (
	//iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3
	J=1<<iota
	K=3<<iota
	L
	M
)
