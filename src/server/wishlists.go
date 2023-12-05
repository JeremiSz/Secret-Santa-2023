package server

import (
	"log"
	"os"
	"strings"
	"sync"
)

const (
	fILE = "saved.txt"
)

type Wishlist struct {
	lock     sync.RWMutex
	wishlist [PEOPLE_COUNT]string
}
type Data struct {
	User       string
	Target     []string
	TargetName string
	UserId     uint8
}

var (
	names = [PEOPLE_COUNT]string{"Admin", "Alex", "Shay", "Rowan", "Remy", "Kit", "Azriel", "Sage"}
)

func (w *Wishlist) SaveWishlist(id uint8, value string) {
	if id > 7 {
		return
	}
	w.lock.Lock()
	defer w.lock.Unlock()
	w.loadFromFile()
	w.wishlist[id] = value
	log.Println("Saved wishlist for id: ", id, "to", value)

}

func (w *Wishlist) LoadWishlists(id uint8) Data {
	if id > 7 {
		return Data{}
	}
	w.lock.RLock()
	defer w.lock.RUnlock()

	targetId := findTarget(id)

	list := w.wishlist[targetId]
	items := strings.Split(list, "\n")
	w.saveToFile()
	return Data{w.wishlist[id], items, names[targetId], id}
}

func findTarget(id uint8) uint8 {
	if id == 0 {
		return 0
	}
	id = id - 1
	id = id + 2
	id = id % (PEOPLE_COUNT - 1)
	return id + 1
}

func (w *Wishlist) saveToFile() {
	f, err := os.OpenFile(fILE, os.O_WRONLY, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	for a := range w.wishlist {
		f.WriteString(w.wishlist[a])
		f.WriteString(":")
	}
	f.Sync()
}
func (w *Wishlist) loadFromFile() {
	f, err := os.OpenFile(fILE, os.O_RDONLY, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	b := make([]byte, 1024)
	f.Read(b)
	s := string(b)
	values := strings.Split(s, ":")
	for i := 0; i < PEOPLE_COUNT; i++ {
		w.wishlist[i] = values[i]
	}
}
