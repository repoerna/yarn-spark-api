package yarn

import (
	"reflect"
	"testing"
)

func TestGetApplicationListQueryParams_GetApplicationList(t *testing.T) {
	type fields struct {
		User              string
		FinalStatus       string
		Queue             string
		Limit             string
		StartedTimeBegin  string
		FinishedTimeBegin string
		FinishedTimeEnd   string
		Name              string
		DeSelect          string
		States            string
		ApplicationTypes  string
		ApplicationTags   string
	}
	type args struct {
		server string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *YarnApplicationList
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"success get yarn app list",
			fields{
				User:             "ampid",
				States:           "RUNNING",
				ApplicationTypes: "SPARK",
			},
			args{
				"http://bdgbnbcldnn01.intra.excelcom.co.id:8088",
			},
			&YarnApplicationList{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GetApplicationListQueryParams{
				User:              tt.fields.User,
				FinalStatus:       tt.fields.FinalStatus,
				Queue:             tt.fields.Queue,
				Limit:             tt.fields.Limit,
				StartedTimeBegin:  tt.fields.StartedTimeBegin,
				FinishedTimeBegin: tt.fields.FinishedTimeBegin,
				FinishedTimeEnd:   tt.fields.FinishedTimeEnd,
				Name:              tt.fields.Name,
				DeSelect:          tt.fields.DeSelect,
				States:            tt.fields.States,
				ApplicationTypes:  tt.fields.ApplicationTypes,
				ApplicationTags:   tt.fields.ApplicationTags,
			}
			got, err := g.GetApplicationList(tt.args.server)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApplicationListQueryParams.GetApplicationList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApplicationListQueryParams.GetApplicationList() = %v, want %v", got, tt.want)
			}
		})
	}
}
