extends Resource
class_name Tick

@export var tick:         int    
@export var tick_started_at: int    
@export var tick_ended_at:   int           
@export var sent_at:        int            
@export var instance:      String          
@export var toggles:       Toggles   
@export var inventory:     Array
@export var experience:    Array
@export var equipment:     Array
@export var player:        EntityDelta    
@export var npcs:          Array[EntityDelta]

static func parse(d: Dictionary) -> Tick:
	var st := Tick.new()
	var foo: float = 0
	assert(d.tick is float)
	assert(d.tickStartedAt is float)
	assert(d.tickEndedAt is float)
	assert(d.sentAt is float)
	assert(d.instance is String)
	assert(d.inventory is Array)
	assert(d.experience is Array)
	assert(d.equipment is Array)
	assert(d.player is Dictionary)
	assert(d.npcs is Array)
	st.tick = d.tick
	st.tick_started_at = d.tickStartedAt
	st.tick_ended_at = d.tickEndedAt
	st.sent_at = d.sentAt
	st.instance = d.instance
	st.toggles = Toggles.parse(d.toggles)
	st.inventory = []
	st.experience = []
	st.equipment = []
	st.player = EntityDelta.parse(d.player)
	st.npcs = []
	return st
