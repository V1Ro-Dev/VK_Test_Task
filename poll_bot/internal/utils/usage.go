package utils

func Usage(command string) string {
	switch command {
	case "create":
		return "Correct usage: /poll create <question> <option1, option2, ..."

	case "vote":
		return "Correct usage: /poll vote <poll_id> <option>"

	case "check_results":
		return "Correct usage: /poll check_results <poll_id>"

	case "end":
		return "Correct usage: /poll end <poll_id>"

	case "del":
		return "Correct usage: /poll del <poll_id>"

	default:
		return "Unknown command: " + command
	}
}
