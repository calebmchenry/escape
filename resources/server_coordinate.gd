extends Resource
class_name ServerCoordinate

@export var x: float
@export var y: float

static func parse(d: Dictionary) -> ServerCoordinate:
	var coord := ServerCoordinate.new()
	assert(d.x == 0 or d.x is float)
	assert(d.y == 0 or d.y is float)
	coord.x = float(d.x)
	coord.y = float(d.y)
	return coord
