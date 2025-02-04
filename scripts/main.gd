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
		if !result.has("position"):
			print("Failed to get position from ray cast:", result)
			return
		print(result.position)
		var target_id: String = str(round(result.position.x)) +","+ str(round(result.position.z))
		print(target_id)
		var payload = JSON.stringify({"action": "move", "targetId": target_id})
		_client.send(payload.to_utf8_buffer())
		

func _handle_client_connected() -> void:
	print("Client connected to server.")
	var rng = RandomNumberGenerator.new()
	var pid = str(rng.randi_range(1, 10000))
	var payload := {"token": pid}
	var str := JSON.stringify(payload)
	var message := str.to_utf8_buffer()
	_client.send(message)
	var c := character_scene.instantiate()
	c.character_id = pid
	$Characters.add_child(c)

func _handle_client_data(data: PackedByteArray) -> void:
	var utf8 := data.get_string_from_utf8()
	var splits := utf8.split("\n", false)
	for s in splits:
		var json = JSON.parse_string(s)
		if json == null:
			print("Failed to parse tick json:", s, str, splits)
		var tick := Tick.parse(json)
		if tick == null:
			print("Failed to parse tick:", s)
			return
		var players: Array = tick.npcs.duplicate()
		players.append(tick.player)
		for delta in players:
			var exists := $Characters.get_children().any(func (n):
				if n is Character:
					return n.character_id == delta.id
				return false
			)
			if not exists:
				var c := character_scene.instantiate()
				c.character_id = delta.id
				$Characters.add_child(c)
			for node in $Characters.get_children():
				if node is Character and node.character_id == delta.id:
					print(delta.position.x, delta.position.y)
					node.move(float(delta.position.x), float(delta.position.y))
					
		for node in $Characters.get_children():
			if node is Character:
				if tick.npcs.map(func (c): c.id).has(node.character_id):
					$Characters.remove_child(node)

func _handle_client_disconnected() -> void:
	print("Client disconnected from server.")
	_connect_after_timeout(RECONNECT_TIMEOUT)

func _handle_client_error() -> void:
	print("Client error.")
	_connect_after_timeout(RECONNECT_TIMEOUT)
