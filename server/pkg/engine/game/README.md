## Game
--> Commands
Apply as it comes in
  - item interactions {actionId: "use:i-0+i-1"} | {actionId: "swap:i-0+i-1"} | {actionId: "drop:i-28"} | {actionId: "equip:i-11"} | {actionId: "consume:i-12"}
    - equip/unequip
    - swap
    - consume
    - drop
  - toggles {actionId: "toggle:uuid-for-toggle"}
    - stances
    - run
    - special attack
  - set interact target {actionId: "move", target:"-1,5"} {actionId: "attack:p-123"} {actionId: "attack", src: null, target: "npc-1"}
    - move
    - attack
    - interact
    - use item on target
  - Examine

\|/ Tick
Evaluate every 600 ms
  1. Inventory interactions
  2. Apply damages from buffer
  3. Player + NPC actions: Move + interact + attack
    - Is this interaction legal?
      - Non attackable
      - Obstructed
      - Unknown interaction
    - check delay timers
    - perform rolls
    - apply xp
    - add damages to buffers

    Mutex on each player controller preventing commands while executing that player's actions and writing the delta?

    How do I handle movement such that things that would move first don't stand still if there target is "in range" to them but will move away? Do I need a way to calculate targets next square?

  5. Send message queue

<-- Send
  - self
    - xp drops
    - toggles
    - inventory diffs {id: 'blah', actions: ["eat", "drink", "equip"]}
    - equipment
    - position
    - animation
    - target
  - entities (nearby only)
    - position
    - animation
    - varius statuses (door open, cupboard open, tree fell)
    - target
    - actions: ["talk to", "trade", "attack"]


## Actions

Move
- --> Set intent
- \|/ Cancel other targets
- <-- send true tile
- Depends on:
  - Obstacles

Attack
- --> Set target
- \|/ Should cancel movement
- \|/ Should cancel previous attack/interactions/process
- \|/ Can reach ? attack : move
- \|/ delay on attack ? wait : attack
- \|/ Calculate damage
- \|/ Add to damage buffer for target 
- <-? send true tile if movement needed + animation + target
- Depends on: 
  - Target's future position
  - Source's current position
  - Obstacles
  - Source's levels + equipment bonuses
  - Target's levels + equipment bonuses
  - Source's attack delay cycle
  - Target's damage buffer

Interact with items
- --> Queue items to process (bump previously queued)
- \|/ Should not cancel movement
- \|/ Should cancel previous attack/interactions/process
- \|/ Do I have room for the new items ? continue : cancel from queue
- \|/ Do I need to remove/add/replace items
- <-- Send item updates
- Depends on:
  - Inventory contents
  - Equipment
  - Sources current interaction

Interact with resource
- --> Set target
- \|/ Can reach ? interact : move
- \|/ Do I have necessary necessary items
- \|/ Auto drop? Inventory space?
- \|/ Delay on interact ? wait : interact
- \|/ Roll for success
- \|/ Items, xp, node depletion???
- Depends on:
  - Target's position
  - Source's current position
  - Obstacles
  - Source's equipment
  - Source's inventory
  - Source's levels
  - Source's interact delay cycle
  - Target's state

Interact(talk to, smelt, cook, trade) with entity (furnace, stove, trader, etc)
- bring up interface?
- --> Set target
- \|/ Can reach ? interact : move
- Depends on:
  - Target's position
  - Source's current position
  - Obstacles
  - Target's ui information

## Instances???
- Router to send commands to the correct instance?
- How do I transfer a player from one instance into another?
  - Some instances have a starting point but what about back to the main world?

## How do I know a targets future location???
* Scenario: target in range but after movement will be out of range: source should move 1 or 2 spaces towards target
* Scenario: 4 players trying to attack each other around an obstacle all going clock wise
