package rpcserver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"time"

	grpcutils "github.com/dvasilas/proteus/internal/grpc"
	"github.com/dvasilas/proteus/internal/libqpu"
	"github.com/dvasilas/proteus/internal/proto/qpu"
	"github.com/dvasilas/proteus/internal/proto/qpu_api"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/reflection"
)

// This package implements the low-level API server functionality of the QPU
// using a grpc server.

// Server represents the QPU's API server.
type Server struct {
	port       string
	tracing    bool
	grpcServer grpcutils.GrpcServer
	api        libqpu.APIProcessor
	// temporarily here
	state libqpu.QPUState
}

// NewServer initializes a grpc server.
func NewServer(port string, tracing bool, api libqpu.APIProcessor, state libqpu.QPUState) (*Server, error) {
	grpcServer, err := grpcutils.NewServer(tracing)
	if err != nil {
		return nil, err
	}

	server := Server{
		port:       port,
		tracing:    tracing,
		grpcServer: grpcServer,
		api:        api,
		state:      state,
	}

	qpu_api.RegisterQPUAPIServer(grpcServer.Server, &server)
	reflection.Register(grpcServer.Server)

	return &server, nil
}

// Serve starts the server and listents for connections.
// It returns only when an error occurs.
func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}

	fmt.Println("QPU service starting")
	return s.grpcServer.Server.Serve(lis)
}

// Query implements the QPU's low level Query API.
func (s *Server) Query(stream qpu_api.QPUAPI_QueryServer) error {
	requestRecord, err := stream.Recv()
	if err == io.EOF {
		return libqpu.Error(errors.New("Query:stream.Recv EOF"))
	}
	if err != nil {
		return err
	}

	switch requestRecord.GetRequest().(type) {
	case *qpu_api.RequestStreamRecord_QueryRequest:
		return s.api.Query(
			libqpu.QueryRequest{Req: requestRecord.GetQueryRequest()},
			libqpu.RequestStream{Stream: stream},
		)
	default:
		return libqpu.Error(errors.New("Query expects RequestStream_Request"))
	}
}

// QueryUnary ...
func (s *Server) QueryUnary(ctx context.Context, req *qpu_api.QueryRequest) (*qpu_api.QueryResponse, error) {
	var querySp opentracing.Span
	querySp = nil
	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		pCtx := parent.Context()
		if tracer := opentracing.GlobalTracer(); tracer != nil {
			querySp = tracer.StartSpan("qpu/query", opentracing.ChildOf(pCtx))
			defer querySp.Finish()
		}
	}

	resp, err := s.api.QueryUnary(libqpu.QueryRequest{Req: req}, querySp)

	results := make([]*qpu.LogOperation, len(resp))
	for i, entry := range resp {
		results[i] = entry.Op
	}

	return &qpu_api.QueryResponse{
		Results: results,
	}, err
}

// QueryNoOp ...
func (s *Server) QueryNoOp(ctx context.Context, req *qpu_api.NoOpReq) (*qpu_api.NoOpResp, error) {

	stateCh, err := s.state.Scan("stateTableJoin", []string{"title", "description", "short_id", "user_id", "vote_sum"}, 5, nil)
	if err != nil {
		return nil, err
	}

	for record := range stateCh {
		record["title"] = ""
	}

	return &qpu_api.NoOpResp{Str: "hi "}, nil
}

// QueryArgs ...
func (s *Server) QueryArgs(ctx context.Context, req *qpu_api.QueryRequest) (*qpu_api.QueryResponse, error) {
	ts, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return nil, err
	}

	return &qpu_api.QueryResponse{
		Results: []*qpu.LogOperation{
			&qpu.LogOperation{
				ObjectId:  "obj0",
				Bucket:    "table",
				Timestamp: &qpu.Vectorclock{Vc: map[string]*timestamp.Timestamp{"ts": ts}},
				Payload: &qpu.Payload{
					Val: &qpu.Payload_State{
						State: &qpu.ObjectState{Attributes: map[string]*qpu.Value{
							"attrA": &qpu.Value{Val: &qpu.Value_Int{Int: 2}},
							"attrB": &qpu.Value{Val: &qpu.Value_Int{Int: 4}},
							"attrC": &qpu.Value{Val: &qpu.Value_Int{Int: 6}},
						},
						},
					},
				},
			},
			&qpu.LogOperation{},
		},
	}, nil
}

// GetConfig implements the QPU's low level GetConfig API.
func (s *Server) GetConfig(ctx context.Context, in *qpu_api.ConfigRequest) (*qpu_api.ConfigResponse, error) {
	return s.api.GetConfig(ctx, in)
}

// GetDataTransfer implements the QPU's low level GetDataTransfer API.
// func (s *Server) GetDataTransfer(ctx context.Context, in *qpu_api.GetDataRequest) (*qpu_api.DataTransferResponse, error) {
// 	return s.api.GetDataTransfer(ctx, in)
// }
