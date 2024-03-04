package folder

type Folder struct {
	Name        string
	Description string
	CreatedAt   int
	File        map[string]File
}

/*
sort.Slice(folders, func(i, j int) bool {
	return users[i].Name < users[j].Name
})


sort.Slice(folders, func(i, j int) bool {
	return users[i].CreatedAt < users[j].CreatedAt
})


*/
