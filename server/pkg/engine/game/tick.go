package game

type ServerTick struct {
	Tick          int               `json:"tick"`
	TickStartedAt int               `json:"tickStartedAt"`
	TickEndedAt   int               `json:"tickEndedAt"`
	SentAt        int               `json:"sentAt"`
	Instance      string            `json:"instance"`
	Toggles       TogglesPayload    `json:"toggles"`
	Inventory     []InventoryDelta  `json:"inventory"`
	Experience    []ExperienceDelta `json:"experience"`
	Equipment     []EquipmentDelta  `json:"equipment"`
	Player        *EntityDelta      `json:"player"`
	NPCs          []*EntityDelta    `json:"npcs"`
}

type EntityDelta struct {
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	Position  Position `json:"position,omitempty"`
	Animation string   `json:"animation,omitempty"`
	TargetID  string   `json:"targetId,omitempty"`
	Actions   []Action `json:"actions"`
}

type Position struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	Size int `json:"size"`
}

type TogglesPayload struct {
	Run bool `json:"run"`
}

type InventoryDelta struct {
	ItemID  string   `json:"itemId"`
	Index   int      `json:"index"`
	Actions []Action `json:"action"`
}

type Action struct {
	Command string `json:"command"`
	Label   string `json:"label"`
}

type ExperienceDelta struct {
	Skill  string `json:"skill"`
	Amount int    `json:"amount"`
}

type EquipmentDelta struct {
	Slot   string `json:"slot"`
	ItemID string `json:"itemId"`

	MeleeAccuracy  int `json:"meleeAccuracy"`
	RangedAccuracy int `json:"rangedAccuracy"`
	MagicAccuracy  int `json:"magicAccuracy"`

	MeleeDamage  int `json:"meleeDamage"`
	RangedDamage int `json:"rangedDamage"`
	MagicDamage  int `json:"magicDamage"`

	MeleeDefense  int `json:"meleeDefense"`
	RangedDefense int `json:"rangedDefense"`
	MagicDefense  int `json:"magicDefense"`
}
