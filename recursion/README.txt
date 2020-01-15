# Purpose

want to compare the performance difference between recursive function call vs loop

# Result

```
Running tool: /usr/local/bin/go test -benchmem -run=^$ github.com/bananaumai/golang-suburi/recursion -bench .

goos: darwin
goarch: amd64
pkg: github.com/bananaumai/golang-suburi/recursion
BenchmarkAddOne05-8         	100000000	        11.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOne10-8         	49529954	        23.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOne15-8         	36443932	        33.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOne20-8         	26475997	        46.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOne25-8         	19672849	        62.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOne30-8         	16826917	        71.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOne35-8         	14811507	        83.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOne40-8         	12815822	        98.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOne45-8         	10645878	       108 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOneLoop05-8     	399350289	         3.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOneLoop10-8     	269373379	         4.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOneLoop15-8     	202453089	         5.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOneLoop20-8     	169631545	         7.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOneLoop25-8     	145909214	         8.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOneLoop30-8     	100000000	        10.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOneLoop35-8     	100000000	        10.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOneLoop40-8     	50129823	        23.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddOneLoop45-8     	91104475	        13.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib05-8            	73202766	        15.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib10-8            	 6904626	       175 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib15-8            	  565112	      1943 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib20-8            	   56378	     22595 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib25-8            	    4630	    246414 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib30-8            	     430	   2703146 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib35-8            	      40	  29954181 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib40-8            	       3	 355101870 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib45-8            	       1	3729820669 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibLoop05-8        	362657804	         3.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibLoop10-8        	184445899	         6.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibLoop15-8        	142342096	         8.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibLoop20-8        	95058772	        11.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibLoop25-8        	92379769	        13.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibLoop30-8        	76265991	        16.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibLoop35-8        	65523163	        19.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibLoop40-8        	54171258	        21.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibLoop45-8        	52074846	        23.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsense05-8       	125444278	         9.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsense10-8       	67205775	        18.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsense15-8       	44349536	        27.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsense20-8       	33906141	        37.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsense25-8       	20221328	        54.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsense30-8       	20227946	        62.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsense35-8       	16453954	        68.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsense40-8       	15788329	        77.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsense45-8       	13805998	        88.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsenseLoop05-8   	346632619	         3.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsenseLoop10-8   	255789908	         4.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsenseLoop15-8   	199439128	         5.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsenseLoop20-8   	156588922	         7.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsenseLoop25-8   	136452508	         8.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsenseLoop30-8   	100000000	        10.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsenseLoop35-8   	99220285	        13.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsenseLoop40-8   	48578025	        23.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNonsenseLoop45-8   	89365604	        14.2 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/bananaumai/golang-suburi/recursion	79.174s
Success: Benchmarks passed.
```
