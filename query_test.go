package main

import (
	"context"
	"testing"

	"github.com/GlitchyGlitch/typinger/config"
	"github.com/stretchr/testify/require"
)

func TestQueryArticles(t *testing.T) {
	conf := config.New()
	conf.Port = "8000"
	c := setup(false, conf)
	defer teardown(conf)

	t.Run("Get all articles", func(t *testing.T) {
		var query struct {
			Articles []article
		}
		err := c.Query(context.Background(), &query, nil)
		require.NoError(t, err)

		expected := articlesDataSet
		require.Equal(t, query.Articles, expected)
	})

	t.Run("Get articles with pagination", func(t *testing.T) {
		var query struct {
			Articles []article `graphql:"articles(first:1, offset:2)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.NoError(t, err)

		expected := []article{articlesDataSet[2]}
		require.Equal(t, query.Articles, expected)
	})

	t.Run("Get articles with filter", func(t *testing.T) {
		var query struct {
			Articles []article `graphql:"articles(filter: {title: \"sec\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.NoError(t, err)

		expected := []article{articlesDataSet[1]}
		require.Equal(t, query.Articles, expected)
	})

	t.Run("No matches to filter", func(t *testing.T) {
		var query struct {
			Articles []article `graphql:"articles(filter: {title: \"nonexistenttitle\"})"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)
		require.Len(t, query.Articles, 0)
	})

	t.Run("Invalid attribute of filter", func(t *testing.T) {
		var query struct {
			Articles []article `graphql:"articles(filter: {title: \"PtYaKbEfujgg012vzU54xnIO8uNFyrWnzT1s8qp469ktBIuyGAMV5DkIteDogufej8rqyOIv73H9dakOfD7gPoMk850GngJA17MyolQ39VwOzV65XlktlOQPf17MJPW56tcSFYuwUmh4tgKZiwYQ8TwC5onHlKiZtCBMVEwdV9Seb6ZzOk0ccvXP0NVStib6eghVeGPOPPtpVplgXasr4cYAp12TWCHXoxjE4JxKZR0c7CqCqQpziJGehZfyEBf9l\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Field \"title\" is invalid.")
	})

	t.Run("Invalid pagination first field", func(t *testing.T) {
		var query struct {
			Articles []article `graphql:"articles(first:-1, offset:1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Field \"first\" is invalid.")
	})

	t.Run("Invalid pagination offset field", func(t *testing.T) {
		var query struct {
			Articles []article `graphql:"articles(first:1, offset:-1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Field \"offset\" is invalid.")
	})
}

func TestQueryUsers(t *testing.T) {
	conf := config.New()
	conf.Port = "8003"
	c := setup(false, conf)
	defer teardown(conf)

	t.Run("Get users forbidden", func(t *testing.T) {
		var query struct {
			Users []user
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("Get users with filter forbidden", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(filter: {name:\"fir\", email:\"fir\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("Get users with pagination forbidden", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(first: 2, offset: 1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("No matches to filter forbidden", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(filter:{name:\"nonexistentname\", email:\"nonexistentemail\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("Invalid name attribute of filter forbidden", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(filter: {name: \"AAZN90ESauHQynfvljT0L9JTamLV1yKThRJhJI4AusfeofWzylykVp1QwTlhVWTkOiYlssW3DEV3HUJFR7zZmkBV5GY2tZTXB4gUPITPsZv7fVbqzPu4D59ac4yTYZu2o\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("Invalid email attribute of filter forbidden", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(filter: {email: \"cL3InRIpfwxbT9POCLM7p11aYIyyzIM3iX16aozwtY0wYXIaMZ6yGlE4O7eQY9l8B2UQK7Qt7bU2gExh9WqLquqkbLWY3a8sBKTD5MY3BfcPs46c6UmzcYxmryyqUjq3rYy4mBsH2v0Sqx4G4sCdC9ioNJJvqsfuTsYVBl1JoS8xhGfVgBCmeRG8xlH1agzQgV0GFNrL3dF9fjhoWwHPPttlh6FWChdZ4dm2xdv3o8W64umcSePUSCgHCf8se0TbTgvFt32zDzQp6gAMWmCPbjCNlEZI3jHpREeTf0e2n8xD9MstK7j5pWFZUokDSp3jx\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("Invalid pagination first field forbidden", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(first:-1, offset:1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("Invalid pagination offset field forbidden", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(first:1, offset:-1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})
}

func TestQueryUsersAuthenticated(t *testing.T) {
	conf := config.New()
	conf.Port = "8004"
	c := setup(true, conf)
	defer teardown(conf)

	t.Run("Get users", func(t *testing.T) {
		var query struct {
			Users []user
		}
		err := c.Query(context.Background(), &query, nil)
		require.NoError(t, err)

		expected := userDataSet
		require.Equal(t, query.Users, expected)
	})

	t.Run("Get users with filter", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(filter:{name:\"second\", email:\"ex\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.NoError(t, err)

		expected := []user{userDataSet[1]}
		require.Equal(t, query.Users, expected)
	})

	t.Run("Get users with pagination", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(first:1, offset:1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.NoError(t, err)

		expected := []user{userDataSet[1]}
		require.Equal(t, query.Users, expected)
	})

	t.Run("No matches to filter", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(filter:{name:\"nonexistentname\", email:\"nonexistentemail\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.NoError(t, err)

		require.Len(t, query.Users, 0)
	})

	t.Run("Invalid name attribute of filter", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(filter: {name: \"AAZN90ESauHQynfvljT0L9JTamLV1yKThRJhJI4AusfeofWzylykVp1QwTlhVWTkOiYlssW3DEV3HUJFR7zZmkBV5GY2tZTXB4gUPITPsZv7fVbqzPu4D59ac4yTYZu2o\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Field \"name\" is invalid.")
	})

	t.Run("Invalid email attribute of filter", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(filter: {email: \"cL3InRIpfwxbT9POCLM7p11aYIyyzIM3iX16aozwtY0wYXIaMZ6yGlE4O7eQY9l8B2UQK7Qt7bU2gExh9WqLquqkbLWY3a8sBKTD5MY3BfcPs46c6UmzcYxmryyqUjq3rYy4mBsH2v0Sqx4G4sCdC9ioNJJvqsfuTsYVBl1JoS8xhGfVgBCmeRG8xlH1agzQgV0GFNrL3dF9fjhoWwHPPttlh6FWChdZ4dm2xdv3o8W64umcSePUSCgHCf8se0TbTgvFt32zDzQp6gAMWmCPbjCNlEZI3jHpREeTf0e2n8xD9MstK7j5pWFZUokDSp3jx\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Field \"email\" is invalid.")
	})

	t.Run("Invalid pagination first field", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(first:-1, offset:1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Field \"first\" is invalid.")
	})

	t.Run("Invalid pagination offset field", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(first:1, offset:-1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Field \"offset\" is invalid.")
	})
}

func TestQueryImage(t *testing.T) {
	conf := config.New()
	conf.Port = "8007"
	c := setup(false, conf)
	defer teardown(conf)

	t.Run("Get images forbidden", func(t *testing.T) {
		var query struct {
			Images []image
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("Get images with pagination forbidden", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(first:1, offset:1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("Get images with filter forbidden", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(filter: {name:\"sec\", slug:\"slu\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("No matches to filter", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(filter:{name:\"nonexistentname\", slug:\"nonexistentslug\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("Invalid name attribute of filter forbidden", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(filter: {name: \"AAZN90ESauHQynfvljT0L9JTamLV1yKThRJhJI4AusfeofWzylykVp1QwTlhVWTkOiYlssW3DEV3HUJFR7zZmkBV5GY2tZTXB4gUPITPsZv7fVbqzPu4D59ac4yTYZu2o\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("Invalid slug attribute of filter forbidden", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(filter: {slug: \"eWUFMDK0FyLmKMhwRR5FvCQoE6ZFlUQ7YzvEx6Sae61DZ8ETBkXMWFW5VMV1758rHlet3ohLptyi4ys1UJ728n5LfG0hLbdjQcWSwdqzTNj9203C3BSg2MyG8zm5QEy1nRaVCz4OWXRvtQw3ryAiWdUfLOyyOpgOXxkf2SWd1ADMUoogP4HWkd2nsZrX1vWRI916oo3efqRcMgqtJCvqAOWhDDOyu4AIf0nariyqq44LuI5yT6F1To9jnLoip1X0I\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("Invalid pagination first field forbidden", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(first:-1, offset:1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})

	t.Run("Invalid pagination offset field forbidden", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(first:1, offset:-1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, forbiddenError)
	})
}
func TestQueryImageAuthentication(t *testing.T) {
	conf := config.New()
	conf.Port = "8008"
	c := setup(true, conf)
	defer teardown(conf)

	t.Run("Get images", func(t *testing.T) {
		var query struct {
			Images []image
		}
		err := c.Query(context.Background(), &query, nil)
		require.NoError(t, err)

		expected := imageDataSet
		require.Equal(t, query.Images, expected)
	})

	t.Run("Get images with pagination", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(first:1, offset:1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.NoError(t, err)

		expected := []image{imageDataSet[1]}
		require.Equal(t, query.Images, expected)
	})

	t.Run("Get images with filter", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(filter:{name:\"sec\", slug:\"slu\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.NoError(t, err)

		expected := []image{imageDataSet[1]}
		require.Equal(t, query.Images, expected)
	})

	t.Run("No matches to filter", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(filter:{name:\"nonexistentname\", slug:\"nonexistentslug\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.NoError(t, err)
		require.Len(t, query.Images, 0)
	})

	t.Run("Invalid name attribute of filter", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(filter: {name: \"AAZN90ESauHQynfvljT0L9JTamLV1yKThRJhJI4AusfeofWzylykVp1QwTlhVWTkOiYlssW3DEV3HUJFR7zZmkBV5GY2tZTXB4gUPITPsZv7fVbqzPu4D59ac4yTYZu2o\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Field \"name\" is invalid.")
	})

	t.Run("Invalid slug attribute of filter", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(filter: {slug: \"eWUFMDK0FyLmKMhwRR5FvCQoE6ZFlUQ7YzvEx6Sae61DZ8ETBkXMWFW5VMV1758rHlet3ohLptyi4ys1UJ728n5LfG0hLbdjQcWSwdqzTNj9203C3BSg2MyG8zm5QEy1nRaVCz4OWXRvtQw3ryAiWdUfLOyyOpgOXxkf2SWd1ADMUoogP4HWkd2nsZrX1vWRI916oo3efqRcMgqtJCvqAOWhDDOyu4AIf0nariyqq44LuI5yT6F1To9jnLoip1X0I\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Field \"slug\" is invalid.")
	})

	t.Run("Invalid pagination first field", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(first:-1, offset:1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Field \"first\" is invalid.")
	})

	t.Run("Invalid pagination offset field", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(first:1, offset:-1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Field \"offset\" is invalid.")
	})
}
