package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	// "github.com/georgechieng-sc/interns-2022/folders"
	// "github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("test_GetAllFolders", func(t *testing.T) {
		// your test/s here
		{
			// Test basic test case with OrgID: c1556e17-b7c0-45a3-a6ae-9546248fb17a
			resFolders, err := folders.GetAllFoldersNew(&folders.FetchFolderRequest{OrgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")})
			if err != nil {
				t.Error(err)
			}
			if len(resFolders.Folders) != 666 {
				t.Error("There should be 666 folders in sample.json with OrgID: c1556e17-b7c0-45a3-a6ae-9546248fb17a")
			}
		}
		
		{
			// Test basic test case with OrgID: 4212d618-66ff-468a-862d-ea49fef5e183
			resFolders, err := folders.GetAllFoldersNew(&folders.FetchFolderRequest{OrgID: uuid.FromStringOrNil("4212d618-66ff-468a-862d-ea49fef5e183")})
			if err != nil {
				t.Error(err)
			}
			if len(resFolders.Folders) != 1 {
				t.Error("There should be 1 folder in sample.json with OrgID: 4212d618-66ff-468a-862d-ea49fef5e183")
			}
		}

		{
			// Test empty OrgID: ""
			resFolders, err := folders.GetAllFoldersNew(&folders.FetchFolderRequest{OrgID: uuid.FromStringOrNil("")})
			if err != nil {
				t.Error(err)
			}
			if len(resFolders.Folders) != 0 {
				t.Error("There should be 0 folders in sample.json with an empty OrgID")
			}
		}

		{
			// Test nil OrgID
			resFolders, err := folders.GetAllFoldersNew(&folders.FetchFolderRequest{})
			if err != nil {
				t.Error(err)
			}
			if len(resFolders.Folders) != 0 {
				t.Error("There should be 0 folders in sample.json with a nil OrgID")
			}
		}

		{
			// Test OrgID which does not exist in sample.json
			resFolders, err := folders.GetAllFoldersNew(&folders.FetchFolderRequest{OrgID: uuid.FromStringOrNil("invalid-OrgID")})
			if err != nil {
				t.Error(err)
			}
			if len(resFolders.Folders) != 0 {
				t.Error("There should be 0 folders in sample.json with this invalid OrgID")
			}
		}
		
		{
			// Test all folders returned from OrgID: c1556e17-b7c0-45a3-a6ae-9546248fb17a have a unique id
			resFolders, err := folders.GetAllFoldersNew(&folders.FetchFolderRequest{OrgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")})
			if err != nil {
				t.Error(err)
			}
			idMap := make(map[string]int)
			for _, folder := range resFolders.Folders {
				_, exists := idMap[folder.Id.String()]
				if exists {
					t.Error("There is duplicate folder id returned from sample.json:", folder.Id)
				}
			}
		}

		// FOLDERS_PAGINATION EXAMPLE USAGE
		{
			resFolders, foldersPaginated, err := folders.GetAllFoldersPagination(&folders.FetchFolderRequest{OrgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")}, 3)
			if err != nil {
				t.Error(err)
			}
			if len(resFolders.Folders) != 666 {
				t.Error("There should be 666 folders in sample.json with OrgID: c1556e17-b7c0-45a3-a6ae-9546248fb17a")
			}

			// request the first page of folders
			foldersPage, index := foldersPaginated.RequestPage()
			if len(foldersPage) != 3 {
				t.Errorf("Expected 3 pages, but got %d pages", len(foldersPage))
			}
			if index != 1 {
				t.Errorf("Expected index of 1, but got %d index", index)
			}

			// check paginated folders are correct
			expectedFolderIds := []string{"7ee73e98-b5a7-4ff5-a710-bfd8077ac0a9", "5c04651c-e7c0-411f-b4e1-2d99627bd376", "ae058f39-6756-44de-a6c6-cf7d8f484d71"}
			for index, folderId := range expectedFolderIds {
				if (foldersPage[index].Id) != uuid.FromStringOrNil(folderId) {
					t.Errorf("Expected folderID '%s' but got folderID '%s'", folderId, foldersPage[index].Id)
				}
			}

			// request the second page of folders
			foldersPage, index = foldersPaginated.RequestPage()
			if len(foldersPage) != 3 {
				t.Errorf("Expected 3 pages, but got %d pages", len(foldersPage))
			}
			if index != 2 {
				t.Errorf("Expected index of 2, but got %d index", index)
			}
			// check paginated folders are correct
			expectedFolderIds = []string{"c03cfef0-3256-4e46-8ec9-280a913c8592", "994cf912-c9a1-4be6-8225-87c63451ef2d", "c0d25887-f71a-4176-ac3b-6a9bcea92ba4"}
			for index, folderId := range expectedFolderIds {
				if (foldersPage[index].Id) != uuid.FromStringOrNil(folderId) {
					t.Errorf("Expected folderID '%s' but got folderID '%s'", folderId, foldersPage[index].Id)
				}
			}
		}
	})
}
