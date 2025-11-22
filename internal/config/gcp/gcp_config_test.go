package gcp

import "testing"

func TestSetRegionOrLocation(t *testing.T) {
	gcp := GoogleCloudConfig{
		"",
		"",
	}
	input := "eu-west1"
	gcp.SetRegionOrLocation(input)

	if gcp.Region != input {
		t.Errorf("expected %v received %v", input, gcp.Region)
	}
}

func TestSetProjectId(t *testing.T) {
	gcp := GoogleCloudConfig{
		"",
		"",
	}
	input := "test-project-name"
	gcp.SetProjectOrSubscription(input)

	if gcp.ProjectId != input {
		t.Errorf("expected %v received %v", input, gcp.Region)
	}
}

func TestSave(t *testing.T) {
	gcp := GoogleCloudConfig{
		"test-save-project",
		"test-save-region",
	}

	if err := gcp.Save(); err != nil {
		t.Errorf("error saving file %v", err)
	}

	var readGcp GoogleCloudConfig

	if err := readGcp.Read(); err != nil {
		t.Errorf("error reading saved file %v", err)
	}

	if readGcp.ProjectId != gcp.ProjectId {
		t.Errorf("expected ProjectId %v, got %v", gcp.ProjectId, readGcp.ProjectId)
	}

	if readGcp.Region != gcp.Region {
		t.Errorf("expected Region %v, got %v", gcp.Region, readGcp.Region)
	}
}

func TestRead(t *testing.T) {
	gcp := GoogleCloudConfig{
		"test-read-project",
		"test-read-region",
	}

	if err := gcp.Save(); err != nil {
		t.Errorf("error saving file for read test %v", err)
	}

	var readGcp GoogleCloudConfig
	if err := readGcp.Read(); err != nil {
		t.Errorf("error reading file %v", err)
	}

	if readGcp.ProjectId != gcp.ProjectId {
		t.Errorf("expected ProjectId %v, got %v", gcp.ProjectId, readGcp.ProjectId)
	}

	if readGcp.Region != gcp.Region {
		t.Errorf("expected Region %v, got %v", gcp.Region, readGcp.Region)
	}
}
