package utils

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_FilesUnder(t *testing.T) {
	err := os.Mkdir("./testDir", os.FileMode(0522))
	if err != nil {
		
	}
	_, err = os.Create("./testDir/testFile1")
	if err != nil {
		
	}
	_, err = os.Create("./testDir/testFile2")
	if err != nil {
		
	}
	defer os.RemoveAll("./testDir")
	
	var filePaths []string = FilesUnder("./testDir")
	var expectedPaths []string = []string{"testDir/testFile1", "testDir/testFile2"} 
	
	assert.Equal(t, expectedPaths, filePaths)
}