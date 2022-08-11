package main

import (
	"github.com/sshintaku/prisma_session"
)

func main() {
	// Create a generic session for running API calls to Prisma Cloud
	var params = prisma_session.ReadParameters()
	session := prisma_session.Session{ApiUrl: params.ApiUrl}
	session.CreateSession()

	// Gathering all image data in the environment
	imageData := session.GetDeployedImages()
	prisma_session.GetDataByCollection(params.CWPQueryFilters.Collections, imageData)
}
