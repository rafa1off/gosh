package server

import (
	"errors"
	"fmt"
	"gosh/pkg/term"
	"net"
	"os"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
)

func Run(host, port string) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Error("Could not start server", "error", err)
	}
	srv, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(fmt.Sprintf("%s/.ssh/gosh_key", home)),
		wish.WithMiddleware(

			bubbletea.Middleware(term.Serve),

			func(next ssh.Handler) ssh.Handler {
				return func(sess ssh.Session) {
					wish.Println(sess, "Have a tea :D!")
					next(sess)
				}
			},

			logging.Middleware(),
		),
	)

	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	log.Info("Starting SSH server from air :D", "host", host, "port", port)
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not start server", "error", err)
	}
}
