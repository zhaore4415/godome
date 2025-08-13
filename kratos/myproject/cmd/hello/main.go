package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"myproject/internal/asset"
	"myproject/internal/micro/app"
	"myproject/internal/micro/config"
	"myproject/internal/micro/otel"
	"myproject/internal/proto/hello"
	"bsi/kratos/micro/auth"
	"bsi/kratos/micro/errors"
	"bsi/kratos/micro/linker"
	"bsi/kratos/micro/server"

	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	registryDC      = ""
	registryAddress = ""
	registryToken   = ""
	configPath      = ""
	localIP         = ""

	Name    = "bsi.hello"
	Path    = "bsi/hello"
	Version = "v0.0.1"
)

func init() {
	flag.StringVar(&registryDC, "registry_dc", "", "consul data center")
	flag.StringVar(&registryAddress, "registry_address", "https://consul.dev.shijizhongyun.com", "consul address")
	flag.StringVar(&registryToken, "registry_token", "3f84201a-31a2-c843-bbc0-0a45983aa7b7", "consul acl token")
	flag.StringVar(&configPath, "config_path", Path, "consul config path")
	flag.StringVar(&localIP, "local_ip", "", "local ip address register to consul. just only ip address. do not include port")
}

func newApp(logger log.Logger, bs *hello.Bootstrap, gs *grpc.Server, hs *http.Server, r registry.Registrar, d registry.Discovery) *kratos.App {
	if len(bs.Service.Name) == 0 {
		bs.Service.Name = Name
	}
	if len(bs.Service.Version) == 0 {
		bs.Service.Version = Version
	}
	servers := []transport.Server{gs}
	if hs != nil {
		servers = append(servers, hs)
	}
	return kratos.New(
		kratos.ID(server.Id()),
		kratos.Name(bs.Service.Name),
		kratos.Version(bs.Service.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Registrar(r),
		kratos.Server(servers...),
		kratos.BeforeStart(func(ctx context.Context) error {
			if err := auth.Init(&auth.Options{SymmetricKey: bs.AuthKey, PublicKey: bs.PublicKey}); err != nil {
				return err
			}
			if err := linker.Init(d); err != nil {
				return err
			}
			if err := errors.Init(d); err != nil {
				return err
			}
			return otel.SetupOTel(ctx, bs)
		}),
		kratos.AfterStart(func(ctx context.Context) error {
			app.OnStart(ctx)
			return auth.InitAsset(bs.Service.Product, bs.Service.Name, asset.GrpcAsset, d)
		}),
		kratos.AfterStop(func(ctx context.Context) error {
			app.OnStop(ctx)
			return nil
		}),
	)
}

func main() {
	flag.Parse()

	writeSyncer := zapcore.AddSync(os.Stdout)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	z := zap.New(core)

	logger := log.With(kratoszap.NewLogger(z),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	log.SetLogger(logger)

	bs, err := config.LoadConfig(configPath, registryDC, registryAddress, registryToken)
	if err != nil {
		panic(err)
	}

	if localIP != "" {
		if ip := net.ParseIP(localIP); ip == nil {
			panic(fmt.Sprintf("invalid local address: %s", localIP))
		}

		_, port, err := net.SplitHostPort(bs.Server.Grpc.Addr)
		if err != nil {
			panic(fmt.Sprintf("invalid local address: %s", localIP))
		}

		bs.Server.Grpc.Addr = fmt.Sprintf("%s:%s", localIP, port)

		if bs.Server.Http != nil && bs.Server.Http.Addr != "" {
			_, port, err := net.SplitHostPort(bs.Server.Http.Addr)
			if err != nil {
				panic(fmt.Sprintf("invalid local address: %s", localIP))
			}
			bs.Server.Http.Addr = fmt.Sprintf("%s:%s", localIP, port)
		}
	}

	app, cleanup, err := wireApp(logger, bs)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
