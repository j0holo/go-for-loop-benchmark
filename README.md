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

## Running the benchmarks

go test -run='^$' -bench=. -benchmem -count=10 ./oop/ > oop.txt
go test -run='^$' -bench=. -benchmem -count=10 ./aos/ > aos.txt
go test -run='^$' -bench=. -benchmem -count=10 ./swa/ > swa.txt

## Results

See benchmark_data.ods or benchmark_data.xslx to view the data in a spreadsheet for a clear view of each benchmark.

```text
goos: linux
goarch: amd64
pkg: inventory_benchmark
cpu: AMD Ryzen 7 5800X3D 8-Core Processor           
                                              │     oop.txt     │                aos.txt                 │               swa.txt                │
                                              │     sec/op      │     sec/op       vs base               │    sec/op     vs base                │
Allocation-16                                       76.73µ ± 4%       78.60µ ± 7%       ~ (p=0.796 n=10)   44.40µ ± 26%  -42.14% (p=0.000 n=10)
TotalWeightOfIndividualItemsRangeLoopValue-16   11106.000n ± 1%   11236.000n ± 2%       ~ (p=0.063 n=10)   7.395n ±  1%  -99.93% (p=0.000 n=10)
TotalWeightOfIndividualItemsRangeLoopIndex-16       4.079n ± 2%       4.085n ± 2%       ~ (p=0.644 n=10)   4.079n ±  1%        ~ (p=0.896 n=10)
TotalWeightOfIndividualItems-16                     4.048n ± 1%       4.059n ± 2%       ~ (p=0.811 n=10)   4.072n ±  1%   +0.61% (p=0.030 n=10)
TotalWeightOfInventory-16                           7.352n ± 1%       7.330n ± 1%       ~ (p=0.280 n=10)   4.912n ±  1%  -33.19% (p=0.000 n=10)
geomean                                             159.6n            160.7n       +0.74%                  30.59n        -80.83%

                                              │    oop.txt     │                aos.txt                │                  swa.txt                  │
                                              │      B/op      │     B/op      vs base                 │     B/op      vs base                     │
Allocation-16                                   336.0Ki ± 0%     336.0Ki ± 0%       ~ (p=1.000 n=10)     336.0Ki ± 0%         ~ (p=1.000 n=10)
TotalWeightOfIndividualItemsRangeLoopValue-16     3.000 ± 0%       3.000 ± 0%       ~ (p=1.000 n=10) ¹     0.000 ± 0%  -100.00% (p=0.000 n=10)
TotalWeightOfIndividualItemsRangeLoopIndex-16     0.000 ± 0%       0.000 ± 0%       ~ (p=1.000 n=10) ¹     0.000 ± 0%         ~ (p=1.000 n=10) ¹
TotalWeightOfIndividualItems-16                   0.000 ± 0%       0.000 ± 0%       ~ (p=1.000 n=10) ¹     0.000 ± 0%         ~ (p=1.000 n=10) ¹
TotalWeightOfInventory-16                         0.000 ± 0%       0.000 ± 0%       ~ (p=1.000 n=10) ¹     0.000 ± 0%         ~ (p=1.000 n=10) ¹
geomean                                                      ²                 +0.00%                ²                 ?                       ² ³
¹ all samples are equal
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean

                                              │   oop.txt    │               aos.txt               │               swa.txt               │
                                              │  allocs/op   │ allocs/op   vs base                 │ allocs/op   vs base                 │
Allocation-16                                   1.000 ± 0%     1.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%       ~ (p=1.000 n=10) ¹
TotalWeightOfIndividualItemsRangeLoopValue-16   0.000 ± 0%     0.000 ± 0%       ~ (p=1.000 n=10) ¹   0.000 ± 0%       ~ (p=1.000 n=10) ¹
TotalWeightOfIndividualItemsRangeLoopIndex-16   0.000 ± 0%     0.000 ± 0%       ~ (p=1.000 n=10) ¹   0.000 ± 0%       ~ (p=1.000 n=10) ¹
TotalWeightOfIndividualItems-16                 0.000 ± 0%     0.000 ± 0%       ~ (p=1.000 n=10) ¹   0.000 ± 0%       ~ (p=1.000 n=10) ¹
TotalWeightOfInventory-16                       0.000 ± 0%     0.000 ± 0%       ~ (p=1.000 n=10) ¹   0.000 ± 0%       ~ (p=1.000 n=10) ¹
geomean                                                    ²               +0.00%                ²               +0.00%                ²
¹ all samples are equal
² summaries must be >0 to compute geomean
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