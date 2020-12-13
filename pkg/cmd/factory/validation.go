package factory

func ValidateCollectorFilter(filter string) bool {
	switch filter {
	case
		"installed",
		"hosted",
		"dead",
		"alive":
		return true
	}
	return false
}

func ValidateRoleOutput(output string) bool {
	switch output {
	case
		"name",
		"description",
		"filterPredicate",
		"users",
		"capabilities",
		"id":
		return true
	}
	return false
}

func ValidateUserSortBy(sortBy string) bool {
	switch sortBy {
	case
		"firstName",
		"lastName",
		"email":
		return true
	}
	return false
}

func ValidateUserOutput(output string) bool {
	switch output {
	case
		"firstName",
		"lastName",
		"email",
		"roleIds",
		"id",
		"isActive",
		"isLocked",
		"isMfaEnabled",
		"lastLoginTimestamp":
		return true
	}
	return false
}
