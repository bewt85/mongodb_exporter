package collector

import(
    "testing"
	"github.com/dcu/mongodb_exporter/shared"
)

func Test_IndexCountersCollectData(t *testing.T) {
	stats := &IndexCounterStats{
	}

	groupName := "index_counters"
	stats.Export(groupName)

	if shared.Groups[groupName] == nil {
		t.Error("Group not created")
	}
}

