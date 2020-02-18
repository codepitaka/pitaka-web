package utils

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_HTMLTemplatePathsUnder(t *testing.T) {
	_ = os.Mkdir("./testDir", os.FileMode(0777))
	_, _ = os.Create("./testDir/testFile1.tmpl")
	_, _ = os.Create("./testDir/testFile2.html")
	_, _ = os.Create("./testDir/testFile3.go")
	_, _ = os.Create("./testDir/testFile4.js")
	defer os.RemoveAll("./testDir")
	
	var filePaths []string = HTMLTemplatePathsUnder("./testDir")
	var expectedPaths []string = []string{"testDir/testFile1.tmpl", "testDir/testFile2.html"} 
	
	assert.Equal(t, expectedPaths, filePaths)
}