### Description

Implement a Queue. A queue is a linear structure which follows a particular order in which the operations are performed. The order is First In First Out (FIFO)

### Example:

```
Input: Enqueue "value1", Enqueue "value2", Enqueue "value3", Dequeue 
Output: ["value2", "value3"]
```

### Run test
```sh
# run local dir mode without any package (need to go to the targeted dir)
go test -run "TestFIFOList" -v --bench "BenchmarkFIFOList" --benchmem

# or run all 
go test -v --bench . --benchmem
```