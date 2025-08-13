package data

import (
	"context"

	"myproject/data/dbset"

	mgorm "bsi/kratos/micro/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type UserRepo struct {
	log *log.Helper
	*mgorm.GormDB
}

func NewUserRepo(logger log.Logger, db *mgorm.GormDB) *UserRepo {
	return &UserRepo{
		log:    log.NewHelper(logger),
		GormDB: db,
	}
}

func (r *UserRepo) Create(ctx context.Context, user *dbset.User) error {
	return r.WithContext(ctx).Create(user).Error
}

func (r *UserRepo) FindUserByID(ctx context.Context, id int64) (*dbset.User, error) {
	var (
		err  error
		user = &dbset.User{}
	)
	if err = r.WithContext(ctx).Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) Update(ctx context.Context, updates map[string]any) (*dbset.User, error) {
	user := &dbset.User{}

	// 模拟 Data 层事务
	err := r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Model(user).Where("id = ?", updates["id"]).Updates(updates).Error; err != nil {
			return err
		}

		// 模拟其他操作
		// Some other operations
		return nil
	})

	return user, err
}

func (r *UserRepo) List(ctx context.Context, page, pageSize int) (int64, []*dbset.User, error) {
	users := []*dbset.User{}
	// 联表预加载角色和部门
	// r.db.Preload("Role").Preload("Department")
	total := int64(0)
	if err := r.WithContext(ctx).Model(&dbset.User{}).Count(&total).Error; err != nil {
		return 0, nil, err
	}
	if err := r.WithContext(ctx).Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return 0, nil, err
	}
	return total, users, nil
}
