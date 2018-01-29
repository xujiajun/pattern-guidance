package proxy

import "fmt"

type Huangniu struct {
}

type ProxyBuyer struct {
	Buyer Huangniu
}

type ProxyBuy interface {
	Buy()
}

func (hn *Huangniu) Buy() {
	fmt.Print("buy a ticket")
}

func (this *ProxyBuyer) Buy() {
	this.Buyer.Buy()
}
