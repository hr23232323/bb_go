// Load Balancer module in ButcherBox TA Supplement program
// Methods used to showcase a simple load balancing algorithm
// Using goroutines and channels


package main

import (
  "time";
  "math/rand";
  "fmt"
)

// Main function to run the load balancing simulation
func runLoadBalancer(nWorkers int, nReq int){
  fmt.Println("Running load balancer program")

  // Create channels for creating
  // Requests and Responses
  req := make(chan Work)
  res := make(chan Work)

  // Generate workers
  for i := 0; i < nWorkers; i++{
    go worker(req, res, i)
  }

  // Generate 100 requests
  go generateReq(req, nReq)

  // Process requests
  processRes(res, nReq)
}


// Function to generate requests
func generateReq (req chan Work, n int){
  // create n requests
  i := 0
  for i < n {
    // Create a new request after waiting for some time
    time.Sleep(time.Duration(1 * time.Millisecond))

    // Add new request to channel
    req <- Work{rand.Intn(1000), 0, 0}
  }

  // No more requests
  close(req)
}


// Struct to represent some work done.
// Input, output and worker ID present
type Work struct{
  in int
  out int
  id int
}

// function used to "work". Simple task is performed
// and some time delay is added. Also worker ID is kept in work
func worker (req chan Work, res chan Work, i int){
  for w := range req {
    w.out = w.in * 10
    w.id = i
    runOper()
    res <- w
  }
}

// Function to process responses being recieved from workers
func processRes(res chan Work, n int){
  // Process all responses
  for i:=0; i < n; i++ {
    select {
    case x, ok := <-res:
          if ok {
              fmt.Printf("Response Received. \nInput %d was converted to output %d by Worker #%d\n\n", x.in, x.out, x.id)
          } else {
              fmt.Println("Channel closed!")
          }
      default:
          // Need to wait for response
          time.Sleep(1)
      }
  }
}
