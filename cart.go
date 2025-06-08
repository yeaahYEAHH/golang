package main

type Cart []Product

func (c *Cart) findProductNyName(name string) (Product, int) {
	for i, product := range *c {
		if product.GetName() == name {
			return product, i
		}
	}

	return nil, -1
}

func (c *Cart) AddProduct(p Product) {
	*c = append(*c, p)
}

func (c *Cart) RemoveProduct(name string) {
	_, index := c.findProductNyName(name)

	if index == -1 {
		return
	}

	*c = append((*c)[:index], (*c)[index+1:]...)
}

func (c Cart) TotalPrice() (total float64) {
	for i := range c {
		total += c[i].GetPrice()
	}
	return
}

func (c Cart) TotalDiscountPrice(discount float64) (total float64) {
	for _, product := range c {
		product.SetPrice(product.GetDiscountPrice(discount))
		total += product.GetPrice()
	}
	return
}

func (c *Cart) Clear() {
	c = nil
}
