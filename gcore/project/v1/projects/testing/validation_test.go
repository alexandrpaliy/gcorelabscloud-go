package testing

import (
	"testing"

	gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"
	"github.com/alexandrpaliy/gcorelabscloud-go/gcore/project/v1/projects"

	"github.com/stretchr/testify/require"
)

func TestUpdateOptsValidation(t *testing.T) {
	opts := projects.UpdateOpts{}
	err := gcorecloud.TranslateValidationError(opts.Validate())
	require.Error(t, err)
	opts = projects.UpdateOpts{
		Name: "test",
	}
	err = gcorecloud.TranslateValidationError(opts.Validate())
	require.NoError(t, err)
}
