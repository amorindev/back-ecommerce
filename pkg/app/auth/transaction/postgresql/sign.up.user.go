package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/user/model"
)

func (t *Transaction) SignUpUser(ctx context.Context, user *model.User) error {
	return errors.New("auth pg tx - SignUpUser unimplement")
}
