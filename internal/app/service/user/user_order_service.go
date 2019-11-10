package user

import "music/internal/pkg/modules"

type OrderService interface {
	Detail(user modules.User)
}
