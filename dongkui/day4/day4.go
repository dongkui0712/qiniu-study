package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"errors"
)

func main() {

	om := &OrderMarket{}
	om.Init()

	cm := &CoderMarket{}
	cm.Init()

	totalMoney := 100.0 //单位：w
	var oldCoder Coder
	var dsNoOrd int

//	for n := 0;n < 10;n ++ {
//		go func(){
			for i := 1;;i++{
				if totalMoney > 1000.0 || totalMoney < 0 {
					fmt.Printf("创业结束. totalMoney:%v", totalMoney)
					return
				}

				fmt.Printf("第%d轮工作开始",i)

				fmt.Println("totalMoney=",totalMoney)

				order, err := om.GetOrder()
				if err != nil {
					dsNoOrd ++
					fmt.Println("没有拉到订单，休息一天后继续")
					time.Sleep(time.Second * 1) //1秒模拟1天
					totalMoney -= 0.1
					continue
				}
				dsNoOrd = 0
				if dsNoOrd > 60 {
					fmt.Printf("连续 %d 天没拿到订单,公司破产",dsNoOrd)
					if oldCoder != nil {
						//遣散费
						fmt.Println("程序员拿到遣散费:",oldCoder.GetSalary())
						totalMoney -= oldCoder.GetSalary()
					}
					fmt.Printf("公司破产,totalMoney:%v",totalMoney)
					return
				}

				var coder Coder
				var timeGetCoder float64
				var moneyToCoder float64
				if oldCoder == nil {
					coder, err = cm.GetCoder()
					if err != nil {
						fmt.Println("没有找到程序员，退回订单后继续")
						om.ReturnOrder(order)
						totalMoney -= order.Money * 2
						continue
					}
					coder.SetSalary(float64(rand.Intn(3)))
					timeGetCoder = float64(rand.Intn(10)) * 0.1 //month
				}else {
					coder = oldCoder
				}



				software, err := coder.WriteCode(order,timeGetCoder)
				if err != nil {
					fmt.Println("开发失败，按订单价格2倍赔偿")

					moneyToCoder := coder.GetSalary()
					coder.GetMoney(moneyToCoder)

					totalMoney -= moneyToCoder + order.Money * 2
					goto flag
				}

				om.DeliverOrder(order, software)
				totalMoney += order.Money

				moneyToCoder = order.Money / 10.0  + coder.GetSalary()
				coder.GetMoney(moneyToCoder)
				totalMoney -= moneyToCoder

				flag:
				if rand.Intn(3) % 3 == 0 {
					oldCoder = coder
					oldCoder.SetSalary(coder.GetSalary() + 0.1 )
					fmt.Println("程序员留在公司,涨薪 0.1W")
				}else {
					oldCoder = nil
					fmt.Println("程序员离开公司")
				}

				fmt.Println("订单间歇，休息两天\n")
				totalMoney -= 0.1 * 2
				time.Sleep(time.Second * 2) //1秒模拟1天
			}
//		}()
//		time.Sleep(time.Second * 3)
//	}

}

//------------------------------------order

type Order struct {
	Money float64 //w
	Name  string
	DeadTime int	//month
}

type OrderMarket struct {
	orderPool []Order
}

func (om *OrderMarket) Init() {

	rand.Seed(time.Now().Unix())
	var orderPool []Order
	for i := 0; i < 10000; i++ {
		money := float64(rand.Int()%50 + 50)
		order := Order{
			Money: money,
			Name:  "order:" + strconv.Itoa(i),
			DeadTime:rand.Intn(int(money)),
		}
		orderPool = append(orderPool, order)
	}

	om.orderPool = orderPool
}

func (om *OrderMarket) GetOrder() (order Order, err error) {
	if rand.Intn(9) % 9 != 0 {
		//拿不到订单
		return order,errors.New("没拿到订单")
	}
	index := rand.Int() % len(om.orderPool)
	order = om.orderPool[index]
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


//---------------------------------

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

func (cm *CoderMarket) GetCoder() (coder Coder, err error) {

	if rand.Intn(3) % 3 == 0 {
		return coder,errors.New("没找到程序员")
	}
	index := rand.Int() % len(cm.coderPool)
	coder = cm.coderPool[index]
	fmt.Println("succeed to get a coder")
	return
}


type SoftWare struct {
	Name string
}


//---------------------------coder

type Coder interface {
	WriteCode(order Order,timeGetCoder float64) (SoftWare, error)
	GetMoney(money float64)
	SetSalary(salary float64)
	GetSalary() float64
}

type GoCoder struct {
	Name string
	Salary float64
}

func (gc *GoCoder) WriteCode(order Order,timeGetCoder float64) (sf SoftWare, err error) {

	fmt.Printf("begin to write code for %v\n", order.Name)

	time.Sleep(time.Second * 2)
	timeUsed := timeGetCoder + float64(rand.Intn(order.DeadTime + 1))

	if timeUsed > float64(order.DeadTime) {
		err = errors.New("time out")
		return
	}

	if rand.Intn(5) % 5 ==0 {
		err = errors.New("failed to writeCode")
	}

	fmt.Println("write code ok.")

	sf.Name = fmt.Sprintf("hello, world! by %v", gc.Name)
	return
}

func (gc *GoCoder) GetMoney(money float64) {
	fmt.Printf("%v get money:%vw\n", gc.Name, money)
}

func (gc *GoCoder) SetSalary(salary float64) {
	gc.Salary = salary
}

func (gc *GoCoder) GetSalary() float64 {
	return gc.Salary
}

//------------------------doctor

type Doctor interface {
	Cure() (err error)
	GetMoney(money float64)
}

type Dentist struct {
	Name string
	Salary float64
}

func (d *Dentist) Cure() (err error) {
	//拔牙
	return
}

func (d *Dentist) GetMoney(money float64) {
	fmt.Printf("%v get money:%vw\n", d.Name, money)
}

func (d *Dentist) WriteCode(order Order,timeGetCoder float64) (sf SoftWare, err error) {

	fmt.Printf("i'm a dentist. begin to write code for %v\n", order.Name)
	time.Sleep(time.Second * 2)

	timeUsed := timeGetCoder + float64(rand.Intn(order.DeadTime + 1))

	if timeUsed > float64(order.DeadTime) {
		err = errors.New("time out")
	}

	if rand.Intn(5) % 5 ==0 {
		err = errors.New("failed to writeCode")
	}

	fmt.Println("write code ok.")

	sf.Name = fmt.Sprintf("hello, world! by dentist:%v", d.Name)
	return
}

func (d *Dentist) SetSalary (salary float64) {
	d.Salary = salary
}

func (d *Dentist) GetSalary() float64 {
	return d.Salary
}