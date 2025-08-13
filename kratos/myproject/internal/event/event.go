package event

// event definition rule:
// event name should be in the format of "event.<module>.<action>"

// ------------------------------------------------------------
// User event
// ------------------------------------------------------------
const (
	UserLogin Event = "event.user.login"

	UserLogout Event = "event.user.logout"

	UserRegister Event = "event.user.register"

	UserUpdate Event = "event.user.update"

	UserDelete Event = "event.user.delete"
)

// ------------------------------------------------------------
// User event
// ------------------------------------------------------------
