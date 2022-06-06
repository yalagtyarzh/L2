package clock

import "testing"

func Test_getDate(t *testing.T) {
	theTests := []struct {
		name    string
		isError bool
		url     string
	}{
		{
			name:    "valid",
			isError: false,
			url:     "0.ru.pool.ntp.org",
		},
		{
			name:    "timeout",
			isError: true,
			url:     "pepega",
		},
	}

	for _, tt := range theTests {
		t.Run(
			tt.name, func(t *testing.T) {
				_, err := GetDate(tt.url)

				if !(isErrorExist(err) == tt.isError) {
					t.Errorf("Expected error existing: %t, got %s", tt.isError, err)
				}
			},
		)
	}
}

func isErrorExist(err error) bool {
	if err == nil {
		return false
	}

	return true
}
