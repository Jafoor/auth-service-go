package user

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}
