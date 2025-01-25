extends Resource
class_name Delta

@export var characters: Array[CharacterDelta]

static func parse(d: Dictionary) -> Delta:
	var delta := Delta.new()
	assert(d.characters is Array)
	delta.characters = d.characters.map(func (c): CharacterDelta.parse(c))
	return delta
