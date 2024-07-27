package folders

import (
	"github.com/gofrs/uuid"
)

// GetAllFolders is a function which retrieves a list of folders from an underlying json file
// (sample.json), turning the json's key-value string data into easier to work with and
// abstracted `Folder` struct objects.
//
// GetAllFolders accepts a FetchFolderRequest specifying the OrgID we want to fetch folders
// from, and returns us a FetchFolderResponse (consisting of a list of pointers to Folder structs of
// that specific orgID) and an error (if any).
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// var (
	// 	err error		// In this function's initial state, this Go program does not run because of compiler errors
	// 	f1  Folder		// related to linting of unused variables (err, f1, fs, k, k1).
	// 	fs  []*Folder	// There is a lot of unnecessary variable declaration and assignment in the code below and the code
	// )				// can be simplified significantly.
	f := []Folder{}		// Lack of descriptive variable names makes the code hard to reason with, especially in the below
						// for range loops.
	r, _ := FetchAllFoldersByOrgID(req.OrgID)	// The error returned by FetchAllFoldersByOrgID should not be discarded
												// and should be appropriately handled by this outer function GetAllFolders.
												// However, as explained in the below comments, the lack of error generation
												// from GetSampleData() is propagated up to this top level function.
												// However, this will not be addressed because GetSampleData lies outside the
												// file folders.go. However, if I could change static.go, proper propagation of
												// errors would be implemented.
	for _, v := range r {	// Lack of descriptive variable names makes code hard to understand.
		f = append(f, *v)
	}
	var fp []*Folder
	for _, v1 := range f { 	// There is no need for these 2 for range loops and we can remove them. 
								// The first for range loop at line 29 is simply dereferencing the pointers to Folder structs, and allocating new memory on the stack for them in 
		fp = append(fp, &v1)	// the variable f, but this loop here at line 33 is simply just collating a slice of pointers once again,
	}							// which is just what FetchAllFoldersByOrgID() initially returned anyways.

	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}		// Code can be made concise by using declaration + assignment with :=
	return ffr, nil 	// Hard-coded error makes the point of having an error in the first place redundant.
}


// NEW IMPROVED IMPLEMENTATION OF GetAllFolders()
func GetAllFoldersNew(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	orgIDFolders, err := FetchAllFoldersByOrgID(req.OrgID)
	return &FetchFolderResponse{Folders: orgIDFolders}, err
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData() 	// As the function which does the underlying reading of the sample.json file, there should be error handling,
								// especially when interacting with I/O (e.g., io.ReadAll() can return an error, but the error is simply discarded).
								// However, the implementation of GetSampleData shows that a lot of the errors are simply discarded. This does not
								// give good visibility to the functions which call GetSampleData() about what could have gone wrong to potentially
								// return an empty list of Folder structs.

	resFolder := []*Folder{}						// resFolders stores pointers to Folder structs filtered by the passed in orgID.
	for _, folder := range folders {				// This for range loop is a good example of efficient memory management.
		if folder.OrgId == orgID {					// Instead of allocating new memory on the heap for each of the returned Folder structs,
			resFolder = append(resFolder, folder)	// we simply store a list of pointers to Folder structs because the memory allocation has
		}											// already been done in GetSampleData() by json.Unmarshal.
	}
	return resFolder, nil 		// The purpose of the error is made redundant if it is hardcoded to return nil.
}
