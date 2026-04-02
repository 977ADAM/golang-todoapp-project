package domain

import (
	"fmt"
	"regexp"

	coreerrors "github.com/977ADAM/golang-todoapp-project/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func NewUser(
	id int,
	version int,
	fullName string,
	phoneNumber *string,
) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUnitialized(
	fullName string,
	phoneNumber *string,
) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		fullName,
		phoneNumber,
	)
}

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func NewUserPatch(
	fullName Nullable[string],
	phoneNumber Nullable[string],
) UserPatch {
	return UserPatch{
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func (u *User) Validate() error {
	fullNameLen := len([]rune(u.FullName))
	if fullNameLen < 3 || fullNameLen > 100 {
		return fmt.Errorf(
			"invalid `FullName` len: %d: %w",
			fullNameLen,
			coreerrors.ErrInvalidArgument,
		)
	}

	if u.PhoneNumber != nil {
		phoneNumberLen := len([]rune(*u.PhoneNumber))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf(
				"invalid `PhoneNumber` len: %d: %w",
				phoneNumberLen,
				coreerrors.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid `PhoneNumber` format: %w",
				coreerrors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

func (p *UserPatch) Validate() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf("`FullName` can't be patched to NULL: %w", coreerrors.ErrInvalidArgument)
	}

	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u

	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}

	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp

	return nil
}
