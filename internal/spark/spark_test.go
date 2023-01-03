package spark

import "testing"

func TestStatus_String(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want string
	}{
		// TODO: Add test cases.
		{
			"succeess get string value of Succeeded status",
			Succeeded,
			"succeeded",
		},
		{
			"succeess get string value of Running status",
			Running,
			"running",
		},
		{
			"succeess get string value of Failed status",
			Failed,
			"failed",
		},
		{
			"succeess get string value of Unknown status",
			Unknown,
			"unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Status.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_EnumIndex(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want int
	}{
		// TODO: Add test cases.
		{
			"succeess get EnumIndex of Running status",
			Running,
			1,
		},
		{
			"succeess get EnumIndex of Succeeded status",
			Succeeded,
			2,
		},
		{
			"succeess get EnumIndex of Failed status",
			Failed,
			3,
		},
		{
			"succeess get EnumIndex of Unknown status",
			Unknown,
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.EnumIndex(); got != tt.want {
				t.Errorf("Status.EnumIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSparkStatus(t *testing.T) {
	type args struct {
		status string
	}
	tests := []struct {
		name    string
		args    args
		want    Status
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"succeess get EnumIndex of Running status",
			args{
				status: "running",
			},
			Running,
			false,
		},
		{
			"succeess get EnumIndex of Succeeded status",
			args{
				status: "succeeded",
			},
			Succeeded,
			false,
		},
		{
			"succeess get EnumIndex of Failed status",
			args{
				status: "failed",
			},
			Failed,
			false,
		},
		{
			"succeess get EnumIndex of Running status",
			args{
				status: "running",
			},
			Running,
			false,
		},
		{
			"succeess get EnumIndex of Unknown status",
			args{
				status: "unknown",
			},
			Unknown,
			false,
		},
		{
			"error get EnumIndex",
			args{
				status: "test",
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSparkStatus(tt.args.status)
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
