package priorityqueue

import (
	"reflect"
	"testing"

	"github.com/gregoryalbouy/go-datastructures/testx"
)

func TestPQEnqueue(t *testing.T) {
	pq := New()
	pq.Enqueue("3", 10).Enqueue("0", -20).Enqueue("2", 10).Enqueue("1", 8)
	exp := []interface{}{"0", "1", "2", "3"}
	got := SliceOf(pq)

	if !reflect.DeepEqual(exp, got) {
		t.Errorf("Invalid priority queue, expected %v, got %v", exp, got)
	}
}

func TestPQDequeue(t *testing.T) {
	testcases := []testx.Testcase{
		{
			Desc:     "empty",
			Input:    New(),
			Expected: []interface{}{nil, nil, nil, nil},
		}, {
			Desc:     "regular value",
			Input:    New().Enqueue("0", 100),
			Expected: []interface{}{"0", nil, nil, nil},
		}, {
			Desc:     "negative priorities",
			Input:    New().Enqueue("1", -1).Enqueue("0", -2).Enqueue("2", 1),
			Expected: []interface{}{"0", "1", "2", nil},
		},
	}

	for _, tc := range testcases {
		q := tc.Input.(Interface)
		gotPeek := []interface{}{}
		gotDequeue := []interface{}{}

		for i := 0; i < 4; i++ {
			n0 := q.Len()
			gotPeek = append(gotPeek, q.Peek())
			gotDequeue = append(gotDequeue, q.Dequeue())
			n1 := q.Len()

			if !(n1 == n0-1 || n0 == 0) {
				t.Errorf(".Len(): expected %v, got %v", n0-1, n1)
			}
		}

		testx.Check(t, tc, gotPeek)
		testx.Check(t, tc, gotDequeue)
	}
}
