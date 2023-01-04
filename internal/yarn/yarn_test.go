package yarn

import "testing"

func TestStates_String(t *testing.T) {
	tests := []struct {
		name string
		s    States
		want string
	}{
		// TODO: Add test cases.
		{
			"succeess get string value of ALL status",
			ALL,
			"ALL",
		},
		{
			"succeess get string value of NEW status",
			NEW,
			"NEW",
		},
		{
			"succeess get string value of NEW_SAVING status",
			NEW_SAVING,
			"NEW_SAVING",
		},
		{
			"succeess get string value of SUBMITTED status",
			SUBMITTED,
			"SUBMITTED",
		},
		{
			"succeess get string value of ACCEPTED status",
			ACCEPTED,
			"ACCEPTED",
		},
		{
			"succeess get string value of RUNNING status",
			RUNNING,
			"RUNNING",
		},
		{
			"succeess get string value of FINISHED status",
			FINISHED,
			"FINISHED",
		},
		{
			"succeess get string value of FAILED status",
			FAILED,
			"FAILED",
		},
		{
			"succeess get string value of KILLED status",
			KILLED,
			"KILLED",
		},
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
		{
			"succeess get EnumIndex of ALL status",
			ALL,
			1,
		},
		{
			"succeess get EnumIndex of NEW status",
			NEW,
			2,
		},
		{
			"succeess get EnumIndex of NEW_SAVING status",
			NEW_SAVING,
			3,
		},
		{
			"succeess get EnumIndex of SUBMITTED status",
			SUBMITTED,
			4,
		},
		{
			"succeess get EnumIndex of ACCEPTED status",
			ACCEPTED,
			5,
		},
		{
			"succeess get EnumIndex of RUNNING status",
			RUNNING,
			6,
		},
		{
			"succeess get EnumIndex of FINISHED status",
			FINISHED,
			7,
		},
		{
			"succeess get EnumIndex of FAILED status",
			FAILED,
			8,
		},
		{
			"succeess get EnumIndex of KILLED status",
			KILLED,
			9,
		},
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
		{
			"succeess get EnumIndex of ALL status",
			args{
				states: "ALL",
			},
			ALL,
			false,
		},
		{
			"succeess get EnumIndex of NEW status",
			args{
				states: "NEW",
			},
			NEW,
			false,
		},
		{
			"succeess get EnumIndex of NEW_SAVING status",
			args{
				states: "NEW_SAVING",
			},
			NEW_SAVING,
			false,
		},
		{
			"succeess get EnumIndex of SUBMITTED status",
			args{
				states: "SUBMITTED",
			},
			SUBMITTED,
			false,
		},
		{
			"succeess get EnumIndex of ACCEPTED status",
			args{
				states: "ACCEPTED",
			},
			ACCEPTED,
			false,
		},
		{
			"succeess get EnumIndex of RUNNING status",
			args{
				states: "RUNNING",
			},
			RUNNING,
			false,
		},
		{
			"succeess get EnumIndex of FINISHED status",
			args{
				states: "FINISHED",
			},
			FINISHED,
			false,
		},
		{
			"succeess get EnumIndex of FAILED status",
			args{
				states: "FAILED",
			},
			FAILED,
			false,
		},
		{
			"succeess get EnumIndex of KILLED status",
			args{
				states: "KILLED",
			},
			KILLED,
			false,
		},
		{
			"error get EnumIndex",
			args{
				states: "test",
			},
			0,
			true,
		},
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
