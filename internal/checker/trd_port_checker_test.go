package checker

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_trdChecker_AssertRunning(t1 *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	t1.Run("trd check on success", func(t *testing.T) {
		trdChecker, _ := NewTRDChecker()
		httpmock.RegisterResponder("GET", trdUrl,
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewJsonResponse(200, map[string]interface{}{})
			},
		)
		err := trdChecker.AssertRunning()
		assert.Nil(t1, err)
	})

	t1.Run("trd check on failure", func(t *testing.T) {
		trdChecker, _ := NewTRDChecker()
		httpmock.RegisterResponder("GET", trdUrl,
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(500, ""), nil
			},
		)
		err := trdChecker.AssertRunning()
		assert.NotNil(t1, err)
	})
}
