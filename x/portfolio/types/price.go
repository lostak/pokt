package types

/*
	TODO:
		- Add SpotPrice CRUD
*/

func createSpotPrice(price float64) *SpotPrice {
	// TODO: add timestamp
	return &SpotPrice{
		Price: price,
	}
}
