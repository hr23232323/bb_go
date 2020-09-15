// A simple project which demonstrates concurrency
// and other distributed system features using
// simple examples
// Built by Harsh Rana using Go

package main

import (
  "flag";
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
    runLoadBalancer(5,100)
  }
}
