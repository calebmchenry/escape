extends Resource
class_name ServerAction

@export var command: String
@export var label: String

static func parse(d: Dictionary) -> ServerAction:
	var action := ServerAction.new()
	assert(d.command is String)
	assert(d.label is String)
	action.command = d.command
	action.label = d.label
	return action
