package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/user/model"
)

	
func (r *Repository) InsertTwoFaSms(ctx context.Context, twoFaSms *model.UserTwoFaSms) error {
	return errors.New("user pg repo - InsertTwoFaSms unimplement")
}