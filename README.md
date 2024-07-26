# Benchmarking loops in various programming styles

This is an experiment to see how much impact different programming styles have and by coincidence discovering
how slow range loop values are when using large structs.

DISCLAIMER: I know the OOP style and the Array of Structs is almost identical.

There are three kind of styles being benchmarked:

- Object-Oriented Programming (OOP): how I was taught
- Array of Structs (AOS): what I normally would do in Go
- Structs with Arrays (SWA): which tries to do [data-oriented design](https://en.wikipedia.org/wiki/Data-oriented_design) which is new to me. 

There are three kind of benchmarks:

- Allocation
- A function which sums the weight of all item in an inventory
  - range with value
  - range with index
  - plain for loop
- A function which sums the weight * amount of all items in an inventory

## Results

See benchmark_data.ods or benchmark_data.xslx to view the data in a spreadsheet for a clear view of each benchmark.

```text
goos: linux
goarch: amd64
pkg: inventory_benchmark
cpu: AMD Ryzen 7 5800X3D 8-Core Processor
BenchmarkOOPAllocation-16                                                          19975             67023 ns/op          344090 B/op          2 allocs/op
BenchmarkAOSAllocation-16                                                          15957             77640 ns/op          344064 B/op          1 allocs/op
BenchmarkSWAAllocation-16                                                          23386             47790 ns/op          344064 B/op          1 allocs/op
BenchmarkOOPInventory_TotalWeightOfIndividualItems_RangeLoopValue-16              220039              5529 ns/op               1 B/op          0 allocs/op
BenchmarkOOPInventory_TotalWeightOfIndividualItems_RangeLoopIndex-16            270968862                4.320 ns/op           0 B/op          0 allocs/op
BenchmarkOOPInventory_TotalWeightOfIndividualItems-16                           160244653                7.340 ns/op           0 B/op          0 allocs/op
BenchmarkAOS_TotalWeightOfIndividualItems_RangeLoopValue-16                        94953             11146 ns/op               3 B/op          0 allocs/op
BenchmarkAOS_TotalWeightOfIndividualItems_RangeLoopIndex-16                     281356263                4.138 ns/op           0 B/op          0 allocs/op
BenchmarkAOS_TotalWeightOfIndividualItems-16                                    282187761                4.069 ns/op           0 B/op          0 allocs/op
BenchmarkSWA_TotalWeightOfIndividualItems_RangeLoopValue-16                     131188309                9.107 ns/op           0 B/op          0 allocs/op
BenchmarkSWA_TotalWeightOfIndividualItems_RangeLoopIndex-16                     161561829                7.336 ns/op           0 B/op          0 allocs/op
BenchmarkSWA_TotalWeightOfIndividualItems-16                                    160562037                7.508 ns/op           0 B/op          0 allocs/op
BenchmarkOOPInventory_TotalWeightOfInventory-16                                 155808315                7.701 ns/op           0 B/op          0 allocs/op
BenchmarkOOPInventory_TotalWeightOfInventory_ExtraCall-16                       154883451                7.763 ns/op           0 B/op          0 allocs/op
BenchmarkAOS_TotalWeightOfInventory-16                                          158777956                7.521 ns/op           0 B/op          0 allocs/op
BenchmarkSWA_TotalWeightOfInventory-16                                          157290012                7.722 ns/op           0 B/op          0 allocs/op
PASS
ok      inventory_benchmark     29.595s
```

## Conclusion

Using the for range iterator can seriously cripple tight loops because each value is a copy and not a reference to the 
actual value. For the Struct with Arrays this causes no real issue (a 2ns difference) but for OOP or Array of Structs 
causes a slowdown of ~1200x or ~2600x for OOP and AOS compared to using a range with index loop.

I'm a bit clueless why the Data-oriented Programming is actually slower than the other styles in the single attribute
benchmark. The weight array should be the most cache-friendly of all implementations.

The disassembled code `go tool compile -S inventory.go > inventory.s` has been included if somebody understands why
the compiler and/or CPU is faster in OOP/AOS style. Maybe because the caches of my CPU are too large to notice
the benefits of Data-Oriented programming. Another option is that the access pattern is still predictable so the 
memory prefetch hides the issue that OOP and AOS style does not fit into the L1 cache.

OOP and AOS have a 22566 byte struct, which total to 22566 * 15 = 337566 bytes.
SWA has 120 bytes for the int arrays so needs 120 or 240 bytes for the loop tests.

337566 / 1024 ~= 330KiB which could fit in the L2 cache which is 512 KiB on and AMD Ryzen 5800x3D.

So that could be it. That the amount of data is not large enough to show the advantage of programming in a 
cache-friendlier way.