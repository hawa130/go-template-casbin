package database

import (
	"context"

	"github.com/hawa130/serverx/internal/perm"
)

func seedData(ctx context.Context) error {
	seeds := []func(context.Context) error{
		seedUser,
	}

	for _, seed := range seeds {
		if err := seed(ctx); err != nil {
			return err
		}
	}

	return nil
}

func seedUser(ctx context.Context) error {
	exists, err := client.User.Query().Exist(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	if err != nil {
		return err
	}

	user, err := client.User.Create().
		SetPhone("12345678910").
		SetPassword("root").
		Save(ctx)
	if err != nil {
		return err
	}

	_, err = perm.AddSubjectRole(user.ID.String(), "root")
	if err != nil {
		return err
	}

	return nil
}
