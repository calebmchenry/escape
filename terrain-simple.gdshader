shader_type spatial;

void vertex() {
  VERTEX.y += cos(VERTEX.x) * sin(VERTEX.z);
}

void fragment() {
  NORMAL = vec3(0, 0, 0);
  ALBEDO.rgb = vec3(5.0, 0.1, 0.2);
}