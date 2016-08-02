package bitmap_test

import (
	"testing"

	"github.com/caoyongzheng/data-structures/bitmap"
)

func TestBitmapSort(t *testing.T) {
	bmap := bitmap.NewBitmapSize(7)

	origal := [5]uint64{4, 6, 3, 1, 7}
	expected := [5]uint64{1, 3, 4, 6, 7}
	actual := [5]uint64{}

	for _, offset := range origal {
		bmap.SetBit(offset, 1)
	}

	var i uint64
	var offset, maxpos uint64 = 0, bmap.Maxpos() + 1
	for ; offset < maxpos; offset++ {
		if bmap.GetBit(offset) == 1 {
			actual[i] = offset
			i++
		}
	}

	if expected != actual {
		t.Errorf("expected:%#v, actual:%#v", expected, actual)
	}
}
