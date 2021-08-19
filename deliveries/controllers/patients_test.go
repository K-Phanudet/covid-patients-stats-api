package controllers

import (
	"reflect"
	"testing"

	"github.com/K-Phanudet/covid-patients-stats-api/models"
)

func TestGroupByProvince(t *testing.T) {
	type args struct {
		in0 []models.PatientInfo
	}
	tests := []struct {
		name string
		args args
		want map[string][]models.PatientInfo
	}{
		{
			name: "Two province",
			args: args{
				in0: []models.PatientInfo{
					{
						Age:      23,
						Gender:   "Male",
						Province: "Bankkok",
					},
					{
						Age:      18,
						Gender:   "Male",
						Province: "Chaing Mai",
					},
					{
						Age:      19,
						Gender:   "Male",
						Province: "Bankkok",
					},
					{
						Age:      20,
						Gender:   "Male",
						Province: "Bankkok",
					},
				},
			},
			want: map[string][]models.PatientInfo{
				"Bankkok": {
					{
						Age:      23,
						Gender:   "Male",
						Province: "Bankkok",
					},
					{
						Age:      19,
						Gender:   "Male",
						Province: "Bankkok",
					},
					{
						Age:      20,
						Gender:   "Male",
						Province: "Bankkok",
					},
				},
				"Chaing Mai": {
					{
						Age:      18,
						Gender:   "Male",
						Province: "Chaing Mai",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupByProvince(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByProvince() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupByAge(t *testing.T) {
	type args struct {
		patients []models.PatientInfo
	}
	tests := []struct {
		name string
		args args
		want map[string][]models.PatientInfo
	}{
		{
			name: "Must PASS",
			args: args{
				patients: []models.PatientInfo{
					{
						Age:      31,
						Gender:   "Male",
						Province: "Bankkok",
					},
					{
						Age:      30,
						Gender:   "Male",
						Province: "Chaing Mai",
					},
					{
						Age:      65,
						Gender:   "Male",
						Province: "Bankkok",
					},
					{
						Age:      20,
						Gender:   "Male",
						Province: "Bankkok",
					},
					{
						Age:      0,
						Gender:   "Male",
						Province: "Bankkok",
					},
					{
						Age:      60,
						Gender:   "Male",
						Province: "Bankkok",
					},
				},
			},
			want: map[string][]models.PatientInfo{
				"61+": {
					{
						Age:      65,
						Gender:   "Male",
						Province: "Bankkok",
					},
				},
				"31-60": {
					{
						Age:      31,
						Gender:   "Male",
						Province: "Bankkok",
					},
					{
						Age:      60,
						Gender:   "Male",
						Province: "Bankkok",
					},
				},
				"0-30": {
					{
						Age:      30,
						Gender:   "Male",
						Province: "Chaing Mai",
					},
					{
						Age:      20,
						Gender:   "Male",
						Province: "Bankkok",
					},
					{
						Age:      0,
						Gender:   "Male",
						Province: "Bankkok",
					},
				},
				"N/A": {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupByAge(tt.args.patients); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountValueInMap(t *testing.T) {
	type args struct {
		in0 map[string][]models.PatientInfo
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "",
			args: args{
				in0: map[string][]models.PatientInfo{
					"61+": {
						{
							Age:      65,
							Gender:   "Male",
							Province: "Bankkok",
						},
					},
					"31-60": {
						{
							Age:      31,
							Gender:   "Male",
							Province: "Bankkok",
						},
						{
							Age:      60,
							Gender:   "Male",
							Province: "Bankkok",
						},
					},
					"0-30": {
						{
							Age:      30,
							Gender:   "Male",
							Province: "Chaing Mai",
						},
						{
							Age:      20,
							Gender:   "Male",
							Province: "Bankkok",
						},
						{
							Age:      0,
							Gender:   "Male",
							Province: "Bankkok",
						},
					},
					"N/A": {},
				},
			},
			want: map[string]int{
				"31-60": 2,
				"0-30":  3,
				"61+":   1,
				"N/A":   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countValueInMap(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountValueInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
