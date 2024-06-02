package module

import (
	"disburse-app/internal/model/wallet"
	mockDB "disburse-app/internal/repo/db/_mock"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
)

func Test_usecase_Disburse(t *testing.T) {
	db, mdb, err := sqlmock.New()
	if err != nil {
		t.Error("failed mock db")
		return
	}
	defer db.Close()

	type args struct {
		userID int64
		amount float64
	}
	tests := []struct {
		name    string
		args    args
		mock    func(ctrl *gomock.Controller) *usecase
		want    wallet.Wallet
		wantErr bool
	}{
		{
			name: "success_disburse_10k_to_1k",
			args: args{
				userID: 1,
				amount: 1000000.00,
			},
			mock: func(ctrl *gomock.Controller) *usecase {
				mdb.ExpectBegin()

				mockdb := mockDB.NewMockDB(ctrl)
				mockdb.EXPECT().GetBalanceUser(gomock.Any(), gomock.Any()).Return(float64(10000000), nil)
				mockdb.EXPECT().UpdateBalanceUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

				mdb.ExpectCommit()
				return &usecase{
					rDB: mockdb,
					db:  db,
				}
			},
			want: wallet.Wallet{
				UserID:  1,
				Balance: 9000000,
			},
			wantErr: false,
		},
		{
			name: "failed_insufficient_10k_with_20k",
			args: args{
				userID: 1,
				amount: 20000000.00,
			},
			mock: func(ctrl *gomock.Controller) *usecase {
				mdb.ExpectBegin()

				mockdb := mockDB.NewMockDB(ctrl)
				mockdb.EXPECT().GetBalanceUser(gomock.Any(), gomock.Any()).Return(float64(10000000), nil)

				mdb.ExpectRollback()
				return &usecase{
					rDB: mockdb,
					db:  db,
				}
			},
			want: wallet.Wallet{
				UserID:  1,
				Balance: 10000000,
			},
			wantErr: true,
		},
		{
			name: "failed_error_get_balance_user",
			args: args{
				userID: 1,
				amount: 20000000.00,
			},
			mock: func(ctrl *gomock.Controller) *usecase {
				mdb.ExpectBegin()

				mockdb := mockDB.NewMockDB(ctrl)
				mockdb.EXPECT().GetBalanceUser(gomock.Any(), gomock.Any()).Return(float64(0), errors.New("test-err"))

				mdb.ExpectRollback()
				return &usecase{
					rDB: mockdb,
					db:  db,
				}
			},
			want: wallet.Wallet{
				UserID:  1,
				Balance: 0,
			},
			wantErr: true,
		},
		{
			name: "failed_error_update_balance_user",
			args: args{
				userID: 1,
				amount: 1000000.00,
			},
			mock: func(ctrl *gomock.Controller) *usecase {
				mdb.ExpectBegin()

				mockdb := mockDB.NewMockDB(ctrl)
				mockdb.EXPECT().GetBalanceUser(gomock.Any(), gomock.Any()).Return(float64(10000000), nil)
				mockdb.EXPECT().UpdateBalanceUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("test-err"))

				mdb.ExpectRollback()
				return &usecase{
					rDB: mockdb,
					db:  db,
				}
			},
			want: wallet.Wallet{
				UserID:  1,
				Balance: 10000000,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			u := tt.mock(ctrl)
			got, err := u.Disburse(tt.args.userID, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Disburse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Disburse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
