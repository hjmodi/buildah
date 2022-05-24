
	"sourcegraph.com/sqs/pbtypes"
// ParseMultiFileDiff parses a multi-file unified diff. It returns an error if parsing failed as a whole, but does its
// best to parse as many files in the case of per-file errors. In the case of non-fatal per-file errors, the error
// return value is null and the Errs field in the returned MultiFileDiff is set.
	if err != nil {
		ts := pbtypes.NewTimestamp(*origTime)
		fd.OrigTime = &ts
		ts := pbtypes.NewTimestamp(*newTime)
		fd.NewTime = &ts
// timestamps).
	return fmt.Sprintf("overflowed into next file: %s", e)
	case (len(fd.Extended) == 3 || len(fd.Extended) == 4 && strings.HasPrefix(fd.Extended[3], "Binary files ")) &&
		strings.HasPrefix(fd.Extended[1], "new file mode ") && strings.HasPrefix(fd.Extended[0], "diff --git "):
		fd.NewName = names[1]
	case (len(fd.Extended) == 3 || len(fd.Extended) == 4 && strings.HasPrefix(fd.Extended[3], "Binary files ")) &&
		strings.HasPrefix(fd.Extended[1], "deleted file mode ") && strings.HasPrefix(fd.Extended[0], "diff --git "):
		fd.OrigName = names[0]
	case len(fd.Extended) == 4 && strings.HasPrefix(fd.Extended[2], "rename from ") && strings.HasPrefix(fd.Extended[3], "rename to ") && strings.HasPrefix(fd.Extended[0], "diff --git "):
	case len(fd.Extended) == 3 && strings.HasPrefix(fd.Extended[2], "Binary files ") && strings.HasPrefix(fd.Extended[0], "diff --git "):
		fd.OrigName = names[0]
		fd.NewName = names[1]
				// Saw start of new hunk, so this hunk is
				// complete. But we've already read in the next hunk's