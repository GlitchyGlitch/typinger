package main

import (
	"context"
	"testing"

	"github.com/GlitchyGlitch/typinger/config"
	"github.com/GlitchyGlitch/typinger/test"
	"github.com/shurcooL/graphql"
	"github.com/stretchr/testify/require"
)

func TestMutationUsers(t *testing.T) {
	conf := config.New()
	conf.Port = "8005"
	c := setup(false, conf)
	defer teardown(conf)

	t.Run("Create user forbidden", func(t *testing.T) {
		var mutation struct {
			User struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"createUser(input: {name:\"Bjarne Stroustrup\", email:\"stroustrup@gmail.com\", password:\"stroustrup\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Update existing user forbidden", func(t *testing.T) {
		var mutation struct {
			User struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"updateUser(id:\"0e38a4bd-87a0-447f-93fd-b904c9f7f303\" input: {name:\"Bjarne Stroustrup\", email:\"stroustrup@gmail.com\", password:\"stroustrup\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Update user not found forbidden", func(t *testing.T) {
		var mutation struct {
			User struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"updateUser(id:\"b0592654-ac3d-4798-baea-3fb9b86a81c8\" input: {name:\"Fourth User\", email:\"fourth@example.com\", password:\"fourthxx\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Delete existing user forbidden", func(t *testing.T) {
		var mutation struct {
			Deleted graphql.Boolean `graphql:"deleteUser(id:\"0e38a4bd-87a0-447f-93fd-b904c9f7f303\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Delete nonexistent user forbidden", func(t *testing.T) {
		var mutation struct {
			Deleted graphql.Boolean `graphql:"deleteUser(id:\"b0592654-ac3d-4798-baea-3fb9b86a81c8\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})
}

func TestMutationUsersAuthenticated(t *testing.T) {
	conf := config.New()
	conf.Port = "8006"
	c := setup(true, conf)
	defer teardown(conf)

	t.Run("Create user", func(t *testing.T) {
		var mutation struct {
			User struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"createUser(input: {name:\"Fourth User\", email:\"fourth@example.com\", password:\"fourthxx\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.NoError(t, err)

		require.Equal(t, "User", string(mutation.User.Typename))
		require.True(t, test.IsValidUUID(string(mutation.User.ID)))
		require.Equal(t, "Fourth User", string(mutation.User.Name))
		require.Equal(t, "fourth@example.com", string(mutation.User.Email))
	})

	t.Run("Create two same users", func(t *testing.T) {
		var mutation struct {
			User struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"createUser(input: {name:\"Fourth User\", email:\"fourth@example.com\", password:\"fourthxx\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Resource already exists.")
	})

	t.Run("Create user with too long name", func(t *testing.T) {
		var mutation struct {
			User struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"createUser(input: {name:\"nlvsfPdxtXAwLUNEhJiyRcs1xWRjES7TIaSpKHyJbGfxGWFJDdHCq0iDykUe2Gaa33lakk7ViFfaSa2BqJovX2lEuMprg0ZHH9pSgfV0A06xvwIDEhHd8KtZ03DOkTq0WdPwJORMtDQ0JZGSZcsHc6kHC6syFdYaTiCGjZKLioQIyi4Wb4Mk20zG0fsCNv7wS4BkA5MrtiYDhYmGhasH8mAHIn8AT2BoohINHR1WGm4AbyE5o5XwKfRzLoC7a1JJG\", email:\"stroustrup@gmail.com\", password:\"st\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Field \"name\" is invalid.")
	})

	t.Run("Create user with invalid email", func(t *testing.T) {
		var mutation struct {
			User struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"createUser(input: {name:\"Fourth User\", email:\"fourth@\", password:\"fourth\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Field \"email\" is invalid.")
	})

	t.Run("Create user with too short password", func(t *testing.T) {
		var mutation struct {
			User struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"createUser(input: {name:\"Fourth User\", email:\"fourth@example.com\", password:\"fo\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Field \"password\" is invalid.")
	})

	t.Run("Update existing user", func(t *testing.T) {
		var mutation struct {
			User struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"updateUser(id:\"0e38a4bd-87a0-447f-93fd-b904c9f7f303\" input: {name:\"Third User Updated\", email:\"thirdupdated@example.com\", password:\"thirdupdated\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.NoError(t, err)

		require.Equal(t, "User", string(mutation.User.Typename))
		require.True(t, test.IsValidUUID(string(mutation.User.ID)))
		require.Equal(t, "Third User Updated", string(mutation.User.Name))
		require.Equal(t, "thirdupdated@example.com", string(mutation.User.Email))
	})

	t.Run("Update nonexistent user", func(t *testing.T) {
		var mutation struct {
			User struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"updateUser(id:\"b0592654-ac3d-4798-baea-3fb9b86a81c8\" input: {name:\"Third User Updated\", email:\"thirdupdated@example.com\", password:\"thirdupdated\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "No data found.")
	})
	t.Run("Update user with invalid id", func(t *testing.T) {
		var mutation struct {
			User struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"updateUser(id:\"b0592654-ac3d-4798-baea-3fb9b81c8\" input: {name:\"Third User Updated\", email:\"thirdupdated@example.com\", password:\"thirdupdated\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Field \"id\" is invalid.")
	})

	t.Run("Delete existing user", func(t *testing.T) {
		var mutation struct {
			Deleted graphql.Boolean `graphql:"deleteUser(id:\"0e38a4bd-87a0-447f-93fd-b904c9f7f303\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.NoError(t, err)

		require.True(t, bool(mutation.Deleted))
	})

	t.Run("Delete user not found", func(t *testing.T) {
		var mutation struct {
			Deleted graphql.Boolean `graphql:"deleteUser(id:\"b0592654-ac3d-4798-baea-3fb9b86a81c8\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "No data found.")
	})

	t.Run("Delete user with invalid id", func(t *testing.T) {
		var mutation struct {
			Deleted graphql.Boolean `graphql:"deleteUser(id:\"b059654\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Field \"id\" is invalid.")
	})
}

func TestMutationArticles(t *testing.T) { // TODO: update test data for articles
	conf := config.New()
	conf.Port = "8001"
	c := setup(false, conf)
	defer teardown(conf)

	t.Run("Creating artcile forbidden", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Fourth article\", content:\"Fourth content.\", thumbnailUrl:\"http://www.example.com/path/to/photo4.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)

		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Updateing article forbidden", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Fourth article\", content:\"Fourth content.\", thumbnailUrl:\"http://www.example.com/path/to/photo4.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)

		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Deleting artcile forbidden", func(t *testing.T) {
		var mutation struct {
			Deleted bool `graphql:"deleteArticle(id:\"d50f5d60-6f59-4605-96b8-a96b9e9b17ea\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)

		require.EqualError(t, err, "Operation forbidden.")
		require.False(t, mutation.Deleted)
	})
}

func TestMutationArticlesAuthenticated(t *testing.T) {
	conf := config.New()
	conf.Port = "8002"
	c := setup(true, conf)
	defer teardown(conf)

	t.Run("Create article", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Fourth article\", content:\"Fourth content.\", thumbnailUrl:\"http://www.example.com/path/to/photo4.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)

		require.NoError(t, err)
		require.Equal(t, "Article", string(mutation.Article.Typename))
		require.True(t, test.IsValidUUID(string(mutation.Article.ID)))
		require.Equal(t, "Fourth article", string(mutation.Article.Title))
		require.Equal(t, "Fourth content.", string(mutation.Article.Content))
		require.Equal(t, "http://www.example.com/path/to/photo4.jpg", string(mutation.Article.ThumbnailURL))
	})

	t.Run("Create two same articles", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Fifth article\", content:\"Fifth content.\", thumbnailUrl:\"http://www.example.com/path/to/photo5.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.NoError(t, err)

		err = c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Resource already exists.")
	})

	t.Run("Create article with no title", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"\", content:\"Fourth article.\", thumbnailUrl:\"http://www.example.com/path/to/photo4.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Field \"title\" is invalid.")
	})

	t.Run("Create article with no content", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Fourth article\", content:\"\", thumbnailUrl:\"http://www.example.com/path/to/photo4.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Field \"content\" is invalid.")
	})

	t.Run("Create article with invalid thumbnailUrl", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Fourth article\", content:\"Fourth content.\", thumbnailUrl:\"invalidurl\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Field \"thumbnailUrl\" is invalid.")
	})

	t.Run("Create article with no thumbnailUrl", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Fourth article\", content:\"Fourth content.\", thumbnailUrl:\"invalidurl\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Field \"thumbnailUrl\" is invalid.")
	})

	t.Run("Delete article", func(t *testing.T) {
		var mutation struct {
			Deleted graphql.Boolean `graphql:"deleteArticle(id:\"82ba242e-e853-499f-8873-f271c53aca6a\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.NoError(t, err)
		require.True(t, bool(mutation.Deleted))
	})

	t.Run("Delete article not found", func(t *testing.T) {
		var mutation struct {
			Deleted graphql.Boolean `graphql:"deleteArticle(id:\"e5f1c9af-fa8a-4a58-9909-d887ddf7e961\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "No data found.")
		require.False(t, bool(mutation.Deleted))
	})

	t.Run("Delete article invalid id", func(t *testing.T) {
		var mutation struct {
			Deleted graphql.Boolean `graphql:"deleteArticle(id:\"invaliduuid\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Field \"id\" is invalid.")
		require.False(t, bool(mutation.Deleted))
	})
}

//TODO: Check password change process.
//TODO: Make sure user doesn't exist after delete.
