package types

/*
	TODO:
		- Add Token CRUD
*/

func createBlankToken(name string) *Token {
	return &Token{
		Name: name,
	}
}

func createToken(name string, state States, amount uint32) *Token {
	return &Token{
		Name:   name,
		States: createStateHistory(state, amount),
	}
}

func createTokens(names []string, states []States, amounts []uint32) []*Token {
	var tokens []*Token

	for i := range names {
		tokens[i] = createToken(names[i], states[i], amounts[i])
	}

	return tokens
}

func createTokenWithHistory(name string, stateHistory []*StateHistory) *Token {
	return &Token{
		Name:   name,
		States: stateHistory,
	}
}

func (t *Token) setTokenStatusWithIndex(name string, state States, index uint32) {
	states := t.GetStates()
	states[index].State = state
}
