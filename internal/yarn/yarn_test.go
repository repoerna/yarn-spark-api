package yarn

import "testing"

func TestStates_String(t *testing.T) {
	tests := []struct {
		name string
		s    States
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("States.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStates_EnumIndex(t *testing.T) {
	tests := []struct {
		name string
		s    States
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.EnumIndex(); got != tt.want {
				t.Errorf("States.EnumIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSparkStatus(t *testing.T) {
	type args struct {
		states string
	}
	tests := []struct {
		name    string
		args    args
		want    States
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSparkStatus(tt.args.states)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSparkStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseSparkStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
