package payment

import (
	handler "payment/internal/payment/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(r *gin.Engine, h *handler.PaymentHandler) {
	payment := r.Group("/payment")
	payment.POST("/", h.CreatePayment)
	payment.GET("/", h.GetAllPayments)
	payment.PUT("/:id", h.UpdatePayment)
}

/*
In Go’s net/http, the request flow starts when a client (browser, curl, or another service) sends an HTTP request to the server over TCP. The Go server, through http.Server, is listening on a port and accepts this incoming TCP connection, which is then internally managed as an HTTP connection.

Once the connection is accepted, Go parses the raw HTTP bytes into a structured *http.Request object. This object contains all the important information like the HTTP method, URL path, headers, body, and context.

After parsing, the request is passed to the router, which in Go’s case is the ServeMux. The ServeMux matches only the URL path (not the HTTP method) and selects the appropriate handler registered for that route.

Finally, the matched handler function is executed with http.ResponseWriter and *http.Request. Inside the handler, we read data from the request and write the response using the response writer. Once writing is complete, Go sends the response back to the client over the same connection.

So overall, the flow can be summarized as: a client sends a request, the server accepts and parses it into a request object, the router matches the URL path to a handler, and the handler processes the request and writes the response back to the client.

*/
