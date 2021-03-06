package zendesk

import (
	"net/http"
	"testing"
)

func TestGetGroups(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "groups.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	groups, _, err := client.GetGroups()
	if err != nil {
		t.Fatalf("Failed to get groups: %s", err)
	}

	if len(groups) != 1 {
		t.Fatalf("expected length of groups is 1, but got %d", len(groups))
	}
}

func TestCreateGroup(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "groups.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateGroup(Group{})
	if err != nil {
		t.Fatalf("Failed to send request to create group: %s", err)
	}
}

func TestGetGroup(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "group.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	group, err := client.GetGroup(123)
	if err != nil {
		t.Fatalf("Failed to get group: %s", err)
	}

	expectedID := int64(360002440594)
	if group.ID != expectedID {
		t.Fatalf("Returned group does not have the expected ID %d. Group ID is %d", expectedID, group.ID)
	}
}
