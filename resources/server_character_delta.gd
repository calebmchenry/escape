extends Resource
class_name CharacterDelta

@export var id: String
@export var position: ServerCoordinate

static func parse(d: Dictionary) -> CharacterDelta:
	var character := CharacterDelta.new()
	assert(d.characterId)
	character.id = d.characterId
	return character
