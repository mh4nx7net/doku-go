package main
import "sync"
func main() {


	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		println(`coba`)
	}()
	go func() {
		println(`hello`)
		wg.Done()
	}()

	go func() {
		println(`world`)
		wg.Done()
	}()
	wg.Wait()
}

