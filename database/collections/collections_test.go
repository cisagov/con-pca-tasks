package collections

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetCycleTargets(t *testing.T) {
	t.Parallel()
	test_cycle_id := ""
	var want []Target
	got, err := GetCycleTargets(test_cycle_id)
	if err != nil {
		t.Fatal(err)
	}
	sort.Slice(got, func(i, j int) bool {
		return got[i].LastName < got[j].LastName
	})
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
