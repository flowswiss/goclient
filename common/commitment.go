package common

import (
	"context"
	"encoding/json"
	"time"

	"github.com/flowswiss/goclient/v2/core"
)

type Commitment struct {
	ID             int                 `json:"id"`
	Type           string              `json:"type"`
	Reference      CommitmentReference `json:"reference"`
	StartDate      time.Time           `json:"start_date"`
	EndDate        time.Time           `json:"end_date"`
	Renew          bool                `json:"renew"`
	AdditionalData string              `json:"additional_data"`
	Cycle          CommitmentCycle     `json:"cycle"`
}

type CommitmentReference struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type CommitmentCycle struct {
	ID   int    `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

type CommitmentService struct {
	client *core.Client
}

func NewCommitmentService(client *core.Client) *CommitmentService {
	return &CommitmentService{client: client}
}

func (c CommitmentService) List(ctx context.Context, cursor core.Cursor) (list List[Commitment], err error) {
	list.Pagination, err = c.client.List(ctx, getCommitmentsPath(), cursor, &list.Items)
	return
}

type CommitmentUpdateReq struct {
	ID uint `json:"-"`

	Renew          bool            `json:"renew"`
	AdditionalData json.RawMessage `json:"additional_data"`
}

func (c CommitmentService) Update(ctx context.Context, req CommitmentUpdateReq) (commitment Commitment, err error) {
	err = c.client.Update(ctx, getSpecificCommitmentPath(req.ID), req, &commitment)
	return
}

const commitmentsSegment = "/v4/commitments"

func getCommitmentsPath() string {
	return commitmentsSegment
}

func getSpecificCommitmentPath(commitmentID uint) string {
	return core.Join(commitmentsSegment, commitmentID)
}
