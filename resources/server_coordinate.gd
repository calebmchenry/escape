extends Resource
class_name ServerCoordinate

@export var x: float
@export var y: float

static func parse(d: Dictionary) -> ServerCoordinate:
	var coord := ServerCoordinate.new()
	assert(d.x is int)
	assert(d.y is int)
	coord.x = float(d.x)
	coord.y = float(d.y)
	return coord
