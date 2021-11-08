package check

import (
	"testing"
)

func TestCheck(t *testing.T) {
	var chkr checker
	var testCases = []struct {
		url         string
		excpChunkNB int
		max         int
		excpErr     error
	}{
		{
			url:         "https://xxxxx.cloudfront.net/gebrish_user_XXXXXX_XXXXXXX/chunked/0.ts",
			excpChunkNB: 14,
			max:         100,
			excpErr:     nil,
		},
		{
			url:         "https://xxxxx.cloudfront.net/gebrish_user_XXXXXX_XXXXXXX/chunked/111111.ts",
			excpChunkNB: 14,
			max:         100,
			excpErr:     nil,
		},
		{
			url:         "https://xxxxx.cloudfront.net/gebrish_user_XXXXXX_XXXXXXX/chunked/0.ts",
			excpChunkNB: 14,
			max:         -1,
			excpErr:     nil,
		},
		{
			url:         "https://xxxxx.cloudfront.net/gebrish_user_XXXXXX_XXXXXXX/chunked/0.ts",
			excpChunkNB: -1,
			max:         5,
			excpErr:     ErrOverMaxInt,
		},
		{
			url:         "https://xxxxx.cloudfront.net/gebrish_user_XXXXXX_XXXXXXX/0.ts",
			excpChunkNB: -1,
			max:         100,
			excpErr:     ErrInvalid,
		},
	}

	for _, tc := range testCases {
		chkr = NewTestCloudfrontChecker(tc.url, tc.max)
		lastChunkNb, err := chkr.Check()
		if tc.excpErr != err {
			t.Fatalf("Expected %s got %s instead.\n", tc.excpErr, err)
		}
		if lastChunkNb != tc.excpChunkNB {
			t.Fatalf("Expected %d , got %d\n", tc.excpChunkNB, lastChunkNb)
		}
	}
}
