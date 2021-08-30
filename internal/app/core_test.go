package app

import (
	"reflect"
	"testing"

	"github.com/go-chi/chi"
)

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string
		want *chi.Mux
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
