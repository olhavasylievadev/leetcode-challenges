func numUniqueEmails(emails []string) int {
    uniqueMails := make(map[string]bool)
    
    for _, email := range emails {
        uniqueMails[formatEmail(email)] = true
    }
    
    return len(uniqueMails)
}

func formatEmail(s string) string {
    slice := strings.Split(s, "@")
	var sb strings.Builder

	for _, char := range slice[0] {
		if char == '.' {
			continue
		} else if char == '+' {
			break
		} else {
			sb.WriteByte(byte(char))
		}
	}
	sb.WriteString(fmt.Sprintf("@%s", slice[1]))

	return sb.String()
}