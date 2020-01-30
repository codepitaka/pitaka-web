package utils

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_FilePathsUnder(t *testing.T) {
	_ = os.Mkdir("./testDir", os.FileMode(0777))
	_, _ = os.Create("./testDir/testFile1")
	_, _ = os.Create("./testDir/testFile2")
	defer os.RemoveAll("./testDir")
	
	var filePaths []string = FilePathsUnder("./testDir")
	var expectedPaths []string = []string{"testDir/testFile1", "testDir/testFile2"} 
	
	assert.Equal(t, expectedPaths, filePaths)
}