package model

import (
	"fmt"
	"spl-users/ent"
)

type UserSearch struct {
	Run       string `json:"run"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
}

func EntUsersToUserSearch(user []*ent.User) []*UserSearch {
	var usersSearch []*UserSearch

	for _, user := range user {
		usersSearch = append(usersSearch, &UserSearch{
			Run:       fmt.Sprintf("%d-%s", user.Run, user.VerificationDigit),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Gender:    string(user.Gender),
		})
	}
	return usersSearch
}
