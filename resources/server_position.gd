extends Resource
class_name ServerPosition

@export var x: int
@export var y: int
@export var size: int

static func parse(d: Dictionary) -> ServerPosition:
	var pos := ServerPosition.new()
	assert(d.size is float)
	pos.size = d.size
	assert(d.x == 0 or d.x is float)
	assert(d.y == 0 or d.y is float)
	pos.x = float(d.x)
	pos.y = float(d.y)
	return pos
