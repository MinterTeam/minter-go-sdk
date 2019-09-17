// +build integration

package api

import (
	"testing"
)

func TestApi_Candidate(t *testing.T) {

	responseCandidates, err := testApi.Candidates(0, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(responseCandidates.Result) == 0 {
		t.Fatal("no candidates")
	}

	response, err := testApi.Candidate(responseCandidates.Result[0].PubKey, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", response)
}
