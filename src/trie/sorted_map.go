package trie

import "fmt"

type Key interface {
	Greater(v interface{}) bool
	Equal(v interface{}) bool
}

type SortedMap struct {
	index []Key
	value map[Key]interface{}
}

// Print debug
func (sm *SortedMap) Print() {
	fmt.Println("-----start-----")
	for k, v := sm.Front(); k != nil && v != nil; k, v = sm.Next(k) {
		fmt.Println("key: ", k, ", value:", v)
	}
	fmt.Println("-----end-----")
}

// PrintBack debug
func (sm *SortedMap) PrintBack() {
	fmt.Println("-----start-----")
	for k, v := sm.Back(); k != nil && v != nil; k, v = sm.Previous(k) {
		fmt.Println("key: ", k, ", value:", v)
	}
	fmt.Println("-----end-----")
}

func NewMap() *SortedMap {
	return new(SortedMap).Init()
}

func (sm *SortedMap) Init() *SortedMap {
	sm.index = []Key{}
	sm.value = make(map[Key]interface{})
	return sm
}

func (sm *SortedMap) Len() int {
	return len(sm.index)
}

// binarySearch true, index; false, last index that less than value
func (sm *SortedMap) binarySearch(value Key) (bool, int) {
	if sm.Len() == 0 {
		return false, 0
	}
	l := 0
	r := sm.Len() - 1
	mid := 0
	for l <= r {
		mid = (l + r) / 2
		if sm.index[mid].Greater(value) { // more
			r = mid - 1
		} else if !sm.index[mid].Equal(value) { // less
			l = mid + 1
		} else { // equal
			break
		}
	}
	if !sm.index[mid].Equal(value) {
		if sm.index[mid].Greater(value) {
			return false, mid
		}
		return false, mid + 1
	}
	return true, mid
}

// Insert true if succeeded; false if its exists
func (sm *SortedMap) Insert(key Key, value interface{}) bool {
	found, index := sm.binarySearch(key)
	if !found {
		sm.index = append(sm.index[:index], append([]Key{key}, sm.index[index:]...)...)
	}
	sm.value[key] = value

	return !found
}

// Remove true if succeeded; false if not found
func (sm *SortedMap) Remove(key Key) bool {
	found, index := sm.binarySearch(key)
	if !found {
		return false
	}
	delete(sm.value, sm.index[index])
	sm.index = append(sm.index[:index], sm.index[index + 1:]...)

	return true
}

// Get value, true if succeeded; nil, false if not found
func (sm *SortedMap) Get(key Key) (interface{}, bool) {
	if value, found := sm.value[key]; found {
		return value, found
	}
	return nil, false
}

// Front key, value if succeeded; nil, nil if its empty
func (sm *SortedMap) Front() (Key, interface{}) {
	if sm.Len() == 0 {
		return nil, nil
	}
	return sm.index[0], sm.value[sm.index[0]]
}

// Back key, value if succeeded; nil, nil if its empty
func (sm *SortedMap) Back() (Key, interface{}) {
	if sm.Len() == 0 {
		return nil, nil
	}
	return sm.index[sm.Len() - 1], sm.value[sm.index[sm.Len() - 1]]
}

// Next key value if succeeded; nil, nil if not found or reach max length
func (sm *SortedMap) Next(key Key) (Key, interface{}) {
	found, index := sm.binarySearch(key)
	if !found || index >= sm.Len() - 1 {
		return nil, nil
	}
	return sm.index[index + 1], sm.value[sm.index[index + 1]]
}

// Previous key value if succeeded; nil, nil if not found or reach minimum length
func (sm *SortedMap) Previous(key Key) (Key, interface{}) {
	found, index := sm.binarySearch(key)
	if !found || index <= 0 {
		return nil, nil
	}
	return sm.index[index - 1], sm.value[sm.index[index - 1]]
}
