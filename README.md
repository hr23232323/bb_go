# ButcherBox Technical Assessment Supplement

Two simple programs built using Go to showcase the basics of distributed system features; in specific distributed computing and load balancing.  


### Usage
The project consists of two parts: concurrency simulation and load balancer simulation. 

To run the concurrency code, use the following command - `go run main.go conc.go lb.go -e=1`

To run the load balancer code, use the following command - `go run main.go conc.go lb.go -e=2`

### More information

#### Problem 1
Given a fixed array of numbers, how fast can we sum them up? A very simple answer would be to traverse the array and sum up numbers as you encounter them. This would be fairly efficient, O(n), but can be improved even further by using goroutines and channels to compute sum of sub arrays in parallel. This same strategy can be used to optimize any computation that can occur independently on separate parts of an array. This is the basis of MapReduce and other distributed computing divide and conquer strategies.  


#### Problem 2
Given a stream of incoming requests and a distributed system of n servers, how do we distribute loads? A basic approach would be to send requests to a single server until it is full and proceed to the next and so on. But this is inefficient because it leaves resources underutilized and will also increase wait times for responses. Another method could be to iterate through all servers and send one request to each server and repeat. This again wouldn’t scale properly as it doesn’t take into account response times and individual server loads. Here, a simple round robin algorithm can be used to distribute requests efficiently. This is the basis of Load Balancers. 
