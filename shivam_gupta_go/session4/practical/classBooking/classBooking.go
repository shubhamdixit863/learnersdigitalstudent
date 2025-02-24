package classbooking

func Classbooking(a[][]string)map[string]int{
	mp:=make(map[string]int)

	for _,num:=range a {
		class1:=num[2]
		mp[class1]++
	}
	return mp
}