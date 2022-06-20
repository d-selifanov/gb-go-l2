package main

import (
	"fmt"
	"testing"
)

func TestFilePathStore(t *testing.T) {
	dupFiles := []struct {
		fileInfo fileInfo
		filepath string
	}{
		{fileInfo{"test_file1", 8}, "/test_duplicates/test_sub_folder1/test_file1"},
		{fileInfo{"test_file1", 8}, "/test_duplicates/test_sub_folder2/test_file1"},
		{fileInfo{"test_file1", 8}, "/test_duplicates/test_sub_folder3/test_file1"},
		{fileInfo{"test_file1", 8}, "/test_duplicates/test_sub_folder4/test_file1"},
	}

	noDupFiles := []struct {
		fileInfo fileInfo
		filepath string
	}{
		{fileInfo{"test_file3", 8}, "/test_duplicates/test_file3"},
		{fileInfo{"test_file2", 16}, "/test_duplicates/test_file2"},
		{fileInfo{"test_file3", 16}, "/test_duplicates/test_file3"},
		{fileInfo{"test_file4", 16}, "/test_duplicates/test_file4"},
	}

	d := &duplicates{filePaths: make(map[fileInfo][]string)}

	for _, file := range dupFiles {
		d.save(file.fileInfo, file.filepath)
	}

	for _, file := range noDupFiles {
		d.save(file.fileInfo, file.filepath)
	}

	duplicates := d.getDup()

	for dup := range duplicates {
		fmt.Println(dup)
		if len(dup) != len(dupFiles) {
			t.Fatalf("Expected %d duplicated, got %d", len(dupFiles), len(dup))
		}

		for i, got := range dup {
			want := dupFiles[i].filepath
			if want != got {
				t.Errorf("Expected '%s', got '%s'", want, got)
			}
		}
	}
}
