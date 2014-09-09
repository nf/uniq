
# Uniq number brenchmark in Golang

Implementation and benchmarks of the two methods of generating unique numbers
described in this blog post:[Concurrency patterns: a source of unique numbers](http://nf.id.au/posts/2010/08/concurrency-patterns-a-source-of-unique-numbe.html)

## Result of 2014-09-09 (go1.3.1)
First the numbers:

	$ go test -bench . -cpu 1,2,4,8
	
	BenchmarkMutex     50000000        50.3 ns/op
	BenchmarkMutex-2   50000000        50.8 ns/op
	BenchmarkMutex-4   50000000        52.5 ns/op
	BenchmarkMutex-8   50000000        51.9 ns/op
	BenchmarkChannel   20000000         130 ns/op
	BenchmarkChannel-2  5000000         500 ns/op
	BenchmarkChannel-4  5000000         435 ns/op
	BenchmarkChannel-8  5000000         521 ns/op

The mutex-based approach is always faster and performs constantly even if
the number of used cores is increased.

The channel-based approach takes 2.58x the time on a single core, and
suffers even more in a multi-core environment.

## Result of 2010-08-29 (pre go-r56)
The output of `gomake bench`:

	uniq.BenchmarkMutex    10000000  173 ns/op
	uniq.BenchmarkChannel  10000000  224 ns/op

The channel-based approach takes 1.29x the time of the mutex-based approach.

Depending on the specific application, and your opinion on the aesthetics
of the two approaches, this can be an acceptable tradeoff (it is, IMO).

Andrew Gerrand <nf@wh3rd.net>

