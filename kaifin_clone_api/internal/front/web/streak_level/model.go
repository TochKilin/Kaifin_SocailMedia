package streaklevel

import "time"

type LevelInfo struct {
	ID         int64  `json:"id" db:"id"`
	Level      int    `json:"level" db:"level"`
	RequiredXP int    `json:"required_xp" db:"required_xp"`
	Title      string `json:"title" db:"title"`
	BadgeIcon  string `json:"badge_icon" db:"badge_icon"`
}

type UserStreak struct {
	ID             int64      `json:"id" db:"id"`
	UserID         int64      `json:"user_id" db:"user_id"`
	CurrentStreak  int        `json:"current_streak" db:"current_streak"`
	LongestStreak  int        `json:"longest_streak" db:"longest_streak"`
	LastActiveDate *time.Time `json:"last_active_date" db:"last_active_date"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}

type UserLevelStatus struct {
	CurrentStreak  int        `json:"current_streak"`
	LongestStreak  int        `json:"longest_streak"`
	LastActiveDate *time.Time `json:"last_active_date"`
	XP             int        `json:"xp"`
	CheckedInToday bool       `json:"checked_in_today"`
	CurrentLevel   *LevelInfo `json:"current_level"`
	NextLevel      *LevelInfo `json:"next_level,omitempty"`
	XPToNextLevel  *int       `json:"xp_to_next_level,omitempty"`
}

type LeaderboardUser struct {
	UserID        int64   `json:"user_id" db:"user_id"`
	Username      string  `json:"user_name" db:"user_name"`
	ProfileImages *string `json:"profile_images" db:"profile_images"`
	CurrentStreak int     `json:"current_streak" db:"current_streak"`
	LongestStreak int     `json:"longest_streak" db:"longest_streak"`
}
