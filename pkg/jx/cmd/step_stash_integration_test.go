// +build integration

package cmd_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/jenkins-x/jx/pkg/gits"
	"github.com/jenkins-x/jx/pkg/helm/mocks"
	"github.com/jenkins-x/jx/pkg/jx/cmd"
	"github.com/jenkins-x/jx/pkg/tests"

	"github.com/stretchr/testify/assert"
)

func TestStepStash(t *testing.T) {
	t.Parallel()
	tempDir, err := ioutil.TempDir("", "test-step-collect")
	assert.NoError(t, err)

	testData := "test_data/step_collect/junit.xml"

	o := &cmd.StepStashOptions{
		StepOptions: cmd.StepOptions{
			CommonOptions: &cmd.CommonOptions{},
		},
	}
	o.StorageLocation.Classifier = "tests"
	o.StorageLocation.BucketURL = "file://" + tempDir
	o.ToPath = "output"
	o.Pattern = []string{testData}
	o.ProjectGitURL = "https://github.com/jenkins-x/dummy-repo.git"
	o.ProjectBranch = "master"
	cmd.ConfigureTestOptions(o.CommonOptions, &gits.GitFake{}, helm_test.NewMockHelmer())

	err = o.Run()
	assert.NoError(t, err)

	generatedFile := filepath.Join(tempDir, o.ToPath, testData)
	assert.FileExists(t, generatedFile)

	tests.AssertTextFileContentsEqual(t, testData, generatedFile)
}
