package repository

import (
	"sync"
	"time"
)

type otpEntry struct {
    Code      string
    ExpiresAt time.Time
}

type OTPRepository interface {
    SaveOTP(phone, code string) error
    ValidateOTP(phone, code string) bool
    CanRequestOTP(phone string) bool
}

type otpRepo struct {
    otps     map[string]otpEntry
    requests map[string][]time.Time
    mu       sync.Mutex
}

func NewOTPRepository() OTPRepository {
    return &otpRepo{
        otps:     make(map[string]otpEntry),
        requests: make(map[string][]time.Time),
    }
}

func (r *otpRepo) SaveOTP(phone, code string) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    r.otps[phone] = otpEntry{
        Code:      code,
        ExpiresAt: time.Now().Add(1 * time.Minute),
    }

    r.requests[phone] = append(r.requests[phone], time.Now())
    return nil
}

func (r *otpRepo) ValidateOTP(phone, code string) bool {
    r.mu.Lock()
    defer r.mu.Unlock()

    entry, exists := r.otps[phone]
    if !exists || time.Now().After(entry.ExpiresAt) || entry.Code != code {
        return false
    }

    delete(r.otps, phone)
    return true
}

func (r *otpRepo) CanRequestOTP(phone string) bool {
    r.mu.Lock()
    defer r.mu.Unlock()

    now := time.Now()
    var recent []time.Time
    for _, t := range r.requests[phone] {
        if now.Sub(t) < 10*time.Minute {
            recent = append(recent, t)
        }
    }

    r.requests[phone] = recent
    return len(recent) < 3
}
