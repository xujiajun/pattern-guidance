package proxy

import "testing"

func TestProxyBuyer_Buy(t *testing.T) {
	huangniu := Huangniu{}
	proxyBuyer := ProxyBuyer{Buyer: huangniu}
	proxyBuyer.Buy()//buy a ticket
}
