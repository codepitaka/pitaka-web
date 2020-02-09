package config

import (
	"errors"
	"os"
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_DatabaseURL(t *testing.T){
	// data for test
	tests := []struct {
		APPMODE  string
		expectedDatabaseURL string
	}{
		{"PRD", "https://pitaka-server.herokuapp.com"},
		{"DEV", "https://pitaka-server-dev.herokuapp.com"},
		{"RVW", "https://pitaka-server-dev.herokuapp.com"},
		{"LOCAL", "https://port-3000-pitaka.run.goorm.io"},
	}
	
	currentAPPMODE := os.Getenv("APPMODE")
	
	for _, test := range tests {
		t.Run(test.APPMODE, func(t *testing.T) {
			os.Setenv("APPMODE", test.APPMODE)
			config := New()
			actualDatabaseURL := config.DatabaseURL()
			assert.Equal(t, test.expectedDatabaseURL, actualDatabaseURL)
		})
	}
	os.Setenv("APPMODE", currentAPPMODE)
}

func Test_getENV(t *testing.T) {
	// data for test
	tests := []struct {
		OSAppMode  string
		expectedAppMode string
	}{
		{"PRD", "PRD"},
		{"DEV", "DEV"},
		{"RVW", "RVW"},
		{"LOCAL", "LOCAL"},
	}
	
	currentAPPMODE := os.Getenv("APPMODE")
	
	for _, test := range tests {
		t.Run(test.OSAppMode, func(t *testing.T) {
			os.Setenv("APPMODE", test.OSAppMode)
			actualAppMode := getENV("APPMODE")
			// Assert we entered title correctly
			assert.Equal(t, test.expectedAppMode, actualAppMode)
		})	
	}
	
	os.Setenv("APPMODE", currentAPPMODE)
}

func Test_initializeDatabaseURL(t *testing.T){
	// data for test
	tests := []struct {
		APPMODE  string
		expectedDatabaseURL string
		expectedError error
	}{
		{"PRD", "https://pitaka-server.herokuapp.com", nil},
		{"DEV", "https://pitaka-server-dev.herokuapp.com", nil},
		{"RVW", "https://pitaka-server-dev.herokuapp.com", nil}, // same dev server
		{"LOCAL", "https://port-3000-pitaka.run.goorm.io", nil},
		{"WRONG APPMODE", "INVALID", errors.New("Invalid AppMode.")},
	}
	
	for _, test := range tests {
		t.Run(test.APPMODE, func(t *testing.T) {
			actualDatabaseURL, e := initializeDatabaseURL(test.APPMODE)
			assert.Equal(t, test.expectedError, e)
			assert.Equal(t, test.expectedDatabaseURL, actualDatabaseURL)
		})
	}
}