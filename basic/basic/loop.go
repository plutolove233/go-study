package main

func main(){
	done := make(chan int, 20)
	for v := range 10 {
		go func() {
			// println(v)
			done<-v
		}()
	}
	for {
		select {
		case data:=<-done:
			println(data)
			break
		}
	}
}