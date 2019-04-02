package metrics

import (
	"time"
	"runtime"
	"runtime/debug"
	"testing"
	"fmt"
	gometrics "github.com/rcrowley/go-metrics"
	"github.com/alipay/sofa-mosn/pkg/metrics/shm"
)

func gcPause() time.Duration {
	runtime.GC()
	var stats debug.GCStats
	debug.ReadGCStats(&stats)
	return stats.PauseTotal
}

const (
	entries = 200000
)

func TestGCPause(t *testing.T) {
	debug.SetGCPercent(10)
	fmt.Println("Number of entries: ", entries)



	r := gometrics.NewRegistry()
	for i := 0; i < entries; i++ {
		key := fmt.Sprintf("key-%010d", i)
		r.GetOrRegister(key, gometrics.NewGauge).(gometrics.Gauge).Update(int64(i))
	}
	key := fmt.Sprintf("key-%010d", 0)
	if r.Get(key).(gometrics.Gauge).Value() != 0 {
		t.Error("gometrics entry value not correct")
	}
	fmt.Println("GC pause for gometrics: ", gcPause())

	ResetAll()
	gcPause()

	//------------------------------------------

	func() {
		zone := shm.InitMetricsZone("testGCPause", 20000000*256)
		defer zone.Detach()

		m, _ := NewMetrics("test", map[string]string{"type": "gc"})
		for i := 0; i < entries; i++ {
			key := fmt.Sprintf("key-%010d", i)
			m.Gauge(key).Update(int64(i))
		}

		key := fmt.Sprintf("key-%010d", 0)
		if m.Gauge(key).Value() != 0 {
			t.Error("shm-based entry value not correct")
		}

		fmt.Println("GC pause for shm-based metrics: ", gcPause())
	}()



}
