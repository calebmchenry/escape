extends Resource
class_name Tick

@export var tick: int
@export var delta: Delta

static func parse(d: Dictionary) -> Tick:
	var st := Tick.new()
	assert(d.tick)
	st.tick = d.tick
	assert(d.delta is Dictionary)
	return st
