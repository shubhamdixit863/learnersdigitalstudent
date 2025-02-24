package passenger

func Passenger(a[][]string)(string,int){
  mp:=make(map[string]int)
	for _,num:=range a{
		name:=num[0]
		mp[name]++
	}
	b:=0
  c:=""
	for key,val:=range mp{

		if(val>b){
			b=val
      c=key
		}
	}
	return c,b
}