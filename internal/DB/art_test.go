//go:build integration
// +build integration

package db

import (
	"ArtAPI/internal/artobj"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestArtDatabase(t *testing.T) {
	t.Run("test create art record", func(t *testing.T){
		db, err := NewDbConnection()
		assert.NoError(t, err)

		art, err := db.PostArt(context.Background(), artobj.ArtObject{
			ObjectID: 2,
			IsHighlight: false,
			AccessionYear: "5432",
			Department: "Test Dept",
			Title: "New Title",
			ObjectName: "Test Obj",
			Culture: "alien",
			Period: "future",
			City: "Mars",
			Country: "Bloop",
		})
		assert.NoError(t, err)

		newArt, err := db.GetArt(context.Background(), art.ID)
		assert.NoError(t, err)
		assert.Equal(t, art.Title, newArt.Title)

	})

	t.Run("test delete art record", func(t *testing.T) {
		db, err := NewDbConnection()
		assert.NoError(t, err)

		art, err := db.PostArt(context.Background(), artobj.ArtObject{
			ObjectID: 2,
			IsHighlight: false,
			AccessionYear: "5432",
			Department: "Test Dept delete",
			Title: "New Title",
			ObjectName: "Test Obj",
			Culture: "alien",
			Period: "future",
			City: "Mars",
			Country: "Bloop",
		})
		assert.NoError(t, err)

		err = db.DeleteArt(context.Background(), art.ID)
		assert.NoError(t, err)

		_, err = db.GetArt(context.Background(), art.ID)
		assert.Error(t, err)
	})

	t.Run("test update art record", func(t *testing.T) {
		db, err := NewDbConnection()
		assert.NoError(t, err)

		art, err := db.PostArt(context.Background(), artobj.ArtObject{
			ObjectID: 2,
			IsHighlight: false,
			AccessionYear: "5432",
			Department: "Test Dept delete",
			Title: "New Title",
			ObjectName: "Test Obj",
			Culture: "alien",
			Period: "future",
			City: "Mars",
			Country: "Bloop",
		})
		assert.NoError(t, err)

		art.Title = "Updated title"
		art.Culture = "Updated Culture"

		updatedArt, err := db.UpdateArt(context.Background(), art.ID, art)
		assert.NoError(t, err)

		getArt, err := db.GetArt(context.Background(), updatedArt.ID)
		assert.Equal(t, getArt, art)

	})
}