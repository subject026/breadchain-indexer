package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/subject026/breadchain-indexer/internal/database"
)

type User struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	WalletAddress string    `json:"wallet_address"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:            dbUser.ID,
		CreatedAt:     dbUser.CreatedAt,
		WalletAddress: dbUser.WalletAddress,
	}
}

type Vote struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID `json:"user_id"`
	ProjectID uuid.UUID `json:"project_id"`
	Value     int32     `json:"value"`
}

func databaseVoteToVote(dbVote database.Vote) Vote {
	return Vote{
		ID:        dbVote.ID,
		CreatedAt: dbVote.CreatedAt,
		UserID:    dbVote.UserID,
		ProjectID: dbVote.ProjectID,
		Value:     dbVote.Value,
	}
}

///////

func databaseVotesToVotes(dbVotes []database.Vote) []Vote {
	votes := []Vote{}
	for _, dbVote := range dbVotes {
		votes = append(votes, Vote{
			ID:        dbVote.ID,
			CreatedAt: dbVote.CreatedAt,
			UserID:    dbVote.UserID,
			ProjectID: dbVote.ProjectID,
			Value:     dbVote.Value,
		})
	}
	return votes
}

type Project struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	WalletAddress string    `json:"wallet_address"`
}

func databaseProjectToProject(dbProject database.Project) Project {
	return Project{
		ID:            dbProject.ID,
		CreatedAt:     dbProject.CreatedAt,
		Name:          dbProject.Name,
		WalletAddress: dbProject.WalletAddress,
	}
}

func databaseProjectsToProjects(dbProjects []database.Project) []Project {
	projects := []Project{}
	for _, dbProject := range dbProjects {
		projects = append(projects, Project{
			ID:            dbProject.ID,
			CreatedAt:     dbProject.CreatedAt,
			Name:          dbProject.Name,
			WalletAddress: dbProject.WalletAddress,
		})
	}
	return projects
}
