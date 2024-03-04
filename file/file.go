package file

// [filename] [description] [created at] [foldername] [username]
type File struct {
	FileName    string
	Description string
	CreatedAt   int

	FolderName string
	UserName   string
}

/*
sort.Slice(files, func(i, j int) bool {
	return users[i].Name < users[j].Name
})


sort.Slice(files, func(i, j int) bool {
	return users[i].CreatedAt < users[j].CreatedAt
})


*/
