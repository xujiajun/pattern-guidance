package abstract_factory

import "fmt"

type UserGroupDao interface {
	GetUserGroup() string
}

type UserProfileDao interface {
	GetUserProfile() string
}

type UserDaoFactory interface {
	GetUserGroupDao() UserGroupDao
	GetUserProfileDao() UserProfileDao
}

type UserGroupDaoImpl struct {
}

func (*UserGroupDaoImpl) GetUserGroup() string {
	return fmt.Sprintf("%s", "GetUserGroup")
}

type UserProfileDaoImpl struct {
}

func (*UserProfileDaoImpl) GetUserProfile() string {
	return fmt.Sprintf("%s", "GetUserProfile")
}

type UserDaoFactoryImpl struct {
}

func (factoryImpl UserDaoFactoryImpl) GetUserGroupDao() UserGroupDao {
	return &UserGroupDaoImpl{}
}

func (factoryImpl UserDaoFactoryImpl) GetUserProfileDao() UserProfileDao {
	return &UserProfileDaoImpl{}
}
