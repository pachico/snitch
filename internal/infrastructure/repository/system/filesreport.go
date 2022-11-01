package system

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"pachico/snitch/internal/domain"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

var skippedFilesReportDirectories = map[string]bool{
	"/boot":  true,
	"/cdrom": true,
	"/dev":   true,
	"/proc":  true,
	"/run":   true,
	"/sys":   true,
	"/tmp":   true,
	"/var":   true,
	"/etc":   true,
}

type FSReportRepository struct{}

func (fr *FSReportRepository) GetFSReport() (domain.FSReport, error) {
	var fsReport domain.FSReport
	maxDepth := 2

	err := filepath.WalkDir("/", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				return nil
			}
			return nil
		}

		currentDepth := strings.Count(path, string(os.PathSeparator))
		if currentDepth > maxDepth {
			return fs.SkipDir
		}

		if _, skipped := skippedFilesReportDirectories[filepath.Clean(path)]; skipped {
			return fs.SkipDir
		}

		if d.Type()&fs.ModeSymlink != 0 {
			// Resolve symlink, but ignore errors resulting from broken links
			resolvedPath, err := filepath.EvalSymlinks(path)
			if err != nil {
				fmt.Printf("Error resolving symlink %s: %v\n", path, err)
				return nil
			}
			path = resolvedPath
		}

		info, err := d.Info()
		if err != nil {
			// Log errors related to retrieving file info but continue the walk.
			fmt.Printf("Error retrieving info for %s: %v\n", path, err)
			return nil
		}

		file := domain.File{
			Name:       path,
			Size:       info.Size(),
			IsDir:      info.IsDir(),
			Owner:      getOwner(info),
			Permission: info.Mode().Perm().String(),
		}
		fsReport = append(fsReport, file)

		return nil
	})

	return fsReport, err
}

func getOwner(info fs.FileInfo) string {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return "Unknown"
	}
	user, err := user.LookupId(strconv.Itoa(int(stat.Uid)))
	if err != nil {
		fmt.Printf("Error looking up user ID %d: %v\n", stat.Uid, err)
		return "Unknown"
	}
	return user.Username
}
