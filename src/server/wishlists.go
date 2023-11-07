package server

type Wishlist struct {
	wishlist [8][]string
}
type Data struct {
	User   []string
	Target []string
}

var mapping = map[uint8]uint8{0: 0, 1: 2, 2: 3, 3: 1, 4: 7, 5: 6, 6: 4, 7: 5}

func (w *Wishlist) SaveWishlist(id uint8, value []string) {
	if id > 7 {
		return
	}
	w.wishlist[id] = value
}

func (w *Wishlist) LoadWishlists(id uint8) Data {
	if id > 7 {
		return Data{}
	}
	return Data{w.wishlist[id], w.wishlist[findTarget(id)]}
}

func findTarget(id uint8) uint8 {
	return mapping[id]
}
