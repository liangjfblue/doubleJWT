package base

// RefreshTokenReq refresh token request
type RefreshTokenReq struct {
}

// RefreshTokenRes refresh token respond
type RefreshTokenRes struct {
    Token           string  `json:"token"`
    RefreshToken    string  `json:"refreshToken"`
}