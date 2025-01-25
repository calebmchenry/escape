extends CharacterBody3D

var m3_pressed: bool = false

var start: Vector3
var destination: Vector3
var duration: float = 0.6

var character_id: String

var radius: float = 2:
	set(value):
		radius = clamp(value, 1, 20)
var phi: float = 0
var theta: float = 70:
	set(value):
		theta = clamp(value, 5, 70)

var left: bool = false
var right: bool = false
var up: bool = false
var down: bool = false

func _process(delta: float) -> void:
	self.position += (destination - start) * (delta / duration)
	var h = $h
	var v = $h/v
	if left:
		h.rotation_degrees.y -= 1
	if right:
		h.rotation_degrees.y += 1
	if up:
		v.rotation_degrees.x = clamp(v.rotation_degrees.x - 1, -90, -20)
	if down:
		v.rotation_degrees.x = clamp(v.rotation_degrees.x + 1, -90, -20)
	$h/v/PlayerCamera.position.z = radius


func _input(event: InputEvent) -> void:
	if Input.is_action_just_pressed("leftarrow"): 
		left = true
	if Input.is_action_just_pressed("rightarrow"):
		right = true
	if Input.is_action_just_pressed("uparrow"):
		up = true
	if Input.is_action_just_pressed("downarrow"): 
		down = true
	if event.is_action_pressed("wheelup"): 
		print("wheel up")
		radius -= 1
	if event.is_action_pressed("wheeldown"): 
		print("wheel down")
		radius += 1
	if Input.is_action_just_released("leftarrow"): 
		left = false
	if Input.is_action_just_released("rightarrow"):
		right = false
	if Input.is_action_just_released("uparrow"):
		up = false
	if Input.is_action_just_released("downarrow"): 
		down = false
	
func move(x: float, z: float) -> void:
	self.start = self.position
	self.destination = Vector3(x, self.position.y, z)
	$MeshInstance3D.look_at(self.destination, Vector3(0, 1000, 0))
	
