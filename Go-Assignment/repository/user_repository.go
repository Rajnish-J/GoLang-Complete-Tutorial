package repository

import (
    "database/sql"
    "go-assignment/models"
)

type UserRepository interface {
    Create(user models.User) error
    GetByID(UserId string) (models.User, error)
    Update(user models.User) error
    Delete(UserId string) error
    GetAll() ([]models.User, error)
}

type userRepo struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepo{db}
}

func (r *userRepo) Create(user models.User) error {
    query := `INSERT INTO user_details 
        (user_id, first_name, last_name, dob, mobile_number, occupation, email, alternate_mobile_number, address) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

    _, err := r.db.Exec(query,
        user.UserId,
        user.FirstName,
        user.LastName,
        user.DateOfBirth,
        user.MobileNumber,
        user.Occupation,
        user.EmailId,
        user.AlternateMobileNumber,
        user.Address)

    return err
}

func (r *userRepo) GetByID(UserId string) (models.User, error) {
    var user models.User
    query := `SELECT * FROM user_details WHERE user_id = ?`
    err := r.db.QueryRow(query, UserId).Scan(
        &user.UserId,
        &user.FirstName,
        &user.LastName,
        &user.DateOfBirth,
        &user.MobileNumber,
        &user.Occupation,
        &user.EmailId,
        &user.AlternateMobileNumber,
        &user.Address,
    )
    return user, err
}

func (r *userRepo) Update(user models.User) error {
    dob := sql.NullString{}
    if user.DateOfBirth != "" {
        dob = sql.NullString{String: user.DateOfBirth, Valid: true}
    }

    query := `UPDATE user_details SET 
        first_name=?, last_name=?, dob=?, mobile_number=?, occupation=?, email=?, alternate_mobile_number=?, address=?
        WHERE user_id=?`

    _, err := r.db.Exec(query,
        user.FirstName,
        user.LastName,
        dob, // ✅ use sql.NullString
        user.MobileNumber,
        user.Occupation,
        user.EmailId,
        user.AlternateMobileNumber,
        user.Address,
        user.UserId)

    return err
}

func (r *userRepo) Delete(UserId string) error {
    _, err := r.db.Exec("DELETE FROM user_details WHERE user_id = ?", UserId)
    return err
}

func (r *userRepo) GetAll() ([]models.User, error) {
    rows, err := r.db.Query("SELECT * FROM user_details")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var u models.User
        err := rows.Scan(
            &u.UserId, &u.FirstName, &u.LastName, &u.DateOfBirth,
            &u.MobileNumber, &u.Occupation, &u.EmailId,
            &u.AlternateMobileNumber, &u.Address,
        )
        if err != nil {
            return nil, err
        }
        users = append(users, u)
    }

    return users, nil
}
