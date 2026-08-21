package streaklevel

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	error_responses "kaifin_clone_api/pkg/responses"
)

const XPPerStreakDay = 100

type LevelRepo interface {
	ListLevels() ([]LevelInfo, *error_responses.ErrorResponse)
	GetStatus(userID int64) (*UserLevelStatus, *error_responses.ErrorResponse)
	CheckIn(userID int64) (*UserLevelStatus, *error_responses.ErrorResponse)
	GetLeaderboard(limit int) ([]LeaderboardUser, *error_responses.ErrorResponse)
}

type LevelRepoImpl struct {
	dbpool *sqlx.DB
}

func NewLevelRepoImpl(db *sqlx.DB) *LevelRepoImpl {
	return &LevelRepoImpl{dbpool: db}
}

func (r *LevelRepoImpl) ListLevels() ([]LevelInfo, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var levels []LevelInfo

	err := r.dbpool.Select(&levels, `
		SELECT id, level, required_xp, title, badge_icon
		FROM tbl_user_levels
		ORDER BY required_xp ASC
	`)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	return levels, nil
}

func (r *LevelRepoImpl) getOrCreateStreak(userID int64) (*UserStreak, error) {
	var s UserStreak
	err := r.dbpool.Get(&s, `
		SELECT id, user_id, current_streak, longest_streak, last_active_date, updated_at
		FROM tbl_user_streaks
		WHERE user_id = $1
	`, userID)

	if err == nil {
		return &s, nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	err = r.dbpool.QueryRow(`
		INSERT INTO tbl_user_streaks (user_id, current_streak, longest_streak, last_active_date, updated_at)
		VALUES ($1, 0, 0, NULL, NOW())
		RETURNING id, user_id, current_streak, longest_streak, last_active_date, updated_at
	`, userID).Scan(&s.ID, &s.UserID, &s.CurrentStreak, &s.LongestStreak, &s.LastActiveDate, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *LevelRepoImpl) buildStatus(s *UserStreak, checkedInToday bool) (*UserLevelStatus, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	xp := s.LongestStreak * XPPerStreakDay

	var current LevelInfo
	err := r.dbpool.Get(&current, `
		SELECT id, level, required_xp, title, badge_icon
		FROM tbl_user_levels
		WHERE required_xp <= $1
		ORDER BY required_xp DESC
		LIMIT 1
	`, xp)

	var currentLevelPtr *LevelInfo
	if err == nil {
		currentLevelPtr = &current
	} else if !errors.Is(err, sql.ErrNoRows) {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	var next LevelInfo
	var nextLevelPtr *LevelInfo
	var xpToNext *int
	err = r.dbpool.Get(&next, `
		SELECT id, level, required_xp, title, badge_icon
		FROM tbl_user_levels
		WHERE required_xp > $1
		ORDER BY required_xp ASC
		LIMIT 1
	`, xp)
	if err == nil {
		nextLevelPtr = &next
		diff := next.RequiredXP - xp
		xpToNext = &diff
	} else if !errors.Is(err, sql.ErrNoRows) {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	return &UserLevelStatus{
		CurrentStreak:  s.CurrentStreak,
		LongestStreak:  s.LongestStreak,
		LastActiveDate: s.LastActiveDate,
		XP:             xp,
		CheckedInToday: checkedInToday,
		CurrentLevel:   currentLevelPtr,
		NextLevel:      nextLevelPtr,
		XPToNextLevel:  xpToNext,
	}, nil
}

func (r *LevelRepoImpl) GetStatus(userID int64) (*UserLevelStatus, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	s, err := r.getOrCreateStreak(userID)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	checkedInToday := isSameDay(s.LastActiveDate, time.Now())
	return r.buildStatus(s, checkedInToday)
}

func (r *LevelRepoImpl) CheckIn(userID int64) (*UserLevelStatus, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	s, err := r.getOrCreateStreak(userID)
	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}

	now := time.Now()
	today := truncateToDate(now)

	if isSameDay(s.LastActiveDate, now) {
		return r.buildStatus(s, true)
	}

	newCurrent := 1

	if s.LastActiveDate != nil {
		yesterday := today.AddDate(0, 0, -1)
		lastDate := truncateToDate(*s.LastActiveDate)
		if lastDate.Equal(yesterday) {
			newCurrent = s.CurrentStreak + 1
		}
	}

	newLongest := s.LongestStreak
	if newCurrent > newLongest {
		newLongest = newCurrent
	}

	_, err = r.dbpool.Exec(`
		UPDATE tbl_user_streaks
		SET current_streak = $1, longest_streak = $2, last_active_date = $3, updated_at = NOW()
		WHERE user_id = $4
	`, newCurrent, newLongest, today, userID)
	if err != nil {
		fmt.Println("CHECKIN UPDATE ERROR:", err)
		return nil, msg.NewErrorResponse("database_error", err)
	}

	s.CurrentStreak = newCurrent
	s.LongestStreak = newLongest
	s.LastActiveDate = &today

	return r.buildStatus(s, true)
}

func truncateToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func isSameDay(t *time.Time, now time.Time) bool {
	if t == nil {
		return false
	}
	a := truncateToDate(*t)
	b := truncateToDate(now)
	return a.Equal(b)
}

func (r *LevelRepoImpl) GetLeaderboard(limit int) ([]LeaderboardUser, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	var rows []LeaderboardUser

	err := r.dbpool.Select(&rows, `
		SELECT
			s.user_id,
			u.user_name,
			u.profile_images,
			s.current_streak,
			s.longest_streak
		FROM tbl_user_streaks s
		LEFT JOIN tbl_users u ON u.id = s.user_id
		ORDER BY s.longest_streak DESC, s.current_streak DESC
		LIMIT $1
	`, limit)

	if err != nil {
		return nil, msg.NewErrorResponse("database_error", err)
	}
	return rows, nil
}
