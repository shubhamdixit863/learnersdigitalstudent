package operations

const size int = 5

var Mp = make(map[int]string)
var slc []int

func Put(key int, value string) {
	_, ok := Mp[key]
	if ok {
		Mp[key] = value
		moveToRecent(key)
		return
	}

	if len(slc) >= size {
		lru := slc[0]
		delete(Mp, lru)
		slc = slc[1:]
	}

	Mp[key] = value
	slc = append(slc, key)
}

func Get(key int) string {
	val, ok := Mp[key]
	if !ok {
		return "-1"
	}
	moveToRecent(key)
	return val
}
func moveToRecent(key int) {
	for i := 0; i < len(slc); i++ {
		if slc[i] == key {
			slc = append(slc[:i], slc[i+1:]...)
			break
		}
	}
	slc = append(slc, key)
}
