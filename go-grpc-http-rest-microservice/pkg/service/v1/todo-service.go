package v1

import (
	"context"
	"database/sql"
	"fmt"
	v1 "github.com/Winszheng/go-grpc-http-rest-microservice/pkg/api/v1"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

const (
	apiVersion = "v1"
)

type toDoServiceServer struct {
	db *sql.DB
}

func NewToDoServiceServer(db *sql.DB) v1.ToDoServiceServer{
	return &toDoServiceServer{db:db}
}

func (s *toDoServiceServer) checkAPI (api string) error{
	//api is "" means use current version of the service
	if len(api)>0{
		if apiVersion != api{
			return status.Errorf(codes.Unimplemented, "不支持的api版本")
		}
	}
	return nil
}

func (s *toDoServiceServer) connect(ctx context.Context) (*sql.Conn, error){
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database: "+err.Error())
	}
	return c, nil
}

func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error){
	//check if the api version req by client is supported by server
	if err:= s.checkAPI(req.Api); err!= nil {
		return nil, err
	}
	//get sql conn from pool
	c, err := s.connect(ctx)
	if err!= nil {
		return nil, err
	}
	defer c.Close()

	//ptypes.Timestamp：把protobuf中的时间戳(指针)转换成golang中的时间戳，比如转换成：
	//2020-04-26 09:54:53.853238 +0000 UTC
	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format: "+err.Error())
	}

	//insert toDo entity data 添加
	res, err := c.ExecContext(ctx, "insert into ToDo(Title, Description, Reminder) values(?, ?, ?)", req.ToDo.Title, req.ToDo.Description, reminder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ToDo-> "+err.Error())
	}

	//get id of created todo
	id ,err := res.LastInsertId()
	if err!= nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ToDo-> "+err.Error())
	}

	return &v1.CreateResponse{Api:apiVersion, Id:id,}, nil
}

func (s *toDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error){
	if err:=s.checkAPI(req.Api); err!= nil {
		return nil, err
	}

	c, err := s.connect(ctx)
	if err!= nil {
		return nil, err
	}

	//query todo by id 查询
	row, err := c.QueryContext(ctx, "select ID, Title, Description, Reminder, from ToDo where ID=?",req.Id)
	if err!= nil {
		return nil, status.Error(codes.Unknown, "failed to select from db: "+err.Error())
	}
	if !row.Next() {
		if err:=row.Err();err!= nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from db: "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("todo task with ID=%d is not found", req.Id))
	}

	var td v1.ToDo
	var reminder time.Time
	if err:= row.Scan(&td.Id, &td.Title, &td.Description, &reminder); err!= nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from row: "+err.Error())
	}
	td.Reminder, err = ptypes.TimestampProto(reminder)
	if err!= nil {
		return nil, status.Error(codes.Unknown, "reminder field has invalid format: "+err.Error())
	}
	if row.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple todo tasks with ID = %d", req.Id))
	}

	return &v1.ReadResponse{
		Api:  apiVersion,
		ToDo: &td,
	}, nil
}

func(s *toDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err:= s.checkAPI(req.Api); err!= nil {
		return nil, err
	}
	c, err := s.connect(ctx)
	if err!= nil {
		return nil, err
	}

	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
	if err!= nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format: "+err.Error())
	}

	res, err := c.ExecContext(ctx, "update ToDo set Title=?, Description=?, Reminder=? where ID=?", req.ToDo.Title, req.ToDo.Description, reminder,req.ToDo.Id)
	if err!= nil {
		return nil, status.Error(codes.Unknown, "failed to update ToDo: "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err!= nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value: "+err.Error())
	}
	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("update id = %d error", req.ToDo.Id))
	}

	return &v1.UpdateResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

func (s *toDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error)  {
	if err:= s.checkAPI(req.Api); err!= nil {
		return nil, err
	}

	c, err := s.connect(ctx)
	if err!= nil {
		return nil, err
	}
	defer c.Close()

	//delete todo task
	res, err := c.ExecContext(ctx, "delete from ToDo where id = ?", req.Id)
	if err!= nil {
		return nil, status.Error(codes.Unknown, "failed to delete todo task: "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err!= nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value: "+ err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("todo with id = %d is not found", req.Id))
	}

	return &v1.DeleteResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}





func (s *toDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	if err := s.checkAPI(req.Api); err!= nil {
		return nil, err
	}

	c, err := s.connect(ctx)
	if err!= nil {
		return nil, err
	}
	defer c.Close()

	rows, err := c.QueryContext(ctx, "select ID, Title, Description, Reminder from ToDo")
	if err!= nil {
		return nil, status.Error(codes.Unknown, "ReadAll failed: "+err.Error())
	}
	defer rows.Close()

	var reminder time.Time
	list := []*v1.ToDo{}
	for rows.Next() {
		td := new(v1.ToDo)
		if err := rows.Scan(&td.Id, &td.Title, &td.Description, &reminder); err!= nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve values from db: "+err.Error())
		}
		td.Reminder, err = ptypes.TimestampProto(reminder)
		if err != nil {
			return nil, status.Error(codes.Unknown, "reminder field has invalid format: "+err.Error())
		}
		list = append(list, td)
	}

	if err:= rows.Err(); err!= nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve values from db: "+err.Error())
	}

	return &v1.ReadAllResponse{
		Api:   apiVersion,
		ToDos: list,
	}, nil
}