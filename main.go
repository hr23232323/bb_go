// A simple project which demonstrates concurrency
// and other distributed system features using
// simple examples
// Built by Harsh Rana using Go

package main

import (
  "fmt";
  "flag";
  "math/rand"
)

func main(){
  // Parse input flag to get which program to run
  exec := flag.Int("e", 1, "Which of two programs to run (concurrency/load balancer)")
  flag.Parse()
  //fmt.Println(*exec)

  // Run program based on flag
  switch *exec{
  case 1:
    runConcurrency(1000)
  case 2:
    runLoadBalancer()
  }
}

// Function to showcase concurrency
// All helper functions can be found in conc.go
func runConcurrency(n int){
  fmt.Println("Running concurrency program")

  // Array of size n to sum up
  arr := rand.Perm(n)
  //fmt.Println(arr)

  // Sum using linear sum
  linearSum,linearTime := concBenchmark(linearSumFunc, arr)
  fmt.Println("\nUsing linear, non-concurrent system-")
  fmt.Println("Sum: ", linearSum)
  fmt.Println("Time taken: ", linearTime)

  // Sum using concurrent program
  concSum,concTime := concBenchmark(concSumFunc, arr)
  fmt.Println("\nUsing Concurrency (2 parallel systems)-")
  fmt.Println("Sum: ", concSum)
  fmt.Println("Time taken: ", concTime)
}


// Function to showcase load balancer
func runLoadBalancer(){
  fmt.Println("Running load balancer program")

}
