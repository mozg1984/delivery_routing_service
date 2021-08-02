package geo

import (
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	type args struct {
		lat1 float64
		lon1 float64
		lat2 float64
		lon2 float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Kazan-Ulyanovsk", args: args{lat1: 55.788455, lon1: 49.138789, lat2: 54.330445, lon2: 48.353620}, want: 169656.32410680264},
		{name: "Voronez-StaryOskal", args: args{lat1: 51.672555, lon1: 39.217412, lat2: 51.251172, lon2: 37.866136}, want: 104683.31838375759},
		{name: "House-School", args: args{lat1: 54.281414, lon1: 48.262210, lat2: 54.277618, lon2: 48.268620}, want: 592.7312698892108},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDistance(tt.args.lat1, tt.args.lon1, tt.args.lat2, tt.args.lon2); got != tt.want {
				t.Errorf("CalculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
