package main

type WinCondition interface {
	IsWon(grid Grid) bool
}

type TopRightToBottomLeft struct{}

func (v TopRightToBottomLeft) IsWon(grid Grid) bool {
	if grid.fields[2] == grid.fields[4] && grid.fields[4] == grid.fields[6] && grid.fields[2] != NoPlayer {
		return true
	}

	return false
}

type TopLeftToBottomRight struct{}

func (v TopLeftToBottomRight) IsWon(grid Grid) bool {
	if grid.fields[0] == grid.fields[4] && grid.fields[4] == grid.fields[8] && grid.fields[0] != NoPlayer {
		return true
	}

	return false
}

type Column struct{}

func (v Column) IsWon(grid Grid) bool {
	for i := 0; i < len(grid.fields); i += 3 {
		if grid.fields[i] == grid.fields[i+1] && grid.fields[i+1] == grid.fields[i+2] && grid.fields[i] != NoPlayer {
			return true
		}
	}

	return false
}

type Row struct{}

func (v Row) IsWon(grid Grid) bool {
	for i := 0; i < len(grid.fields); i += 3 {
		if grid.fields[i] == grid.fields[i+1] && grid.fields[i+1] == grid.fields[i+2] && grid.fields[i] != NoPlayer {
			return true
		}
	}

	return false
}

func GetWinConditions() []WinCondition {
	return []WinCondition{
		TopLeftToBottomRight{},
		TopRightToBottomLeft{},
		Row{},
		Column{},
	}
}
