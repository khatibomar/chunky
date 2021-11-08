package check

import (
	"testing"
)

func TestCheck(t *testing.T) {
	var chkr checker
	excpected := 15

	chkr = NewTestCloudfrontChecker("https://xxxxx.cloudfront.net/gebrish_user_XXXXXX_XXXXXXX/chunked/0.ts", 200)
	nb, err := chkr.Check()
	if err != nil {
		t.Fatal(err)
	}
	if nb != excpected {
		t.Fatalf("Expected %d , got %d\n", excpected, nb)
	}
}
