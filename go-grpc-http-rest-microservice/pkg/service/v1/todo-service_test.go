package v1

import (
	"context"
	"database/sql"
	v1 "github.com/Winszheng/go-grpc-http-rest-microservice/pkg/api/v1"
	"reflect"
	"testing"
)

func Test_toDoServiceServer_ReadAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		req *v1.ReadAllRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.ReadAllResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &toDoServiceServer{
				db: tt.fields.db,
			}
			got, err := s.ReadAll(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewToDoServiceServer(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want v1.ToDoServiceServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewToDoServiceServer(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewToDoServiceServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDoServiceServer_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		req *v1.CreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.CreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &toDoServiceServer{
				db: tt.fields.db,
			}
			got, err := s.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDoServiceServer_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		req *v1.DeleteRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.DeleteResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &toDoServiceServer{
				db: tt.fields.db,
			}
			got, err := s.Delete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDoServiceServer_Read(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		req *v1.ReadRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.ReadResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &toDoServiceServer{
				db: tt.fields.db,
			}
			got, err := s.Read(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDoServiceServer_ReadAll1(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		req *v1.ReadAllRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.ReadAllResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &toDoServiceServer{
				db: tt.fields.db,
			}
			got, err := s.ReadAll(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDoServiceServer_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		req *v1.UpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.UpdateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &toDoServiceServer{
				db: tt.fields.db,
			}
			got, err := s.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDoServiceServer_checkAPI(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		api string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &toDoServiceServer{
				db: tt.fields.db,
			}
			if err := s.checkAPI(tt.args.api); (err != nil) != tt.wantErr {
				t.Errorf("checkAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_toDoServiceServer_connect(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *sql.Conn
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &toDoServiceServer{
				db: tt.fields.db,
			}
			got, err := s.connect(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() got = %v, want %v", got, tt.want)
			}
		})
	}
}