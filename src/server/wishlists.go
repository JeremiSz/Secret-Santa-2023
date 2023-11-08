package server

import "sync"

type Wishlist struct {
	lock     sync.RWMutex
	wishlist [8][]string
}
type Data struct {
	User       []string
	Target     []string
	TargetName string
}

var (
	mapping = map[uint8]uint8{0: 0, 1: 2, 2: 3, 3: 1, 4: 7, 5: 6, 6: 4, 7: 5}
	names   = []string{"Admin", "Alicja", "Jeremi", "Pjotrek", "Gosia", "Pawel", "Gosia", "Fletcher"}
)

func (w *Wishlist) SaveWishlist(id uint8, value []string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	if id > 7 {
		return
	}
	w.wishlist[id] = value

}

func (w *Wishlist) LoadWishlists(id uint8) Data {
	w.lock.RLock()
	defer w.lock.RUnlock()
	if id > 7 {
		return Data{}
	}
	targetId := findTarget(id)
	return Data{w.wishlist[id], w.wishlist[targetId], names[targetId]}
}

func findTarget(id uint8) uint8 {
	return mapping[id]
}
