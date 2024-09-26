package errs

import "errors"

// Ошибки аутентификации
var (
	ErrEmptyAuthHeader             = errors.New("empty auth header")              // Пустой заголовок аутентификации
	ErrInvalidAuthHeader           = errors.New("invalid auth header")            // Неверный заголовок аутентификации
	ErrTokenParsingFailed          = errors.New("token parsing failed")           // Ошибка обработки токена
	ErrUserNotAuthenticated        = errors.New("user not authenticated")         // Пользователь не аутентифицирован
	ErrIncorrectUsernameOrPassword = errors.New("incorrect username or password") // Неверное имя пользователя или пароль
)

// Ошибки разрешений
var (
	ErrPermissionDenied                  = errors.New("access denied")                    // Доступ запрещен
	ErrPermissionDeniedOnlyForAdmin      = errors.New("access denied: admins only")       // Доступ только для администраторов
	ErrPermissionDeniedOnlyForSuperAdmin = errors.New("access denied: super admins only") // Доступ только для супер администраторов
)

// Ошибки пользователей
var (
	ErrUserNotFound             = errors.New("user not found")          // Пользователь не найден
	ErrUsersNotFound            = errors.New("users not found")         // Пользователи не найдены
	ErrUsernameUniquenessFailed = errors.New("username must be unique") // Имя пользователя должно быть уникальным
	ErrRecordNotFound           = errors.New("record not found")        // Запись не найдена
	ErrUserAlreadyDeleted       = errors.New("user already deleted")    // Пользователь уже удалён
	ErrUserNotDeleted           = errors.New("user not deleted")        // Пользователь не был удалён
	ErrSomethingWentWrong       = errors.New("something went wrong")    // Что-то пошло не так
)

// Ошибки категорий
var (
	ErrCategoryNameExists = errors.New("category name already exists") // Категория с таким именем уже существует
	ErrCategoryNotFound   = errors.New("category not found")           // Категория не найдена
)

// Ошибки компьютеров
var (
	ErrComputerNotFound      = errors.New("computer not found")      // Компьютер не найден
	ErrComputerAlreadyExists = errors.New("computer already exists") // Компьютер уже существует
)

// Ошибки бронирования
var (
	ErrBookingNotFound         = errors.New("booking not found")         // Бронирование не найдено
	ErrBookingAlreadyCompleted = errors.New("booking already completed") // Бронирование уже завершено
)

// Ошибки баланса
var (
	ErrBalanceInsufficient = errors.New("insufficient balance") // Недостаточно средств
	ErrBalanceNotFound     = errors.New("balance not found")    // Баланс не найден
)

// Общие ошибки
var (
	ErrUnauthorized = errors.New("unauthorized access") // Неавторизованный доступ
	ErrServerError  = errors.New("server error")        // Ошибка сервера
)

var (
	ErrValidationFailed = errors.New("validation failed") // Ошибка: валидация данных не пройдена...
	//ErrUsernameUniquenessFailed    = errors.New("username must be unique")        // Ошибка: имя пользователя должно быть уникальным...
	//ErrIncorrectUsernameOrPassword = errors.New("incorrect username or password") // Ошибка: неверное имя пользователя или пароль...
)
