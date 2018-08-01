package orderbook

import "fmt"

type Orderbook struct {
	Ask []*Order
	Bid []*Order
	Mrt []*Order
	tr  []*Trade
}

func New() *Orderbook {
	orbook := Orderbook{}

	return &orbook
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	//tr := []*Trade{}
	//	fmt.Printf("hello")
	switch order.Kind.String() {
	case "MARKET":
		orderbook.Mrt = append(orderbook.Mrt, order)
	case "LIMIT":
		switch order.Side.String() {
		case "ASK":
			orderbook.OrderdAsk(order)
		case "BID":
			orderbook.OrderdBid(order)
		}
	}

	/*if ask.Price == bid.Price && ask.Volume == bid.Volume {
		newtr := Trade{Volume: bid.Volume, Price: ask.Price}
		tr = append(tr, &newtr)
		return tr, nil
	}

	/*if ask.Price <= bid.Price {
		//newtr := Trade{Volume: bid.Volume, Price: ask.Price}
		//tr = append(tr, &newtr)
		if (bid.Volume - ask.Volume) <= 0 {
			orderbook.Bid = append(orderbook.Bid[bid.ID:], orderbook.Bid[bid.ID+1:]...)
		}
		ask.Volume -= bid.Volume
		if ask.Volume <= 0 {
			orderbook.Ask = append(orderbook.Ask[ask.ID:], orderbook.Ask[ask.ID+1:]...)
		}
	}*/
	return orderbook.tr, nil
}

func (orderbook *Orderbook) OrderdAsk(order *Order) {
	flag := false
	for _, bid := range orderbook.Bid {
		if order.Price <= bid.Price {
			if order.Volume >= bid.Volume {
				newtr := Trade{Volume: bid.Volume, Price: order.Price}
				orderbook.tr = append(orderbook.tr, &newtr)
				flag = true
				order.Volume -= bid.Volume
			}
		}
	}
	if !flag {
		orderbook.Ask = append(orderbook.Ask, order)
		//fmt.Printf("asdasd")
	}
}

func (orderbook *Orderbook) OrderdBid(order *Order) {
	flag := false
	fmt.Println("12order - bid ", order)
	for _, ask := range orderbook.Ask {
		fmt.Println("12ask ", ask)
		if ask.Price <= order.Price {
			if ask.Volume <= order.Volume {
				newtr := Trade{Volume: ask.Volume, Price: ask.Price}
				orderbook.tr = append(orderbook.tr, &newtr)
				fmt.Println("trade was")
				flag = true
				order.Volume -= ask.Volume
				//orderbook.Ask = append(orderbook.Ask[ask.ID:], orderbook.Ask[ask.ID+1:]...)
			} else {
				newtr := Trade{Volume: order.Volume, Price: ask.Price}
				orderbook.tr = append(orderbook.tr, &newtr)
				fmt.Println("trade was is else")
				flag = true
				ask.Volume -= order.Volume
			}
		}
	}
	if !flag {
		orderbook.Bid = append(orderbook.Bid, order)
	}
}
