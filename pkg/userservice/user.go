package userservice

type Repository interface {
	GetUser(userID string) (User, error)
	InsertUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(userID string) (User, error)
}

type Service interface {
	GetUser(userID string) (User, error)
	CreateUser(firstName, lastName string) (User, error)
	UpdateUser(userID, firstName, lastName string) (User, error)
	DeleteUser(userID string) (User, error)

	CheckPalindrome(userID string) (PalindromeResult, error)
}

type User struct {
	UserID    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type PalindromeResult struct {
	FirstName bool `json:"firstname"`
	LastName  bool `json:"lastname"`
}

type service struct {
	repo Repository
}

func New(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetUser(userID string) (User, error) {
	return s.repo.GetUser(userID)
}

func (s *service) CreateUser(firstName, lastName string) (User, error) {
	return s.repo.InsertUser(User{
		UserID:    "",
		FirstName: firstName,
		LastName:  lastName,
	})
}

func (s *service) UpdateUser(userID, firstName, lastName string) (User, error) {
	return s.repo.UpdateUser(User{
		UserID:    userID,
		FirstName: firstName,
		LastName:  lastName,
	})
}

func (s *service) DeleteUser(userID string) (User, error) {
	return s.repo.DeleteUser(userID)
}

func (s *service) CheckPalindrome(userID string) (PalindromeResult, error) {
	user, err := s.repo.GetUser(userID)
	if err != nil {
		return PalindromeResult{}, err
	}

	return PalindromeResult{
		FirstName: isPalindrome(user.FirstName),
		LastName:  isPalindrome(user.LastName),
	}, nil
}
