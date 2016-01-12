package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	om := &OrderMarket{}
	om.Init()

	cm := &CoderMarket{}
	cm.Init()

	totalMoney := 100.0 //单位：w
	noOrderTime := 0 //连续未接到订单的时间统计
	isHavaCoder :=false

	var coder Coder //现成的程序员
	coder = nil
	basicSalary := float64(0) //程序员对应的基本工资

	num := 0 //循环次数

	for {
		if totalMoney > 1000.0 || totalMoney < 0 {
			fmt.Println("创业结束. totalMoney:%v", totalMoney)
			return
		}

		num++
		fmt.Println("新一轮工作开始~~~~~~~~第", num, "轮")
		order, err := om.GetOrder()
		if err != nil {
			fmt.Println("没有拉到订单，休息一天后继续")
			time.Sleep(time.Second * 1) //1秒模拟1天
			totalMoney -= 0.1
			noOrderTime +=1
			continue
		}
		noOrderTime = 0

		if noOrderTime >= 90 {
			fmt.Println("连续3个月未拉倒订单，创业失败!")
			return
		}


		if isHavaCoder == false {
			fmt.Println("没有现成的程序员,需要从人才市场中去拉")
			coderGet,basicSalaryGet, err := cm.GetCoder()
			if err != nil {
				fmt.Println("没有找到程序员，退回订单后继续")
				om.ReturnOrder(order)
				totalMoney -= order.Money * 2
				continue
			}
			coder = coderGet
			basicSalary = basicSalaryGet
		}


		fmt.Println("订单要求的完成时间为:", order.DeadLine)
		software, err := coder.WriteCode(order)
		if software.line > order.DeadLine {
			fmt.Println("订单延期,退回订单,赔付业务方2倍的违约金,并支付程序员基本工资")
			om.ReturnOrder(order)
			totalMoney -= order.Money * 2

			moneyToCoder := basicSalary
			coder.GetMoney(moneyToCoder)
			totalMoney -= basicSalary

			fmt.Println("创业总资金剩余:", totalMoney)
			continue
		}


		om.DeliverOrder(order, software)
		totalMoney += order.Money

		moneyToCoder := order.Money / 10.0 + basicSalary
		coder.GetMoney(moneyToCoder)
		totalMoney -= moneyToCoder

		left := om.Consult() // 协商结果
		isHavaCoder = left

		fmt.Println("创业总资金剩余:", totalMoney)
		fmt.Println("给程序员买盒饭闪了腰，休息两天\n")
		totalMoney -= 0.1 * 2
		time.Sleep(time.Second * 2) //1秒模拟1天
		noOrderTime = 0
	}
}

type Order struct {
	Money    float64 //w
	Name     string
	DeadLine int
}

type OrderMarket struct {
	orderPool []Order
}

func (om *OrderMarket) Init() {

	rand.Seed(time.Now().Unix())
	var orderPool []Order
	for i := 0; i < 10000; i++ {
		order := Order{
			Money: float64(rand.Int()%10 + 10),
			Name:  "order:" + strconv.Itoa(i),
			DeadLine: int(rand.Int()%2 + 10),
		}
		orderPool = append(orderPool, order)
	}

	om.orderPool = orderPool
}

func (om *OrderMarket) GetOrder() (order Order, err error) {
	index := rand.Int() % len(om.orderPool)
	order = om.orderPool[index]
//	om.orderPool = append(om.orderPool[:1], om.orderPool[:1]...)
	fmt.Printf("get an order:%v money:%vw\n", order.Name, order.Money)
	return
}

func (om *OrderMarket) ReturnOrder(order Order) {
	om.orderPool = append(om.orderPool, order)
	fmt.Println("ret an no complete order:", order.Name)
}

func (om *OrderMarket) DeliverOrder(order Order, sf SoftWare) {
	fmt.Println("deliver order success. software info:", sf.Name)
}

func (om *OrderMarket) Consult() (left bool) {
	left = bool(rand.Intn(3)%3 == 0)
	if left == true{
		fmt.Println("老A与程序员协商的结果是:留下来")
	} else {
		fmt.Println("老A与程序员协商的结果是:滚蛋")
	}

	return
}

type CoderMarket struct {
	coderPool []Coder
}

func (cm *CoderMarket) Init() {

	gc := &GoCoder{
		Name: "xiaomin",
	}
	cm.coderPool = append(cm.coderPool, gc)

	gc = &GoCoder{
		Name: "laoA",
	}
	cm.coderPool = append(cm.coderPool, gc)

	dc := &Dentist{
		Name: "yu_ba_ya",
	}
	cm.coderPool = append(cm.coderPool, dc)
}

func (cm *CoderMarket) GetCoder() (coder Coder,basicSalary float64, err error) {
	basicSalary = float64(rand.Int()%10 + 1)
	index := rand.Int() % len(cm.coderPool)
	coder = cm.coderPool[index]
	fmt.Println("success get a order.")
	return
}

type SoftWare struct {
	Name string
	line int
}

type Coder interface {
	WriteCode(order Order) (SoftWare, error)
	GetMoney(money float64)
}

type GoCoder struct {
	Name   string
}

func (gc *GoCoder) WriteCode(order Order) (sf SoftWare, err error) {
	line := int(rand.Int()%2 + 10)
	fmt.Printf("begin to write code for %v\n", order.Name)
	time.Sleep(time.Second * 3)
	fmt.Println("write code ok.")
	fmt.Println("本次开发周期为:", line)
	sf.line = line

	sf.Name = fmt.Sprintf("hello, world! by %v", gc.Name)
	return
}

func (gc *GoCoder) GetMoney(money float64) {
	fmt.Printf("%v get money:%vw\n", gc.Name, money)
}

type Doctor interface {
	Cure() (err error)
	GetMoney(money float64)
}

type Dentist struct {
	Name   string
}

func (d *Dentist) Cure() (err error) {
	//拔牙
	return
}

func (d *Dentist) GetMoney(money float64) {
	fmt.Printf("%v get money:%vw\n", d.Name, money)
}

func (d *Dentist) WriteCode(order Order) (sf SoftWare, err error) {
	line := int(rand.Int()%2 + 10)
	fmt.Printf("i'm a dentist. begin to write code for %v\n", order.Name)
	time.Sleep(time.Second * 3)
	fmt.Println("write code ok.")
	fmt.Println("本次开发周期为:", line)
	sf.line = line

	sf.Name = fmt.Sprintf("hello, world! by dentist:%v", d.Name)
	return
}
