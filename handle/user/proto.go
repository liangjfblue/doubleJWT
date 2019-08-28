package user

type RegisterReq struct {
    Username    string  `json:"username"`
    Password    string  `json:"password"`
}

type RegisterRes struct {
    UUID            string  `json:"uuid"`
}

type LoginReq struct {
    Username    string  `json:"username"`
    Password    string  `json:"password"`
}

type LoginRes struct {
    Token           string  `json:"token"`
    RefreshToken    string  `json:"refreshToken"`
}
