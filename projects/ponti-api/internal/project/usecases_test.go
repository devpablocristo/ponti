package project

import (
	"context"
	"errors"
	"fmt"
	"testing"

	cropdom "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/crop/usecases/domain"
	customer "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/customer/mocks"
	customerdom "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/customer/usecases/domain"
	field "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/field/mocks"
	fielddom "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/field/usecases/domain"
	investor "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/investor/mocks"
	investordom "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/investor/usecases/domain"
	lot "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/lot/mocks"
	lotdom "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/lot/usecases/domain"
	manager "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/manager/mocks"
	managerdom "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/manager/usecases/domain"
	"github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/project/mocks"
	"github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/project/usecases/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type fields struct {
	repo *mocks.MockRepository
	cu   *customer.MockUseCases
	ma   *manager.MockUseCases
	in   *investor.MockUseCases
	fu   *field.MockUseCases
	lo   *lot.MockUseCases
	uc   UseCases
}

func setupMocks(ctrl *gomock.Controller) fields {
	repo := mocks.NewMockRepository(ctrl)
	cu := customer.NewMockUseCases(ctrl)
	ma := manager.NewMockUseCases(ctrl)
	in := investor.NewMockUseCases(ctrl)
	fu := field.NewMockUseCases(ctrl)
	lo := lot.NewMockUseCases(ctrl)

	return fields{
		repo: repo,
		cu:   cu,
		ma:   ma,
		in:   in,
		fu:   fu,
		lo:   lo,
		uc:   NewUseCases(repo, cu, ma, in, fu, lo),
	}
}

func TestCreateProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		p   *domain.Project
	}

	// Base project without IDs to force creation of all related entities
	base := &domain.Project{
		Name: "Project X",
		Customer: customerdom.Customer{
			Name: "Client A",
		},
		Managers: []managerdom.Manager{
			{Name: "Manager A"},
		},
		Investors: []investordom.Investor{
			{Name: "Investor A", Percentage: 25},
		},
		Fields: []fielddom.Field{
			{
				Name:        "Field A",
				LeaseTypeID: 1,
				Lots: []lotdom.Lot{
					{
						Name:         "Lot A",
						Hectares:     10,
						PreviousCrop: cropdom.Crop{ID: 1},
						CurrentCrop:  cropdom.Crop{ID: 2},
						Season:       "2025",
					},
				},
			},
		},
	}

	tests := []struct {
		name    string
		setup   func(f *fields)
		args    args
		wantID  int64
		wantErr bool
	}{
		{
			name: "repo success",
			setup: func(f *fields) {
				// Create customer
				f.cu.EXPECT().
					CreateCustomer(gomock.Any(), &customerdom.Customer{Name: "Client A"}).
					Return(int64(10), nil)
				// Create manager
				f.ma.EXPECT().
					CreateManager(gomock.Any(), &managerdom.Manager{Name: "Manager A"}).
					Return(int64(20), nil)
				// Create investor
				f.in.EXPECT().
					CreateInvestor(gomock.Any(), &investordom.Investor{Name: "Investor A", Percentage: 25}).
					Return(int64(30), nil)
				// Create field (handles nested lots within the Field service)
				f.fu.EXPECT().
					CreateField(gomock.Any(), &fielddom.Field{
						Name:        "Field A",
						LeaseTypeID: 1,
						Lots: []lotdom.Lot{{
							Name:         "Lot A",
							Hectares:     10,
							PreviousCrop: cropdom.Crop{ID: 1},
							CurrentCrop:  cropdom.Crop{ID: 2},
							Season:       "2025",
						}},
					}).
					Return(int64(40), nil)
				// Persist project
				f.repo.EXPECT().
					CreateProject(gomock.Any(), gomock.Any()).
					Return(int64(99), nil)
			},
			args:   args{ctx: context.Background(), p: base},
			wantID: 99,
		},
		{
			name: "repo error",
			setup: func(f *fields) {
				// In this scenario all IDs are non-zero so only CreateField + CreateProject are called
				f.fu.EXPECT().
					CreateField(gomock.Any(), gomock.Any()).
					Return(int64(40), nil)
				f.repo.EXPECT().
					CreateProject(gomock.Any(), gomock.Any()).
					Return(int64(0), errors.New("repo error"))
				// Rollback: delete project, field and customer
				f.repo.EXPECT().
					DeleteProject(gomock.Any(), int64(0)).
					Return(nil)
				f.fu.EXPECT().
					DeleteField(gomock.Any(), int64(40)).
					Return(nil)
				f.cu.EXPECT().
					DeleteCustomer(gomock.Any(), int64(10)).
					Return(nil)
			},
			args: args{ctx: context.Background(), p: &domain.Project{
				ID:   99,
				Name: "Project X",
				Customer: customerdom.Customer{
					ID:   10,
					Name: "Client A",
				},
				Managers: []managerdom.Manager{
					{ID: 20, Name: "Manager A"},
				},
				Investors: []investordom.Investor{
					{ID: 30, Name: "Investor A", Percentage: 25},
				},
				Fields: []fielddom.Field{{
					ID:          40,
					Name:        "Field A",
					LeaseTypeID: 1,
					Lots: []lotdom.Lot{{
						ID:           50,
						Name:         "Lot A",
						Hectares:     10,
						PreviousCrop: cropdom.Crop{ID: 1},
						CurrentCrop:  cropdom.Crop{ID: 2},
						Season:       "2025",
					}},
				}},
			}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := setupMocks(ctrl)
			tt.setup(&f)

			gotID, err := f.uc.CreateProject(tt.args.ctx, tt.args.p)
			if tt.wantErr {
				assert.Error(t, err, "expected error from CreateProject")
			} else {
				assert.NoError(t, err, "unexpected error from CreateProject")
				assert.Equal(t, tt.wantID, gotID, "CreateProject: returned ID mismatch")
			}
		})
	}
}

func TestGetProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		setup   func(f *fields)
		args    args
		want    *domain.Project
		wantErr bool
	}{
		{
			name: "success + enrich",
			setup: func(f *fields) {
				f.repo.EXPECT().
					GetProject(gomock.Any(), int64(42)).
					Return(&domain.Project{
						ID:        42,
						Name:      "P1",
						Customer:  customerdom.Customer{ID: 10},
						Managers:  []managerdom.Manager{{ID: 20}},
						Investors: []investordom.Investor{{ID: 30}},
						Fields:    []fielddom.Field{{ID: 40}},
					}, nil)
				// enrichProject()
				f.cu.EXPECT().GetCustomer(gomock.Any(), int64(10)).
					Return(&customerdom.Customer{ID: 10, Name: "C1"}, nil)
				f.ma.EXPECT().GetManager(gomock.Any(), int64(20)).
					Return(&managerdom.Manager{ID: 20, Name: "M1"}, nil)
				f.in.EXPECT().GetInvestor(gomock.Any(), int64(30)).
					Return(&investordom.Investor{ID: 30, Name: "I1"}, nil)
				f.fu.EXPECT().GetField(gomock.Any(), int64(40)).
					Return(&fielddom.Field{ID: 40, Name: "F1"}, nil)
			},
			args: args{ctx: context.Background(), id: 42},
			want: &domain.Project{
				ID:        42,
				Name:      "P1",
				Customer:  customerdom.Customer{ID: 10, Name: "C1"},
				Managers:  []managerdom.Manager{{ID: 20, Name: "M1"}},
				Investors: []investordom.Investor{{ID: 30, Name: "I1"}},
				Fields:    []fielddom.Field{{ID: 40, Name: "F1"}},
			},
		},
		{
			name: "repo error",
			setup: func(f *fields) {
				f.repo.EXPECT().
					GetProject(gomock.Any(), int64(99)).
					Return(nil, errors.New("not found"))
			},
			args:    args{ctx: context.Background(), id: 99},
			wantErr: true,
		},
		{
			name: "enrich error",
			setup: func(f *fields) {
				f.repo.EXPECT().GetProject(gomock.Any(), int64(42)).
					Return(&domain.Project{
						ID:        42,
						Name:      "P1",
						Customer:  customerdom.Customer{ID: 10},
						Managers:  []managerdom.Manager{{ID: 20}},
						Investors: []investordom.Investor{{ID: 30}},
						Fields:    []fielddom.Field{{ID: 40}},
					}, nil)

				f.cu.EXPECT().
					GetCustomer(gomock.Any(), int64(10)).
					Return(nil, errors.New("customer fetch error"))
			},
			args:    args{ctx: context.Background(), id: 42},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := setupMocks(ctrl)
			tt.setup(&f)

			got, err := f.uc.GetProject(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				assert.Error(t, err, "expected error from GetProject")
			} else {
				assert.NoError(t, err, "unexpected error from GetProject")
				assert.Equal(t, tt.want, got, "GetProject: returned project mismatch")
			}
		})
	}
}

func TestListProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		setup   func(f *fields)
		args    args
		want    []domain.Project
		wantErr bool
	}{
		{
			name: "success + enrich all",
			setup: func(f *fields) {
				list := []domain.Project{
					{
						ID:        1,
						Name:      "P1",
						Customer:  customerdom.Customer{ID: 10},
						Managers:  []managerdom.Manager{{ID: 20}},
						Investors: []investordom.Investor{{ID: 30}},
						Fields:    []fielddom.Field{{ID: 40}},
					},
					{
						ID:        2,
						Name:      "P2",
						Customer:  customerdom.Customer{ID: 11},
						Managers:  []managerdom.Manager{{ID: 21}},
						Investors: []investordom.Investor{{ID: 31}},
						Fields:    []fielddom.Field{{ID: 41}},
					},
				}
				f.repo.EXPECT().
					ListProjects(gomock.Any()).
					Return(list, nil)

				// enrich All
				for _, p := range list {
					f.cu.EXPECT().GetCustomer(gomock.Any(), p.Customer.ID).
						Return(&customerdom.Customer{ID: p.Customer.ID, Name: fmt.Sprintf("C%d", p.Customer.ID)}, nil)
					f.ma.EXPECT().GetManager(gomock.Any(), p.Managers[0].ID).
						Return(&managerdom.Manager{ID: p.Managers[0].ID, Name: fmt.Sprintf("M%d", p.Managers[0].ID)}, nil)
					f.in.EXPECT().GetInvestor(gomock.Any(), p.Investors[0].ID).
						Return(&investordom.Investor{ID: p.Investors[0].ID, Name: fmt.Sprintf("I%d", p.Investors[0].ID)}, nil)
					f.fu.EXPECT().GetField(gomock.Any(), p.Fields[0].ID).
						Return(&fielddom.Field{ID: p.Fields[0].ID, Name: fmt.Sprintf("F%d", p.Fields[0].ID)}, nil)
				}
			},
			args: args{ctx: context.Background()},
			want: []domain.Project{
				{ID: 1, Name: "P1",
					Customer:  customerdom.Customer{ID: 10, Name: "C10"},
					Managers:  []managerdom.Manager{{ID: 20, Name: "M20"}},
					Investors: []investordom.Investor{{ID: 30, Name: "I30"}},
					Fields:    []fielddom.Field{{ID: 40, Name: "F40"}}},
				{ID: 2, Name: "P2",
					Customer:  customerdom.Customer{ID: 11, Name: "C11"},
					Managers:  []managerdom.Manager{{ID: 21, Name: "M21"}},
					Investors: []investordom.Investor{{ID: 31, Name: "I31"}},
					Fields:    []fielddom.Field{{ID: 41, Name: "F41"}}},
			},
		},
		{
			name: "repo error",
			setup: func(f *fields) {
				f.repo.EXPECT().
					ListProjects(gomock.Any()).
					Return(nil, errors.New("not found"))
			},
			args:    args{ctx: context.Background()},
			wantErr: true,
		},
		{
			name: "enrich error - GetInvestor fails",
			setup: func(f *fields) {
				list := []domain.Project{
					{
						ID:        1,
						Name:      "P1",
						Customer:  customerdom.Customer{ID: 10},
						Managers:  []managerdom.Manager{{ID: 20}},
						Investors: []investordom.Investor{{ID: 30}},
						Fields:    []fielddom.Field{{ID: 40}},
					},
				}
				f.repo.EXPECT().
					ListProjects(gomock.Any()).
					Return(list, nil)

				f.cu.EXPECT().GetCustomer(gomock.Any(), int64(10)).
					Return(&customerdom.Customer{ID: 10, Name: "C10"}, nil)
				f.ma.EXPECT().GetManager(gomock.Any(), int64(20)).
					Return(&managerdom.Manager{ID: 20, Name: "M20"}, nil)
				f.in.EXPECT().GetInvestor(gomock.Any(), int64(30)).
					Return(nil, errors.New("investor fetch fail"))
				// GetField is not called due to the previous error
			},
			args:    args{ctx: context.Background()},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := setupMocks(ctrl)
			tt.setup(&f)

			got, err := f.uc.ListProjects(tt.args.ctx)
			if tt.wantErr {
				assert.Error(t, err, "expected error from ListProjects")
			} else {
				assert.NoError(t, err, "unexpected error from ListProjects")
				assert.Equal(t, tt.want, got, "ListProjects: returned project list mismatch")
			}
		})

	}
}

