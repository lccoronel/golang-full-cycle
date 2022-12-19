package carro

type Carro struct {
	Marca string
}

func (c Carro) MostraMarca() string {
	return c.Marca
}
