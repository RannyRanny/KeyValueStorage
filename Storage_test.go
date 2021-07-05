package KeyValueStorage

import (
	"testing"
	"time"
)

func TestStorage_AddValue(t *testing.T) {
	var storage = NewStorage()
	var result = storage.AddValue(1, 2)
	if result != 2 {
		t.Error("Result should be 2, but got ", result)
	}

}
func TestWhenAddAndRead_ThenGetsValue(t *testing.T) {
	var storage = NewStorage()
	storage.AddValue("123", 1)
	var result = storage.ReadValue("123")
	if result != 1 {
		t.Error("Result should be 1, but got ", result)
	}
}

func TestWhenAddAndUpdateAndRead_ThenGetsUpdatedValue(t *testing.T) {
	var storage = NewStorage()
	storage.AddValue("123", 1)
	storage.UpdateValue("123", 4)
	var result = storage.ReadValue("123")
	if result != 4 {
		t.Error("Result should be 4, but got ", result)
	}
}

func TestWhenAddAndAddExistingKeyAndRead_ThenGetsInitialValue(t *testing.T) {
	var storage = NewStorage()
	storage.AddValue("123", 1)
	storage.AddValue("123", 4)
	var result = storage.ReadValue("123")
	if result != 1 {
		t.Error("Result should be 4, but got ", result)
	}
}

func TestWhenAddAnDeleteAndRead_ThenGetsNil(t *testing.T) {
	var storage = NewStorage()
	storage.AddValue("123", 1)
	storage.DeleteValue("123")
	var result = storage.ReadValue("123")
	if result != nil {
		t.Error("Result should be <nil>, but got ", result)
	}
}
func TestWhenAddAsync_ThenAllSaved(t *testing.T) {
	var storage = NewStorage()
	for i := 0; i < 10; i++ {
		go storage.AddValue(i, i)
	}
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {

		var result = storage.ReadValue(i)
		if result != i {
			t.Errorf("Result should be %v, but got %v", i, result)
		}
	}
}
