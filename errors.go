package framework

import "fmt"

// InputError error de input de services
type InputError struct {
	Message string
}

func (i *InputError) Error() string {
	return fmt.Sprintf("Input inválido: %v", i.Message)
}

// GatewayError error de service externo
type GatewayError struct {
	Message string
	Cause   error
}

func (g *GatewayError) Error() string {
	return fmt.Sprintf("Error de Gateway : %v", g.Message)
}

//NotFoundError errores codificados de backend
type NotFoundError struct {
	Code    string
	Message string
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("Error NotFound : %v", n.Message)
}

//CustomError errores customizables
type CustomError struct {
	Code       string
	Message    string
	StatusCode int
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("Error: %v", c.Message)
}

//BackendCodedError errores codificados de backend
type BackendCodedError struct {
	Code    string
	Message string
}

func (g *BackendCodedError) Error() string {
	return fmt.Sprintf("Error. Backend code: %q, message: %q", g.Code, g.Message)
}
