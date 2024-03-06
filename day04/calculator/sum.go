package calculator

// go不会提供publi或者privite
//如需将某些内容设为专用内容，请以小写字母开始。
//如需将某些内容设为公共内容，请以大写字母开始。

var logMessage = "[LOG]"

var Version = "1.0"

func internalSum(number int) int  {
	return number - 1
}

func Sum(number1, number2 int) int  {
	return number1 + number2
}