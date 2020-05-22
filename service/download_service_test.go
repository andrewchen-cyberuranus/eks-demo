package service_test

import (
	"eks-demo/model"
	"eks-demo/service"
	"os"
	"testing"
)

func TestGetDownloadURL(t *testing.T) {

	req := &model.GetDownloadURLRequest{
		Scheme: "https",
		Domain: os.Getenv("CF_DOMAIN_NAME"),
		File:   "test.jpg",
	}

	downloadSvc := service.NewDownloadService()

	resp, err := downloadSvc.GetDownloadURL(req)
	if err != nil {
		t.Fatalf("GetDownloadURL failed: %v", err)
	}
	t.Logf("GetDownloadURL: %v", resp)
}