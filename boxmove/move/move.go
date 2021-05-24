package move

import "time"

type Move struct {
	repo Storage
	dto  DTO
	id   string
}

type View struct {
	ID        string     `json:"id"`
	FromID    string     `json:"from_id"`
	ToID      string     `json:"to_id"`
	Date      time.Time  `json:"date"`
	Count     uint64     `json:"count"`
	CreatedAt time.Time  `json:"create_at"`
}

func (t *Move) View() *View {
	return t.dto.View()
}

func (t *Move) Refresh() error {
	dto, err := t.repo.FindById(t.id)
	if err != nil {
		return err
	}

	t.dto = dto

	return nil
}

func (t *Move) EqualTo(other *Move) bool {
	return t.id == other.id
}