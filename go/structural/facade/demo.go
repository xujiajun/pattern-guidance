package facade

type UserFacade struct {
	U User
	G Group
}

type User struct {
	name string
}

type Group struct {
	name string
}

func (user *User)getUserName(name string) string  {
	return "username:"+name
}

func (user *Group)getGroupName(name string) string  {
	return "groupname:"+name
}

func (facade *UserFacade)GetUserProfile(username string,groupname string) string  {
	return facade.U.getUserName(username)+"@"+facade.G.getGroupName(groupname)
}

