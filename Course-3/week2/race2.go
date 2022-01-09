/*
Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.
Submission: Upload your source code for the program along with your written explanation of race conditions.
*/

package main
import "fmt"
import "sync"
import "time"

var x int 
var wg sync.WaitGroup

func producer(){
   for i:=0; i<10000; i++ {
      x=i
   }
   wg.Done()
}

func consumer(){
	if x!=-1 {
		fmt.Printf("x now is : %d  (CTRL-C to quit)\n",x)
	}
	wg.Done()
}



func main(){
	for {
		x=-1
		wg.Add(1)
		go producer()
		wg.Add(1)
		go consumer()   //This goroutine call demonstrate the race condition: value of x is printed whenever we had a value different from "-1"
						//demonstrating that the producer goroutine gets interrupted in a non-deterministic moment.
						//Thus, we can not determine value of x, its value at this stage is random.
						
		wg.Wait()							//these are to not write resource killing program.
		t,_:=time.ParseDuration("10ms")		//these are to not write resource killing program.
		time.Sleep(t) 
	}
}

