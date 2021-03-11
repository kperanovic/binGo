package socketio

import (
	"context"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kperanovic/tombola/internal/engine"
	"go.uber.org/zap"
)

type SocketIO struct {
	ctx    context.Context
	log    *zap.Logger
	engine *engine.Engine
	Server *socketio.Server
}

func NewSocketIO(ctx context.Context, log *zap.Logger, engine *engine.Engine) *SocketIO {
	return &SocketIO{
		ctx:    ctx,
		log:    log,
		engine: engine,
	}
}

func (s *SocketIO) Init() error {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return err
	}

	s.Server = server

	server.OnConnect("/", func(socket socketio.Conn) error {
		socket.SetContext(s.ctx)

		s.log.Info("client connected", zap.String("connection ID", socket.ID()))
		return nil
	})

	server.OnError("/", func(socket socketio.Conn, e error) {
		s.log.Error("error occured", zap.Error(err))
		socket.Emit("test", "testing123")
	})

	server.OnDisconnect("/", func(socket socketio.Conn, reason string) {
		s.log.Info("client disconnected", zap.String("reason", reason), zap.String("string", socket.ID()))
	})

	return nil
}
