package main

func main() {
	ch := make(chan<- int, 2) // This is BUFFERING

	ch <- 21
	ch <- 9

	/*
				"fmt.Println(<-ch)"...This will not work because the is send (write to) channel and not a recieve (read) channel.
				 It is denoted with the sign "<-" when creating the channel.

				*example:
		        ch := make(chan <- int, n) ---------- This is a send only channel

				 ch := make(chan int, n)

				Values can only sent to this channel. Hence, no need fot the fmt package

	*/
}


//This Code here, will not work. It is an example of a recieve only channel

// func buff() {
// 	ch := make(<- chan int, 2) // This is BUFFERING

// 	ch <- 21
// 	ch <- 9
// }

