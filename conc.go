// Concurrency module in ButcherBox TA Supplement program
// Methods used to showcase the difference between
// Run a program linearly V/S using concurrency

package main

import (
  "time";
  "fmt";
  "math/rand"
)


// Function to showcase concurrency
// All helper functions can be found in conc.go
func runConcurrency(n int){
  fmt.Println("Running concurrency program")

  // Array of size n to sum up
  arr := rand.Perm(n)

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

// Function to run a passed input function and return
// result and time taken
func concBenchmark(f func([]int) int, arr []int) (int, time.Duration){
  start := time.Now()
  sum := f(arr)
  timeTaken := time.Since(start)

  return sum, timeTaken
}

// Function to emulate time taken for 1 operation
func runOper(){
  time.Sleep(1 * time.Millisecond)
}


// Function which uses simple linear traversal to find sum
// Also uses runOper method to mimic time taken for operation
func linearSumFunc(arr []int) int {
  linearSum := 0
  for _,v := range arr {
    linearSum += v
    runOper()
  }
  return linearSum
}


// Function to sum up numbers in array and
// Write to channel
func cSum(arr []int, c chan int){
  c <- linearSumFunc(arr)
}


// Function which uses concurrency to find sum
// Also uses runOper method to mimic time taken for operation
func concSumFunc(arr []int) int {
  // Create channels
  c1 := make(chan int)
  c2 := make(chan int)

  // Start goroutines to run parallelly
  go cSum(arr[:len(arr)/2], c1)
  go cSum(arr[len(arr)/2:], c2)

  // Get summed up answers
  concSum := <-c1 + <-c2
  return concSum
}
