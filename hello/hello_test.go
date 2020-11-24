package main

import "testing"

func Test_makeUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test1",
			want: "123-456-789",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeUUID(); got != tt.want {
				t.Errorf("makeUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
