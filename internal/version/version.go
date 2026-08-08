package version

var (
    Version = "dev"
    Commit  = "none"
    Date    = "unknown"
)

func Full() string {
    return "PortKiller " + Version + " (" + Commit + ", " + Date + ")"
}