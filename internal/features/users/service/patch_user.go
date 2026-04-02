package usersservice

import (
	"context"
	"fmt"

	"github.com/977ADAM/golang-todoapp-project/internal/core/domain"
)

func (s *UsersService) PatchUser(
	ctx context.Context,
	id int,
	patch domain.UserPatch,
) (domain.User, error) {
	// 1. Достаем из репозитория USER с этим ID если юзера не существуюет
	// то возвращаем ошибку с пустым юзером.
	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from response: %w", err)
	}
	// если получили юзера
	// то применяем патч
	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("apply user patch: %w", err)
		// если ошибка возрощаем пустого юзера и ошибку
	}

	patchedUser, err := s.usersRepository.PatchUser(
		ctx,
		id,
		user,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("patch user: %w", err)
	}

	return patchedUser, nil
}
