extends Resource
class_name Delta

@export var characters: Array[CharacterDelta]

static func parse(d: Dictionary) -> Delta:
	var delta := Delta.new()
	var cs: Array[CharacterDelta] = []
	for cd in d.characters:
		cs.append(CharacterDelta.parse(cd))
	delta.characters = cs
	return delta
