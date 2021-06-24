package main

import "fmt"

func main() {

	intChan := make(chan int, 1000)

	primeChan := make(chan int, 2000)

	//标识退出的管道
	exitChan := make(chan bool, 4)

	//开启一个协程,向intChan放入 1-8000个数
	go putNum(intChan)
	//开启四个协程,从intChan取出数据,并判断是否为素数,如果是素数,就放到primeChan内
	for i := 0; i < 4; i++ {
		go primeChanM(intChan, primeChan, exitChan)
	}
	flag := make(chan bool)
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		flag <- true
		close(exitChan)
	}()
	res, ok := <-flag

	for ok && res {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(res)
	}
}

func primeChanM(intChan chan int, primeChan chan int, exitChan chan bool) {
	//var num int
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			//管道内无值
			break
		}
		flag = true
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				//说明不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	exitChan <- true
}

func putNum(intChan chan int) {
	for i := 0; i <= 8000; i++ {
		intChan <- i
	}
	//关闭管道
	close(intChan)
}
