package main

type Smartphone struct {
	name    string
	price   float64
	Brand   string
	Storage uint
	RAM     uint
}

func (s *Smartphone) GetName() string {
	return s.name
}

func (s *Smartphone) GetPrice() float64 {
	return s.price
}

func (s *Smartphone) SetPrice(price float64) {
	s.price = price
}

func (s *Smartphone) GetDiscountPrice(discount float64) float64 {
	if discount == 100 {
		return -1
	}
	return s.price * (100 - discount) / 100
}

type TV struct {
	name       string
	price      float64
	Brand      string
	ScreenSize float32
	Resolution string
}

func (t *TV) GetName() string {
	return t.name
}

func (t *TV) GetPrice() float64 {
	return t.price
}

func (t *TV) SetPrice(price float64) {
	t.price = price
}

func (t *TV) GetDiscountPrice(discount float64) float64 {
	if discount == 100 {
		return -1
	}
	return t.price * (100 - discount) / 100
}

type Laptop struct {
	name        string
	price       float64
	Brand       string
	CPU         string
	RAM         uint
	BatteryLife uint
}

func (l *Laptop) GetName() string {
	return l.name
}

func (l *Laptop) GetPrice() float64 {
	return l.price
}

func (l *Laptop) SetPrice(price float64) {
	l.price = price
}

func (l *Laptop) GetDiscountPrice(discount float64) float64 {
	if discount == 100 {
		return -1
	}
	return l.price * (100 - discount) / 100
}

type Product interface {
	GetPrice() float64
	GetName() string
	GetDiscountPrice(discount float64) float64
	SetPrice(price float64)
}
