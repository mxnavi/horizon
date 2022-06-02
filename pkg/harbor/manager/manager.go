package manager

import (
	"context"

	"g.hz.netease.com/horizon/pkg/harbor/dao"
	"g.hz.netease.com/horizon/pkg/harbor/models"
)

var (
	// Mgr is the global region manager
	Mgr = New()
)

type Manager interface {
	// Create a harbor
	Create(ctx context.Context, harbor *models.Harbor) (uint, error)
	// Update update a harbor
	Update(ctx context.Context, id uint, harbor *models.Harbor) error
	// DeleteByID delete a harbor by id
	DeleteByID(ctx context.Context, id uint) error
	// GetByID get by id
	GetByID(ctx context.Context, id uint) (*models.Harbor, error)
	// ListAll list all harbors
	ListAll(ctx context.Context) ([]*models.Harbor, error)
}

type manager struct {
	harborDAO dao.DAO
}

func New() Manager {
	return &manager{
		harborDAO: dao.NewDAO(),
	}
}

func (m manager) Create(ctx context.Context, harbor *models.Harbor) (uint, error) {
	return m.harborDAO.Create(ctx, harbor)
}

func (m manager) GetByID(ctx context.Context, id uint) (*models.Harbor, error) {
	return m.harborDAO.GetByID(ctx, id)
}

func (m manager) ListAll(ctx context.Context) ([]*models.Harbor, error) {
	return m.harborDAO.ListAll(ctx)
}

func (m manager) Update(ctx context.Context, id uint, harbor *models.Harbor) error {
	return m.harborDAO.Update(ctx, id, harbor)
}

func (m manager) DeleteByID(ctx context.Context, id uint) error {
	return m.harborDAO.DeleteByID(ctx, id)
}
