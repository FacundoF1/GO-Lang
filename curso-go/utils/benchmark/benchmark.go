package benchmark

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

var mem1_before, mem2_before, mem1_after, mem2_after runtime.MemStats

func BenchmarkTest() map[string]interface{} {

    rand.Seed(time.Now().UnixNano())

    tr := testing.Benchmark(BenchmarkNothing)
	
	out := map[string]interface{}{}
    out["Allocs/operation"] = tr.AllocsPerOp()
    out["Byte/operation"] = tr.AllocedBytesPerOp()
    out["Precise allocs/operation:"] = float64(tr.MemAllocs)/float64(tr.N)
	out["testing.Benchmark"] = tr

	return out;
}


func BenchmarkNothing(b *testing.B) {
    for i := 0; i < b.N; i++ {
        measure_nothing(&mem1_before, &mem1_after)
    }
}

func measure_nothing(before, after *runtime.MemStats) {
    runtime.GC()
    runtime.ReadMemStats(before)

    runtime.GC()
    runtime.ReadMemStats(after)
}