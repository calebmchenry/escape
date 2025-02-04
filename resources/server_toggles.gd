extends Resource
class_name Toggles

@export var run: bool

static func parse(d: Dictionary) -> Toggles:
	var toggles := Toggles.new()
	assert(d.run is bool)
	toggles.run = d.run
	return toggles
