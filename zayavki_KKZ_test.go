//Задача "Заявки в киноконцертный зал".

//Сколько макс заявок мы можем расположить в рассписании зала.

package main

import (
	"reflect"
	"testing"
)

func Test_maxZayavok(t *testing.T) {
	type args struct {
		zal     beginend
		actions []beginend
	}

	tests := []struct {
		name    string
		args    args
		wantRet []beginend
	}{
		{"1",
			args{
				NewBeginend("09:00AM", "09:00PM"),
				[]beginend{
					NewBeginend("08:20AM", "09:30AM"),
					NewBeginend("09:00AM", "6:00PM"),
					NewBeginend("11:00AM", "12:00PM"),
					NewBeginend("09:10AM", "10:00AM"),
				},
			},

			[]beginend{
				NewBeginend("09:10AM", "10:00AM"),
				NewBeginend("11:00AM", "12:00PM"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := maxZayavok(tt.args.zal, tt.args.actions); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("maxZayavok() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
