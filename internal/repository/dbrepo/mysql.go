package dbrepo

import (
	"bookings/internal/models"
	"context"
	"time"
)

func (m *mysqlDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *mysqlDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into reservations (first_name, last_name, email, phone, start_date,
			 end_date, room_id, created_at, updated_at)
			 values (?, ?, ?, ?, ?, ?, ?, ?, ?) returning id`

	var newID int
	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now().Local(),
		time.Now().Local(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// InsertRoomRestriction inserts a room restriction into database
func (m *mysqlDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into room_restrictions (start_date, end_date, room_id, reservation_id,
		created_at, updated_at, restriction_id)
		values (?, ?, ?, ?, ?, ?, ?)`
	_, err := m.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		time.Now().Local(),
		time.Now().Local(),
		r.RestrictionID,
	)

	if err != nil {
		return err
	}

	return nil
}
