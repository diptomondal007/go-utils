package variables

import (
	"reflect"
	"testing"
)

func TestBoolP(t *testing.T) {
	var a = true
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want *bool
	}{
		// TODO: Add test cases.
		{
			name: "t-01",
			args: args{v: false},
			want: new(bool),
		},
		{
			name: "t-02",
			args: args{v: true},
			want: &a,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolP(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32P(t *testing.T) {
	var a = [2]float32{1.00, 2.00}
	type args struct {
		v float32
	}
	tests := []struct {
		name string
		args args
		want *float32
	}{
		// TODO: Add test cases.
		{
			name: "t-01",
			args: args{v: a[0]},
			want: &a[0],
		},
		{
			name: "t-02",
			args: args{v: a[1]},
			want: &a[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32P(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32P() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64P(t *testing.T) {
	var a = [2]float64{1.00, 2.00}
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want *float64
	}{
		// TODO: Add test cases.
		{
			name: "t-01",
			args: args{v: a[0]},
			want: &a[0],
		},
		{
			name: "t-02",
			args: args{v: a[1]},
			want: &a[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64P(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64P() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32P(t *testing.T) {
	var a = [2]int32{1, 2}
	type args struct {
		v int32
	}
	tests := []struct {
		name string
		args args
		want *int32
	}{
		// TODO: Add test cases.
		{
			name: "t-01",
			args: args{v: a[0]},
			want: &a[0],
		},
		{
			name: "t-02",
			args: args{v: a[1]},
			want: &a[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32P(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32P() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64P(t *testing.T) {
	var a = [2]int64{1, 2}
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want *int64
	}{
		// TODO: Add test cases.
		{
			name: "t-01",
			args: args{v: a[0]},
			want: &a[0],
		},
		{
			name: "t-02",
			args: args{v: a[1]},
			want: &a[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64P(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64P() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntP(t *testing.T) {
	var a = [2]int{1, 2}
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		// TODO: Add test cases.
		{
			name: "t-01",
			args: args{v: a[0]},
			want: &a[0],
		},
		{
			name: "t-02",
			args: args{v: a[1]},
			want: &a[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntP(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t-01",
			args: args{a: 1, b: 2},
			want: 2,
		},
		{
			name: "t-01",
			args: args{a: 10, b: 15},
			want: 15,
		},
		{
			name: "t-01",
			args: args{a: 20, b: 1},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t-01",
			args: args{a: 1, b: 2},
			want: 1,
		},
		{
			name: "t-01",
			args: args{a: 10, b: 15},
			want: 10,
		},
		{
			name: "t-01",
			args: args{a: 20, b: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringP(t *testing.T) {
	var test = [2]string{"test", "test1"}
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		// TODO: Add test cases.
		{
			name: "t-01",
			args: args{v: ""},
			want: new(string),
		},
		{
			name: "t-02",
			args: args{v: test[0]},
			want: &test[0],
		},
		{
			name: "t-03",
			args: args{v: test[1]},
			want: &test[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringP(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{
			name: "t-01",
			args: args{value: "10"},
			want: 10,
		},
		{
			name: "t-02",
			args: args{value: "true"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt32(tt.args.value); got != tt.want {
				t.Errorf("ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}
