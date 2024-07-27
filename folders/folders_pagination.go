package folders

import (
	"fmt"

	"github.com/gofrs/uuid"
)

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started

// SOLUTION EXPLANATION
// An API on FoldersPaginated struct has been created with the ability to:
//  - RequestPage(): continually request the next page until there are no more pages to request.
//  - GetCurrPage(): returns the specified pageNumth folderPage.
//  - ResetCurrPage(): resets the current page of pagination back to the first folder page.
//  - SetCurrPage(): sets the current page of pagination to the specified pageNumth page.

// Example usage of folders_pagination located at the bottom of folders_test.go at line 85

// Object to store the current state of the pagination.
type FoldersPaginated struct {
	currPage int
	folderPages [][]*Folder
}

// Works like an iterator, returning the current page of folders and the page index to get the next page thereafter.
// Iterates the current page to the next. Repeated calls to RequestPage() will keep on returning the next
// available page.
// When we reach the last page, it keeps returning the last page and the last page index until reset.
func (fp *FoldersPaginated) RequestPage() ([]*Folder, int) {
	if len(fp.folderPages) == 0 {
		return []*Folder{}, 0
	}

	prevPage := fp.currPage
	if fp.currPage < len(fp.folderPages) - 1 {
		fp.currPage += 1
	}
	return fp.folderPages[prevPage], fp.currPage
}

// Returns the pageNumth folderPage given a pageNum. Returns an error if the provided pageNum is out of range.
func (fp *FoldersPaginated) GetCurrPage(pageNum int) ([]*Folder, error) {
	if pageNum < 0 || pageNum >= len(fp.folderPages) {
		return []*Folder{}, fmt.Errorf("pageNum %d out of range. pageNum must be in the range [0, %d)", pageNum, len(fp.folderPages))
	}
	return fp.folderPages[pageNum], nil
}

// Resets the current page of FoldersPaginated struct to the first folderPage with index 0.
func (fp *FoldersPaginated) ResetCurrPage() {
	fp.currPage = 0
}

// Sets the current page to specified pageNum. Returns an error immediately if the pageNum is not in an appropriate
// range.
func (fp *FoldersPaginated) SetCurrPage(pageNum int) error {
	if pageNum < 0 || pageNum >= len(fp.folderPages) {
		return fmt.Errorf("pageNum %d out of range. pageNum must be in the range [0, %d)", pageNum, len(fp.folderPages))
	}
	fp.currPage = pageNum
	return nil
}

// Works the same as the original GetAllFolders() function. It however also returns us a FoldersPaginated struct which
// provides pagination functionality via the API of functions above.
func GetAllFoldersPagination(req *FetchFolderRequest, pageSize int) (*FetchFolderResponse, *FoldersPaginated, error) {
	orgIDFolders, err := FetchAllFoldersByOrgID(req.OrgID)
	foldersPaginated := FoldersPaginated{currPage: 0}
	
	currPage := -1
	for index, folder := range orgIDFolders {
		if index % pageSize == 0 {
			foldersPaginated.folderPages = append(foldersPaginated.folderPages, []*Folder{})
			currPage += 1
		}
		foldersPaginated.folderPages[currPage] = append(foldersPaginated.folderPages[currPage], folder)
	}

	return &FetchFolderResponse{Folders: orgIDFolders}, &foldersPaginated, err
}

func FetchAllFoldersByOrgIDPagination(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
