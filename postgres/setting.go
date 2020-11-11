package postgres

import (
	"github.com/GlitchyGlitch/typinger/models"

	"github.com/go-pg/pg"
)

type SettingRepo struct {
	DB *pg.DB
}

func (s *SettingRepo) GetSettings() ([]*models.Setting, error) {
	var settings []*models.Setting

	err := s.DB.Model(&settings).Order("id").Select()

	if err != nil {
		return nil, err
	}
	return settings, nil
}
