extends Resource
class_name CharacterDelta

@export var id: String
@export var position: ServerCoordinate

static func parse(d: Dictionary) -> CharacterDelta:
	var character := CharacterDelta.new()
	assert(d.characterId is String)
	character.id = d.characterId
	assert(d.position is Dictionary)
	character.position = ServerCoordinate.parse(d.position)
	return character
