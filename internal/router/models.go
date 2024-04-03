package router

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
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	WalletAddress string    `json:"walletAddress"`
	ProjectID     uuid.UUID `json:"projectId"`
	Value         int32     `json:"value"`
}

func databaseVoteToVote(dbVote database.GetVotesInRangeRow) Vote {
	return Vote{
		ID:            dbVote.ID,
		CreatedAt:     dbVote.CreatedAt,
		WalletAddress: dbVote.WalletAddress,
		ProjectID:     dbVote.ProjectID,
		Value:         dbVote.Value,
	}
}

///////

func databaseVotesToVotes(dbVotes []database.GetVotesInRangeRow) []Vote {
	votes := []Vote{}
	for _, dbVote := range dbVotes {
		votes = append(votes, Vote{
			ID:            dbVote.ID,
			CreatedAt:     dbVote.CreatedAt,
			WalletAddress: dbVote.WalletAddress,
			ProjectID:     dbVote.ProjectID,
			Value:         dbVote.Value,
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

type Slice struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	StartedAt time.Time `json:"started_at"`
}

func databaseSliceToSlice(dbSlice database.Slice) Slice {
	return Slice{
		ID:        dbSlice.ID,
		CreatedAt: dbSlice.CreatedAt,
		StartedAt: dbSlice.StartedAt,
	}
}

func databaseSlicesToSlices(dbSlices []database.GetSlicesRow) []Slice {
	slices := []Slice{}
	for _, dbSlice := range dbSlices {
		slices = append(slices, Slice{
			ID:        dbSlice.ID,
			CreatedAt: dbSlice.CreatedAt,
			StartedAt: dbSlice.StartedAt,
		})
	}
	return slices
}

type SliceProject struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	ProjectID uuid.UUID `json:"project_id"`
	SliceID   uuid.UUID `json:"slice_id"`
	Value     int32     `json:"value"`
}

func databaseSliceProjectToSliceProject(dbSliceProject database.SliceProject) SliceProject {
	return SliceProject{
		ID:        dbSliceProject.ID,
		CreatedAt: dbSliceProject.CreatedAt,
		ProjectID: dbSliceProject.ProjectID,
		SliceID:   dbSliceProject.SliceID,
		Value:     dbSliceProject.Value,
	}
}

func databaseSliceProjectsToSliceProjects(dbSliceProjects []database.SliceProject) []SliceProject {
	sliceProjects := []SliceProject{}
	for _, dbSliceProject := range dbSliceProjects {
		sliceProjects = append(sliceProjects, SliceProject{
			ID:        dbSliceProject.ID,
			CreatedAt: dbSliceProject.CreatedAt,
			ProjectID: dbSliceProject.ProjectID,
			SliceID:   dbSliceProject.SliceID,
			Value:     dbSliceProject.Value,
		})
	}
	return sliceProjects
}
