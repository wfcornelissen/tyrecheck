# Tyrecheck Project Plan

**Tyrecheck (to be renamed to Wieleman upon completion)**  
From hereon referred to as **WM** or **TC**

## Aim
Tyre management CLI tool for transport companies. Allows data entry specifically related to tyres in companies that manage their own fleets.

## How
WM stores **Truck, Trailer, and Tyre** information in an **SQL database**. During development, **SQLite** is used for ease of use. This information is stored in tables, and all functions of the program are used to either **view, change, add, or remove** data. 

> Note: Removal is never actually done; instead, data is marked as **scrap (archived)**.

## Functionality

### Base
- **Add/Edit/Remove/View** the following:
  - Truck
  - Trailer
  - Tyre

### Higher Level
- **Link Truck and Trailer** to create a combo (combination)
  - **Combo Struct**: This link is saved in its own table with the `fleetnum` of the truck and the `fleetnum` of the trailer.
- **Truck and Tyre**: Assign a tyre to a specific position on the truck.
  - Implementation is undecided.
  - The `tyre` struct has a `location` and `position` attribute.
  - Ensure only one is not null, indicating either location or position.
  - Position can be saved as a combination of `fleetnum` of the truck/trailer with the position on the truck/trailer.
- **Trailer and Tyre**: Implement the same logic as truck and tyre linking.

### Work Done
- **Swop**: Swaps the position of two tyres to reflect physical swapping.
- **Remove**: Clears the position of a tyre and assigns a location for it.
- **Fit**: Adds a tyre to the database and assigns a position to it.
- **Retread (optional at end of Phase 3)**:
  - Marks a tyre as sent away for retreading.
  - When received back from the supplier, it must be either marked as scrap or assigned new attributes.
- **Repair**: Logs a repair on a tyre and stores it in the database.

> With every job done, the odometer reading of the truck must be recorded so that the total amount of kilometers traveled by the tyre can be updated.
> If a tyre is linked to a trailer position, the **combo struct/table** must be used to determine the odometer reading.

## Phases

### Phase 1
- Base commands: **Add, Remove, Edit, View**.
- Base commands completed:
  - Add: ✓
  - Remove: ✓ 
  - Edit: ✓
  - View: ✓

- Base functions to reflect the commands.
- Base functions completed:
  - Add: ✓
  - Remove: ✓ 
  - Edit: ✓
  - View: ✓
    - Need to format the repr better for all structs.

- SQL database interactions for each.
- SQL database interactions completed:
  - Add: ✓
  - Remove: ✓ 
  - Edit: ✓
  - View: ✓

### Phase 2
- Refactor code to move database interactions to their own functions instead of being built in to specific command functions.
- Higher-level commands: **Link and Work Done** commands.
  - Link:
    - Truck and Trailer ✓
      - Funcs finished  ✓
    - Truck and Tyre ✓
      - Funcs finished  ✓
    - Trailer and Tyre ✓
      - Funcs finished  ✓
  - Work Done (tracking work done wrt tyres):
    - Fit ✓
      - Funcs needs ironing
    - Remove
    - Repair
    - Check ✓
      - Funcs finished  ✓
    - Retread
      - Sent
      - Received
    - Swop
    - Rotate (on rim)
- Functions to reflect the commands.
- SQL database interactions for each.

### Phase 3
- **Fault logging**
  - Tyre
    - Delamination
    - Cut
    - Abrasion
    - Uneven wear
      - Diagonal Pattern
        - Likely caused by worn suspension.
      - Side to side Pattern
        - Likely caused by worn or loose bearings.
    - Blow-out
    - Other
  - Valve
    - Leak
    - Failure
  - Rim
    - Bent
    - Worn
    - Other
  
- **Report commands** with functions.
- **Migration to an online database**.

### Phase 4
- **Implementation**
- **Testing**
- **Building for use on other computers**



## Notes
- **Tyre Position** has a widely accepted standard. A truck typically has 10 positions and a trailer has 16 positions. Positions are numbered from 1 to 26 from front left to back right. Each new axle is counted from the left to the right. eg. First axle left is 1, right is 2. Second axle moving from left to right is 3,4,5 and 6. Same goes for each axle thereafter.

