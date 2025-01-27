extends CharacterBody3D

var m3_pressed: bool = false

var direction: Vector3
var true_tile := Vector2(position.x, position.z)
const DURATION: float = 0.6
const gravity = 9.1

var character_id: String

var phi: float = 0
var theta: float = 70:
	set(value):
		theta = clamp(value, 5, 70)

var left: bool = false
var right: bool = false
var up: bool = false
var down: bool = false

#func done_moving() -> bool:
	## P = start + t * direction
	#return ((position - start) / velocity) >= 1

func _physics_process(delta: float) -> void:
	if not is_on_floor():
		velocity.y -= gravity * delta
	move_and_slide()
	var desired_angle = atan2(direction.x, direction.z)
	$MeshInstance3D.rotation.y = lerp_angle($MeshInstance3D.rotation.y, desired_angle, .1)
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


func _input(event: InputEvent) -> void:
	if event is InputEventPanGesture:
		$h/v/PlayerCamera.position.z = clamp($h/v/PlayerCamera.position.z + event.delta.y, 1, 20)
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
		$h/v/PlayerCamera.position.z = clamp($h/v/PlayerCamera.position.z - 1, 1, 20)
	if event.is_action_pressed("wheeldown"): 
		print("wheel down")
		$h/v/PlayerCamera.position.z = clamp($h/v/PlayerCamera.position.z + 1, 1, 20)
	if Input.is_action_just_released("leftarrow"): 
		left = false
	if Input.is_action_just_released("rightarrow"):
		right = false
	if Input.is_action_just_released("uparrow"):
		up = false
	if Input.is_action_just_released("downarrow"): 
		down = false
	
func move(x: float, z: float) -> void:
	if(true_tile == Vector2(x, z)):
		velocity= Vector3(0, 0, 0)
	else:
		true_tile = Vector2(x, z)
		$TrueTile.position = Vector3(x, .01, z)
		direction = (transform.basis * (Vector3(x, position.y, z) - position)).normalized()
		velocity = (Vector3(x, position.y, z) - position) / DURATION
	
