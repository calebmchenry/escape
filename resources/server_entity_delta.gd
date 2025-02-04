extends Resource
class_name EntityDelta

@export var id:       String  
@export var type:      String   
@export var position:  ServerPosition
@export var animation: String
@export var target_id:  String
@export var actions:   Array[ServerAction]

static func parse(d: Dictionary) -> EntityDelta:
	var entity := EntityDelta.new()
	assert(d.id is String)
	assert(d.type is String)
	assert(d.position == null || d.position is Dictionary)
	if d.has("position"):
		assert(d.position is Dictionary)
		entity.position = ServerPosition.parse(d.position)
	if d.has("animation"):
		assert(d.animation is String)
		entity.animation = d.animation
	if d.has("targetId"):
		assert(d.targetId is String)
		entity.target_id = d.targetId
	assert(d.actions is Array)
	entity.id = d.id
	entity.type = d.type
	entity.actions = []
	for action in d.actions: 
		entity.actions.append(ServerAction.parse(action))
	return entity
