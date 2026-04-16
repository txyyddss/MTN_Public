package monitor

import (
	"reflect"
	"testing"

	"github.com/mcstatus-io/mcutil/v4/formatting"
	"github.com/mcstatus-io/mcutil/v4/response"
)

func TestSplitPlayers(t *testing.T) {
	sample := []response.SamplePlayer{
		{ID: "uuid1", Name: formatting.Result{Clean: "JavaPlayer"}},
		{ID: "uuid2", Name: formatting.Result{Clean: ".BedrockPlayer1"}},
		{ID: "uuid3", Name: formatting.Result{Clean: "BE_BedrockPlayer2"}},
		{ID: "00000000-0000-0000-0000-000000000000", Name: formatting.Result{Clean: "Bot"}},
		{ID: "uuid4", Name: formatting.Result{Clean: ""}},
	}

	java, bedrock, uuids := splitPlayers(sample)

	expectedJava := []string{"JavaPlayer"}
	expectedBedrock := []string{".BedrockPlayer1", "BE_BedrockPlayer2"}
	expectedUUIDs := []string{"uuid1", "uuid2", "uuid3", "uuid4"}

	if !reflect.DeepEqual(java, expectedJava) {
		t.Errorf("expected java %v, got %v", expectedJava, java)
	}

	if !reflect.DeepEqual(bedrock, expectedBedrock) {
		t.Errorf("expected bedrock %v, got %v", expectedBedrock, bedrock)
	}

	if !reflect.DeepEqual(uuids, expectedUUIDs) {
		t.Errorf("expected uuids %v, got %v", expectedUUIDs, uuids)
	}
}
