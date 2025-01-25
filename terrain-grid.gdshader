shader_type spatial;

uniform float height_scale = 0.1;
uniform float grid_size = 1;
uniform sampler2D noise;
uniform float xz_scale = 1.0;

varying vec2 tex_position;

void vertex() {
  tex_position = VERTEX.xz / 2.0 + 0.5;
  float height = texture(noise, tex_position).x;
  VERTEX.y += height * height_scale;
}

void fragment() {
  //if (abs(fract(VERTEX.x / grid_size.x)) < 0.01 || abs(fract(VERTEX.z / grid_size.y)) < 0.01) {
  	//ALBEDO = vec3(1.0, 1.0, 1.0);
  //} else {
	ALBEDO = vec3(abs((VERTEX.x + xz_scale / 2.0) / xz_scale), 0.0, abs((VERTEX.z + xz_scale / 2.0 ) / xz_scale));
  //}
}