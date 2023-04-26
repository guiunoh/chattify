package gin

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"thumbnailer/pkg/shutdown"
	"time"

	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	return gin.Default()
}

func Serve(router *gin.Engine, port int) {
	const idleTimeout = 10 * time.Second

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	go func(s *http.Server) {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}(srv)
	log.Println("listen:", port)

	shutdown.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), idleTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("forced to shutdown: ", err)
	}

	log.Println("exiting")
}
