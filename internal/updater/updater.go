package updater

var endpoint = "http://localhost:4000/updates"

type UpdateInfo struct {
	Version string
}

func CheckForUpdates() {
	// Every x time, check updates from server y
	// Cross-check current version against update info api call version

	// Should try to update if it conditions are met
}
