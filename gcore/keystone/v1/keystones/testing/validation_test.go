package testing

import (
	"testing"

	gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"
	"github.com/alexandrpaliy/gcorelabscloud-go/gcore/keystone/v1/types"

	"github.com/alexandrpaliy/gcorelabscloud-go/gcore/keystone/v1/keystones"
	"github.com/stretchr/testify/require"
)

func TestUpdateOptsValidation(t *testing.T) {
	opts := keystones.UpdateOpts{}
	err := gcorecloud.TranslateValidationError(opts.Validate())
	require.Error(t, err)
	opts = keystones.UpdateOpts{
		State: types.KeystoneStateDeleted,
	}
	err = gcorecloud.TranslateValidationError(opts.Validate())
	require.NoError(t, err)
}
