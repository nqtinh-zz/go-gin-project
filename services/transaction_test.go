package services

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nqtinh/go-gin-project/models"
	"github.com/nqtinh/go-gin-project/repositories"
	"github.com/nqtinh/go-gin-project/repositories/mock"
)

func Test_transactionService_CreateUserTransaction(t *testing.T) {
	ctx := context.TODO()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	createUserTransactionReq := &models.CreateUserTransactionReq{
		UserID:          1,
		AccountID:       1,
		Amount:          12345,
		TransactionType: models.Deposit,
	}
	type fields struct {
		repo *repositories.Repository
	}
	type args struct {
		ctx context.Context
		req *models.CreateUserTransactionReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.UserTransactionResp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "CreateUserTransaction - Valid",
			fields: fields{
				repo: &repositories.Repository{
					TransactionRepository: func() repositories.TransactionRepository {
						mockTransactionRepository := mock.NewMockTransactionRepository(mockCtrl)
						mockTransactionRepository.EXPECT().CreateUserTransaction(ctx, createUserTransactionReq).Return(
							1, nil,
						)
						mockTransactionRepository.EXPECT().GetUserTransaction(ctx, 1).Return(
							nil, nil,
						)
						return mockTransactionRepository
					}(),
				},
			},
			args: args{
				ctx: ctx,
				req: createUserTransactionReq,
			},
		},
		{
			name: "CreateUserTransaction - Invalid",
			fields: fields{
				repo: &repositories.Repository{
					TransactionRepository: func() repositories.TransactionRepository {
						mockTransactionRepository := mock.NewMockTransactionRepository(mockCtrl)
						mockTransactionRepository.EXPECT().CreateUserTransaction(ctx, createUserTransactionReq).Return(
							1, errors.New("error"),
						)
						return mockTransactionRepository
					}(),
				},
			},
			args: args{
				ctx: ctx,
				req: createUserTransactionReq,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &transactionService{
				repo: tt.fields.repo,
			}
			got, err := s.CreateUserTransaction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactionService.CreateUserTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transactionService.CreateUserTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionService_GetUserTransactions(t *testing.T) {
	ctx := context.TODO()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	getUserTransactionsReq := &models.GetUserTransactionsReq{
		UserID: 1,
	}
	type fields struct {
		repo *repositories.Repository
	}
	type args struct {
		ctx context.Context
		req *models.GetUserTransactionsReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.UserTransactionResp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GetUserTransactions - Valid",
			fields: fields{
				repo: &repositories.Repository{
					TransactionRepository: func() repositories.TransactionRepository {
						mockTransactionRepository := mock.NewMockTransactionRepository(mockCtrl)
						mockTransactionRepository.EXPECT().GetUserTransactions(ctx, getUserTransactionsReq).Return(
							nil, nil,
						)
						return mockTransactionRepository
					}(),
				},
			},
			args: args{
				ctx: ctx,
				req: getUserTransactionsReq,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &transactionService{
				repo: tt.fields.repo,
			}
			got, err := s.GetUserTransactions(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactionService.GetUserTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transactionService.GetUserTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionService_UpdateUserTransactions(t *testing.T) {
	ctx := context.TODO()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	updateUserTransactionsReq := &models.UpdateUserTransactionsReq{
		UserID:          1,
		Amount:          12345,
		TransactionType: models.Deposit,
	}
	type fields struct {
		repo *repositories.Repository
	}
	type args struct {
		ctx context.Context
		req *models.UpdateUserTransactionsReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.UserTransactionResp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "UpdateUserTransactions - Valid",
			fields: fields{
				repo: &repositories.Repository{
					TransactionRepository: func() repositories.TransactionRepository {
						mockTransactionRepository := mock.NewMockTransactionRepository(mockCtrl)
						mockTransactionRepository.EXPECT().UpdateUserTransactions(ctx, updateUserTransactionsReq).Return(
							nil,
						)
						mockTransactionRepository.EXPECT().GetUserTransactions(ctx, &models.GetUserTransactionsReq{
							UserID: updateUserTransactionsReq.UserID,
						}).Return(
							nil, nil,
						)
						return mockTransactionRepository
					}(),
				},
			},
			args: args{
				ctx: ctx,
				req: updateUserTransactionsReq,
			},
		},
		{
			name: "UpdateUserTransactions - Invalid",
			fields: fields{
				repo: &repositories.Repository{
					TransactionRepository: func() repositories.TransactionRepository {
						mockTransactionRepository := mock.NewMockTransactionRepository(mockCtrl)
						mockTransactionRepository.EXPECT().UpdateUserTransactions(ctx, updateUserTransactionsReq).Return(
							errors.New("error"),
						)
						return mockTransactionRepository
					}(),
				},
			},
			args: args{
				ctx: ctx,
				req: updateUserTransactionsReq,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &transactionService{
				repo: tt.fields.repo,
			}
			got, err := s.UpdateUserTransactions(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactionService.UpdateUserTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transactionService.UpdateUserTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}
