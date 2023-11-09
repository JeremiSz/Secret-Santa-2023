package server

import (
	"sync"
	"testing"
)

func TestLoad(t *testing.T) {
	wishlist := Wishlist{sync.RWMutex{}, [8]string{"Best\ntest"}}
	data := wishlist.LoadWishlists(0)
	if data.TargetName != "Admin" {
		t.Error("Expected Admin, got ", data.TargetName)
	}
	if data.UserId != 0 {
		t.Error("Expected 0, got ", data.UserId)
	}
	if data.Target[0] != "Best" {
		t.Error("Expected Best, got ", data.Target[0])
	}
}

func TestSave(t *testing.T) {
	wishlist := Wishlist{sync.RWMutex{}, [8]string{"Best\ntest"}}
	wishlist.SaveWishlist(0, "test")
	data := wishlist.LoadWishlists(0)
	if data.TargetName != "Admin" {
		t.Error("Expected Admin, got ", data.TargetName)
	}
	if data.UserId != 0 {
		t.Error("Expected 0, got ", data.UserId)
	}
	if data.Target[0] != "test" {
		t.Error("Expected test, got ", data.Target[0])
	}
}

func TestLoadInvalid(t *testing.T) {
	wishlist := Wishlist{sync.RWMutex{}, [8]string{"Best\ntest"}}
	data := wishlist.LoadWishlists(8)
	if data.TargetName != "" {
		t.Error("Expected empty, got ", data.TargetName)
	}

}

func TestSaveInvalid(t *testing.T) {
	wishlist := Wishlist{sync.RWMutex{}, [8]string{"Best\ntest"}}
	wishlist.SaveWishlist(8, "test")
	data := wishlist.LoadWishlists(8)
	if data.TargetName != "" {
		t.Error("Expected empty, got ", data.TargetName)
	}
}

func TestFindTarget(t *testing.T) {
	if findTarget(0) != 0 {
		t.Error("Expected 0, got ", findTarget(0))
	}
	if findTarget(1) != 2 {
		t.Error("Expected 2, got ", findTarget(1))
	}
	if findTarget(2) != 3 {
		t.Error("Expected 3, got ", findTarget(2))
	}
	if findTarget(3) != 1 {
		t.Error("Expected 1, got ", findTarget(3))
	}
	if findTarget(4) != 7 {
		t.Error("Expected 7, got ", findTarget(4))
	}
	if findTarget(5) != 6 {
		t.Error("Expected 6, got ", findTarget(5))
	}
	if findTarget(6) != 4 {
		t.Error("Expected 4, got ", findTarget(6))
	}
	if findTarget(7) != 5 {
		t.Error("Expected 5, got ", findTarget(7))
	}
	if findTarget(8) != 0 {
		t.Error("Expected 0, got ", findTarget(8))
	}
}
