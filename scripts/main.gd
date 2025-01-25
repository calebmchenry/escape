extends Node3D

const HOST: String = "127.0.0.1"
const PORT: int = 8080
const RECONNECT_TIMEOUT: float = 0.6

const Client = preload("res://scripts/client.gd")
const character_scene = preload("res://scenes/character.tscn")
const Character = preload("res://scripts/character.gd")
var _client: Client = Client.new()

func _ready() -> void:
	_client.connected.connect(_handle_client_connected)
	_client.disconnected.connect(_handle_client_disconnected)
	_client.error.connect(_handle_client_error)
	_client.data.connect(_handle_client_data)
	add_child(_client)
	_client.connect_to_host(HOST, PORT)
	

func _connect_after_timeout(timeout: float) -> void:
	await get_tree().create_timer(timeout).timeout
	_client.connect_to_host(HOST, PORT)
	
func _process(delta: float) -> void:
	var camera := get_viewport().get_camera_3d()
	if Input.is_action_just_pressed("primaryclick"):
		var query = PhysicsRayQueryParameters3D.new()
		query.from = camera.project_ray_origin(get_viewport().get_mouse_position())
		query.to = query.from + camera.project_ray_normal(get_viewport().get_mouse_position()) * 1000
		var result = get_world_3d().direct_space_state.intersect_ray(query)
		print(result)
		var actionId = JSON.stringify({
			"x": round(result.position.x),
			"y": round(result.position.z)
		})
		var payload = JSON.stringify({"actionId": actionId})
		print("payload:", payload)
		_client.send(payload.to_utf8_buffer())
		

func _handle_client_connected() -> void:
	print("Client connected to server.")
	var rng = RandomNumberGenerator.new()
	var pid = str(rng.randi_range(1, 10000))
	print("token: ", pid)
	var payload := {"token": pid}
	var str := JSON.stringify(payload)
	var message := str.to_utf8_buffer()
	_client.send(message)
	var c := character_scene.instantiate()
	c.character_id = pid
	$Characters.add_child(c)

func _handle_client_data(data: PackedByteArray) -> void:
	var tick = JSON.parse_string(data.get_string_from_utf8())
	if not tick:
		print("i don't know what this is:", data.get_string_from_utf8())
		return
	for character_delta in tick.delta.characters:
		var exists := $Characters.get_children().any(func (n):
			if n is Character:
				return n.character_id == character_delta.characterId
			return false
		)
		if not exists:
			var c := character_scene.instantiate()
			c.character_id = character_delta.characterId
			$Characters.add_child(c)
		for node in $Characters.get_children():
			print("node: ", node.character_id)
			print("delta :", character_delta)
			if node is Character and node.character_id == character_delta.characterId:
				node.move(float(character_delta.position.x), float(character_delta.position.y))
				
	for node in $Characters.get_children():
		if node is Character:
			if tick.delta.characters.map(func (c): c.characterId).has(node.character_id):
				$Characters.remove_child(node)
	
	

func _handle_client_disconnected() -> void:
	print("Client disconnected from server.")
	_connect_after_timeout(RECONNECT_TIMEOUT)

func _handle_client_error() -> void:
	print("Client error.")
	_connect_after_timeout(RECONNECT_TIMEOUT)