func TestListProjectsByCustomerID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx        context.Context
		customerID int64
	}

	tests := []struct {
		name    string
		setup   func(f *fields)
		args    args
		want    []domain.Project
		wantErr bool
	}{
		{
			name: "success + enrich",
			setup: func(f *fields) {
				list := []domain.Project{
					{ID: 1, Name: "A",
						Customer:  customerdom.Customer{ID: 5},
						Managers:  []managerdom.Manager{{ID: 7}},
						Investors: []investordom.Investor{{ID: 9}},
						Fields:    []fielddom.Field{{ID: 11}}},
				}
				f.repo.EXPECT().
					ListProjectsByCustomerID(gomock.Any(), int64(5)).
					Return(list, nil)

				//enrich
				f.cu.EXPECT().GetCustomer(gomock.Any(), int64(5)).Return(&customerdom.Customer{ID: 5, Name: "C5"}, nil)
				f.ma.EXPECT().GetManager(gomock.Any(), int64(7)).Return(&managerdom.Manager{ID: 7, Name: "M7"}, nil)
				f.in.EXPECT().GetInvestor(gomock.Any(), int64(9)).Return(&investordom.Investor{ID: 9, Name: "I9"}, nil)
				f.fu.EXPECT().GetField(gomock.Any(), int64(11)).Return(&fielddom.Field{ID: 11, Name: "F11"}, nil)
			},
			args: args{ctx: context.Background(), customerID: 5},
			want: []domain.Project{
				{ID: 1, Name: "A",
					Customer:  customerdom.Customer{ID: 5, Name: "C5"},
					Managers:  []managerdom.Manager{{ID: 7, Name: "M7"}},
					Investors: []investordom.Investor{{ID: 9, Name: "I9"}},
					Fields:    []fielddom.Field{{ID: 11, Name: "F11"}}},
			},
		},
		{
			name: "repo error",
			setup: func(f *fields) {
				f.repo.EXPECT().ListProjectsByCustomerID(gomock.Any(), int64(99)).Return(nil, errors.New("not found"))
			},
			args:    args{ctx: context.Background(), customerID: 99},
			wantErr: true,
		},
		{
			name: "enrich error - GetField fails",
			setup: func(f *fields) {
				list := []domain.Project{
					{ID: 1, Name: "A",
						Customer:  customerdom.Customer{ID: 5},
						Managers:  []managerdom.Manager{{ID: 7}},
						Investors: []investordom.Investor{{ID: 9}},
						Fields:    []fielddom.Field{{ID: 11}}},
				}
				f.repo.EXPECT().
					ListProjectsByCustomerID(gomock.Any(), int64(5)).
					Return(list, nil)

				f.cu.EXPECT().GetCustomer(gomock.Any(), int64(5)).Return(&customerdom.Customer{ID: 5, Name: "C5"}, nil)
				f.ma.EXPECT().GetManager(gomock.Any(), int64(7)).Return(&managerdom.Manager{ID: 7, Name: "M7"}, nil)
				f.in.EXPECT().GetInvestor(gomock.Any(), int64(9)).Return(&investordom.Investor{ID: 9, Name: "I9"}, nil)
				f.fu.EXPECT().GetField(gomock.Any(), int64(11)).Return(nil, errors.New("field not found"))
			},
			args:    args{ctx: context.Background(), customerID: 5},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := setupMocks(ctrl)
			tt.setup(&f)

			got, err := f.uc.ListProjectsByCustomerID(tt.args.ctx, tt.args.customerID)
			if tt.wantErr {
				assert.Error(t, err, "expected error from ListProjectsByCustomerID")
			} else {
				assert.NoError(t, err, "unexpected error from ListProjectsByCustomerID")
				assert.Equal(t, tt.want, got, "ListProjectsByCustomerID: returned project list mismatch")
			}
		})
	}
}

func TestUpdateProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo *mocks.MockRepository
		uc   UseCases
	}
	type args struct {
		ctx context.Context
		p   *domain.Project
	}

	tests := []struct {
		name    string
		setup   func(f *fields)
		args    args
		wantErr bool
	}{
		{
			name: "success",
			setup: func(f *fields) {
				f.repo.EXPECT().
					UpdateProject(gomock.Any(), &domain.Project{ID: 1, Name: "P2"}).
					Return(nil)
			},
			args: args{ctx: context.Background(), p: &domain.Project{ID: 1, Name: "P2"}},
		},
		{
			name: "error",
			setup: func(f *fields) {
				f.repo.EXPECT().UpdateProject(gomock.Any(), gomock.Any()).
					Return(errors.New("update failed"))
			},
			args:    args{ctx: context.Background(), p: &domain.Project{ID: 1, Name: "P3"}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repoMock := mocks.NewMockRepository(ctrl)

			uc := NewUseCases(repoMock, nil, nil, nil, nil, nil)
			f := fields{repo: repoMock, uc: uc}
			tt.setup(&f)

			err := f.uc.UpdateProject(tt.args.ctx, tt.args.p)
			if tt.wantErr {
				assert.Error(t, err, "expected error from UpdateProject")
			} else {
				assert.NoError(t, err, "unexpected error from UpdateProject")
			}
		})
	}
}

func TestDeleteProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo *mocks.MockRepository
		uc   UseCases
	}

	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		setup   func(f *fields)
		args    args
		wantErr bool
	}{
		{
			name: "success",
			setup: func(f *fields) {
				f.repo.EXPECT().DeleteProject(gomock.Any(), int64(10)).
					Return(nil)
			},
			args: args{ctx: context.Background(), id: 10},
		},
		{
			name: "repo error",
			setup: func(f *fields) {
				f.repo.EXPECT().DeleteProject(gomock.Any(), gomock.Any()).
					Return(errors.New("delete fail"))
			},
			args:    args{ctx: context.Background(), id: 99},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repoMock := mocks.NewMockRepository(ctrl)
			uc := NewUseCases(repoMock, nil, nil, nil, nil, nil)

			f := fields{repo: repoMock, uc: uc}
			tt.setup(&f)

			err := f.uc.DeleteProject(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				assert.Error(t, err, "expected error from DeleteProject")
			} else {
				assert.NoError(t, err, "unexpected error from DeleteProject")
			}
		})
	}
}
