package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testBridgeName = "libsacloud_test_archive"

func TestBridgeCRUD(t *testing.T) {
	defer initBridge()()

	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	api := client.Bridge

	//CREATE
	newBr := api.New()
	newBr.Name = testBridgeName
	newBr.Description = "before"

	br, err := api.Create(newBr)

	assert.NoError(t, err)
	assert.NotEmpty(t, br)

	id := br.ID

	//READ
	br, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, br)

	//UPDATE
	br.Description = "after"
	br, err = api.Update(id, br)

	assert.NoError(t, err)
	assert.NotEqual(t, br.Description, "before")

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func initBridge() func() {
	cleanupBridge()
	return cleanupBridge
}

func cleanupBridge() {
	items, _ := client.Bridge.Reset().WithNameLike(testBridgeName).Find()
	for _, item := range items.Bridges {
		client.Bridge.Delete(item.ID)
	}
}
