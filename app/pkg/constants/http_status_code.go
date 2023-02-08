package constants

const (
	Continue           int = 100
	SwitchingProtocols int = 101
	Processing         int = 102
	EarlyHints         int = 103

	OK       int = 200
	Created  int = 201
	Accepted int = 202

	MultipleChoices int = 300

	MovedPermanently int = 301
	Found            int = 302
	SeeOther         int = 303
	NotModified      int = 304
	UseProxy         int = 305

	BadRequest       int = 400
	Unauthorized     int = 401
	PaymentRequired  int = 402
	Forbidden        int = 403
	NotFound         int = 404
	MethodNotAllowed int = 405

	InternalServerError     int = 500
	NotImplemented          int = 501
	BadGateway              int = 502
	ServiceUnavailable      int = 503
	GatewayTimeout          int = 504
	HttpVersionNotSupported int = 505
)
