package singleton

func main(){
	for i:=0;i<100;i++{
		instance := getSingleInstance()
	}
	
	// do something with the instance
}