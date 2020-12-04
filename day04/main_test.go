package main

import "testing"

func Test_validateHeight(t *testing.T) {
	tests := []struct {
		name       string
		passport   map[string]string
		wantAnswer bool
	}{
		{"1", map[string]string{"hgt": "60in"}, true},
		{"2", map[string]string{"hgt": "190cm"}, true},
		{"3", map[string]string{"hgt": "190in"}, false},
		{"4", map[string]string{"hgt": "190"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := validateHeight(tt.passport); gotAnswer != tt.wantAnswer {
				t.Errorf("validateHeight() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}

func Test_validateByr(t *testing.T) {
	tests := []struct {
		name       string
		passport   map[string]string
		wantAnswer bool
	}{
		{"1", map[string]string{"byr": "2002"}, true},
		{"2", map[string]string{"byr": "2003"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := validateByr(tt.passport); gotAnswer != tt.wantAnswer {
				t.Errorf("validateByr() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}

func Test_validateIyr(t *testing.T) {
	tests := []struct {
		name       string
		passport   map[string]string
		wantAnswer bool
	}{
		{"1", map[string]string{"iyr": "2010"}, true},
		{"2", map[string]string{"iyr": "2020"}, true},
		{"3", map[string]string{"iyr": "2022"}, false},
		{"4", map[string]string{"iyr": "2009"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := validateIyr(tt.passport); gotAnswer != tt.wantAnswer {
				t.Errorf("validateIyr() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}

func Test_validateEyr(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name     string
		passport map[string]string

		wantAnswer bool
	}{
		{"1", map[string]string{"eyr": "2020"}, true},
		{"2", map[string]string{"eyr": "2030"}, true},
		{"3", map[string]string{"eyr": "2031"}, false},
		{"4", map[string]string{"eyr": "2019"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := validateEyr(tt.passport); gotAnswer != tt.wantAnswer {
				t.Errorf("validateEyr() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}

func Test_validateHairColor(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name       string
		passport   map[string]string
		wantAnswer bool
	}{
		{"1", map[string]string{"hcl": "#123abc"}, true},
		{"2", map[string]string{"hcl": "#123abz"}, false},
		{"3", map[string]string{"hcl": "123abc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := validateHairColor(tt.passport); gotAnswer != tt.wantAnswer {
				t.Errorf("validateHairColor() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}

func Test_validateEyeColor(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name     string
		passport map[string]string

		wantAnswer bool
	}{
		{"1", map[string]string{"ecl": "brn"}, true},
		{"2", map[string]string{"ecl": "wat"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := validateEyeColor(tt.passport); gotAnswer != tt.wantAnswer {
				t.Errorf("validateEyeColor() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}

func Test_validatePID(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name       string
		passport   map[string]string
		wantAnswer bool
	}{
		{"1", map[string]string{"pid": "000000001"}, true},
		{"2", map[string]string{"pid": "0123456789"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := validatePID(tt.passport); gotAnswer != tt.wantAnswer {
				t.Errorf("validatePID() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
